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

// PutRuntimeScheduleScanConfigReader is a Reader for the PutRuntimeScheduleScanConfig structure.
type PutRuntimeScheduleScanConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRuntimeScheduleScanConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPutRuntimeScheduleScanConfigCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRuntimeScheduleScanConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutRuntimeScheduleScanConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutRuntimeScheduleScanConfigCreated creates a PutRuntimeScheduleScanConfigCreated with default headers values
func NewPutRuntimeScheduleScanConfigCreated() *PutRuntimeScheduleScanConfigCreated {
	return &PutRuntimeScheduleScanConfigCreated{}
}

/* PutRuntimeScheduleScanConfigCreated describes a response with status code 201, with default header values.

Success
*/
type PutRuntimeScheduleScanConfigCreated struct {
	Payload interface{}
}

func (o *PutRuntimeScheduleScanConfigCreated) Error() string {
	return fmt.Sprintf("[PUT /runtime/scheduleScan/config][%d] putRuntimeScheduleScanConfigCreated  %+v", 201, o.Payload)
}
func (o *PutRuntimeScheduleScanConfigCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *PutRuntimeScheduleScanConfigCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRuntimeScheduleScanConfigBadRequest creates a PutRuntimeScheduleScanConfigBadRequest with default headers values
func NewPutRuntimeScheduleScanConfigBadRequest() *PutRuntimeScheduleScanConfigBadRequest {
	return &PutRuntimeScheduleScanConfigBadRequest{}
}

/* PutRuntimeScheduleScanConfigBadRequest describes a response with status code 400, with default header values.

Failed to set scheduled scan config
*/
type PutRuntimeScheduleScanConfigBadRequest struct {
	Payload string
}

func (o *PutRuntimeScheduleScanConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /runtime/scheduleScan/config][%d] putRuntimeScheduleScanConfigBadRequest  %+v", 400, o.Payload)
}
func (o *PutRuntimeScheduleScanConfigBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PutRuntimeScheduleScanConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRuntimeScheduleScanConfigDefault creates a PutRuntimeScheduleScanConfigDefault with default headers values
func NewPutRuntimeScheduleScanConfigDefault(code int) *PutRuntimeScheduleScanConfigDefault {
	return &PutRuntimeScheduleScanConfigDefault{
		_statusCode: code,
	}
}

/* PutRuntimeScheduleScanConfigDefault describes a response with status code -1, with default header values.

unknown error
*/
type PutRuntimeScheduleScanConfigDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the put runtime schedule scan config default response
func (o *PutRuntimeScheduleScanConfigDefault) Code() int {
	return o._statusCode
}

func (o *PutRuntimeScheduleScanConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /runtime/scheduleScan/config][%d] PutRuntimeScheduleScanConfig default  %+v", o._statusCode, o.Payload)
}
func (o *PutRuntimeScheduleScanConfigDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *PutRuntimeScheduleScanConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
