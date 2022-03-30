// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cisco-open/kubei/sbom_db/api/server/models"
)

// PutSbomDBResourceHashCreatedCode is the HTTP code returned for type PutSbomDBResourceHashCreated
const PutSbomDBResourceHashCreatedCode int = 201

/*PutSbomDBResourceHashCreated SBOM created in DB.

swagger:response putSbomDBResourceHashCreated
*/
type PutSbomDBResourceHashCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SuccessResponse `json:"body,omitempty"`
}

// NewPutSbomDBResourceHashCreated creates PutSbomDBResourceHashCreated with default headers values
func NewPutSbomDBResourceHashCreated() *PutSbomDBResourceHashCreated {

	return &PutSbomDBResourceHashCreated{}
}

// WithPayload adds the payload to the put sbom d b resource hash created response
func (o *PutSbomDBResourceHashCreated) WithPayload(payload *models.SuccessResponse) *PutSbomDBResourceHashCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put sbom d b resource hash created response
func (o *PutSbomDBResourceHashCreated) SetPayload(payload *models.SuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSbomDBResourceHashCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutSbomDBResourceHashDefault Unknown error

swagger:response putSbomDBResourceHashDefault
*/
type PutSbomDBResourceHashDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPutSbomDBResourceHashDefault creates PutSbomDBResourceHashDefault with default headers values
func NewPutSbomDBResourceHashDefault(code int) *PutSbomDBResourceHashDefault {
	if code <= 0 {
		code = 500
	}

	return &PutSbomDBResourceHashDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put sbom d b resource hash default response
func (o *PutSbomDBResourceHashDefault) WithStatusCode(code int) *PutSbomDBResourceHashDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put sbom d b resource hash default response
func (o *PutSbomDBResourceHashDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put sbom d b resource hash default response
func (o *PutSbomDBResourceHashDefault) WithPayload(payload *models.APIResponse) *PutSbomDBResourceHashDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put sbom d b resource hash default response
func (o *PutSbomDBResourceHashDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutSbomDBResourceHashDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
