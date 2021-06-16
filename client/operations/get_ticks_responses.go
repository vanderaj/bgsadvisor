// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bgsadvisor/v2/models"
)

// GetTicksReader is a Reader for the GetTicks structure.
type GetTicksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTicksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTicksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTicksOK creates a GetTicksOK with default headers values
func NewGetTicksOK() *GetTicksOK {
	return &GetTicksOK{}
}

/* GetTicksOK describes a response with status code 200, with default header values.

An array of systems with historical data
*/
type GetTicksOK struct {
	Payload []*models.TickTimesV5
}

func (o *GetTicksOK) Error() string {
	return fmt.Sprintf("[GET /ticks][%d] getTicksOK  %+v", 200, o.Payload)
}
func (o *GetTicksOK) GetPayload() []*models.TickTimesV5 {
	return o.Payload
}

func (o *GetTicksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
