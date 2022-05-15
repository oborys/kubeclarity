package scheduler

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/openclarity/kubeclarity/api/server/models"
	"github.com/openclarity/kubeclarity/backend/pkg/database"
	"github.com/openclarity/kubeclarity/backend/pkg/runtime_scanner"
)

type Scheduler struct {
	stopChan chan struct{}
	// Send scan requests through here.
	scanChan  chan *runtime_scanner.ScanConfig
	dbHandler database.Database
}

type SchedulerParams struct {
	Namespaces                    []string
	CisDockerBenchmarkScanEnabled bool
	Interval                      time.Duration
	StartTime                     time.Time
	SingleScan                    bool
}

const (
	ByDaysScheduleScanConfig  = "ByDaysScheduleScanConfig"
	ByHoursScheduleScanConfig = "ByHoursScheduleScanConfig"
	SingleScheduleScanConfig  = "SingleScheduleScanConfig"
	WeeklyScheduleScanConfig  = "WeeklyScheduleScanConfig"
)

func CreateScheduler(scanChan chan *runtime_scanner.ScanConfig, dbHandler database.Database) *Scheduler {
	return &Scheduler{
		stopChan:  make(chan struct{}),
		scanChan:  scanChan,
		dbHandler: dbHandler,
	}
}

func (s *Scheduler) Init() {
	// read last schedule scan config from db
	sched, err := s.dbHandler.SchedulerTable().Get()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("No scheduler config")
			return
		}
		log.Errorf("Failed to get scheduler table: %v", err)
		return
	}
	scanConfig := models.RuntimeScheduleScanConfig{}
	if err := json.Unmarshal([]byte(sched.Config), &scanConfig); err != nil {
		log.Errorf("Failed to unmarshal scheduler config: %v", err)
		return
	}
	startTime, err := time.Parse(time.RFC3339, sched.NextScanTime)
	if err != nil {
		log.Errorf("Failed to parse nextScanTime: %v. %v", sched.NextScanTime, err)
		return
	}

	// start schedule scan
	s.Schedule(&SchedulerParams{
		Namespaces:                    scanConfig.Namespaces,
		CisDockerBenchmarkScanEnabled: scanConfig.CisDockerBenchmarkScanEnabled,
		Interval:                      time.Duration(sched.Interval),
		StartTime:                     startTime,
		SingleScan:                    scanConfig.ScanConfigType().ScheduleScanConfigType() == SingleScheduleScanConfig,
	})
}

func (s *Scheduler) Schedule(params *SchedulerParams) {
	// Clear
	close(s.stopChan)
	s.stopChan = make(chan struct{})

	startsAt := getStartsAt(time.Now().UTC(), params.StartTime, params.Interval)

	go s.spin(params, startsAt)
}

// get the next scan, that is after timeNow. if currentScanTime is already after timeNow, it will be return.
func getNextScanTime(timeNow, currentScanTime time.Time, interval time.Duration) time.Time {
	// if current scan time is before timeNow, jump to the next future scan time
	if currentScanTime.Before(timeNow) {
		// if scan time has passed in less then a second, start a scan now.
		timePassed := timeNow.Sub(currentScanTime)
		if timePassed < time.Second {
			return timeNow
		}
		remainingInterval := timePassed % interval
		if remainingInterval == 0 {
			currentScanTime = timeNow
		} else {
			currentScanTime = timeNow.Add(interval - remainingInterval)
		}
	}
	return currentScanTime
}

// get the time in Duration that the next scan should start at.
func getStartsAt(timeNow time.Time, startTime time.Time, interval time.Duration) time.Duration {
	nextScanTime := getNextScanTime(timeNow, startTime, interval)

	startsAt := nextScanTime.Sub(timeNow)

	return startsAt
}

func (s *Scheduler) spin(params *SchedulerParams, startsAt time.Duration) {
	log.Debugf("Starting a new schedule scan. interval: %v, start time: %v, starts in: %v, namespaces: %v, cisDockerBenchmarkScanEnabled: %v",
		params.Interval, params.StartTime, startsAt, params.Namespaces, params.CisDockerBenchmarkScanEnabled)
	singleScan := params.SingleScan
	interval := params.Interval

	timer := time.NewTimer(startsAt)
	select {
	case <-s.stopChan:
		return
	case <-timer.C:
		go func() {
			ticker := time.NewTicker(interval)
			defer ticker.Stop()
			for {
				if err := s.sendScan(params); err != nil {
					log.Errorf("Failed to send scan: %v", err)
				}
				if singleScan {
					return
				}
				select {
				case <-ticker.C:
				case <-s.stopChan:
					log.Debugf("Received a stop signal...")
					return
				}
			}
		}()
	}
}

func (s *Scheduler) sendScan(params *SchedulerParams) error {
	scanConfig := &runtime_scanner.ScanConfig{
		ScanType:                      models.ScanTypeSCHEDULE,
		CisDockerBenchmarkScanEnabled: params.CisDockerBenchmarkScanEnabled,
		Namespaces:                    params.Namespaces,
	}
	select {
	case s.scanChan <- scanConfig:
	default:
		return fmt.Errorf("failed to send scan config to channel")
	}
	// update next scan time.
	nextScanTime := time.Now().Add(params.Interval).UTC().Format(time.RFC3339)
	if err := s.dbHandler.SchedulerTable().UpdateNextScanTime(nextScanTime); err != nil {
		return fmt.Errorf("UpdateNextScanTime failed: %v", err)
	}
	return nil
}