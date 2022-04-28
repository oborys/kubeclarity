// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

<<<<<<< HEAD
	"github.com/cisco-open/kubei/api/client/models"
=======
	"github.com/openclarity/kubeclarity/api/client/models"
>>>>>>> scheduele scan api
)

// GetRuntimeScheduleScanConfigReader is a Reader for the GetRuntimeScheduleScanConfig structure.
type GetRuntimeScheduleScanConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeScheduleScanConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeScheduleScanConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRuntimeScheduleScanConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuntimeScheduleScanConfigOK creates a GetRuntimeScheduleScanConfigOK with default headers values
func NewGetRuntimeScheduleScanConfigOK() *GetRuntimeScheduleScanConfigOK {
	return &GetRuntimeScheduleScanConfigOK{}
}

/* GetRuntimeScheduleScanConfigOK describes a response with status code 200, with default header values.

Success
*/
type GetRuntimeScheduleScanConfigOK struct {
	Payload *models.RuntimeScheduleScanConfig
}

func (o *GetRuntimeScheduleScanConfigOK) Error() string {
	return fmt.Sprintf("[GET /runtime/scheduleScan/config][%d] getRuntimeScheduleScanConfigOK  %+v", 200, o.Payload)
}
func (o *GetRuntimeScheduleScanConfigOK) GetPayload() *models.RuntimeScheduleScanConfig {
	return o.Payload
}

func (o *GetRuntimeScheduleScanConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeScheduleScanConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeScheduleScanConfigDefault creates a GetRuntimeScheduleScanConfigDefault with default headers values
func NewGetRuntimeScheduleScanConfigDefault(code int) *GetRuntimeScheduleScanConfigDefault {
	return &GetRuntimeScheduleScanConfigDefault{
		_statusCode: code,
	}
}

/* GetRuntimeScheduleScanConfigDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetRuntimeScheduleScanConfigDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get runtime schedule scan config default response
func (o *GetRuntimeScheduleScanConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetRuntimeScheduleScanConfigDefault) Error() string {
	return fmt.Sprintf("[GET /runtime/scheduleScan/config][%d] GetRuntimeScheduleScanConfig default  %+v", o._statusCode, o.Payload)
}
func (o *GetRuntimeScheduleScanConfigDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetRuntimeScheduleScanConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
