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

// GetFactionsReader is a Reader for the GetFactions structure.
type GetFactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFactionsOK creates a GetFactionsOK with default headers values
func NewGetFactionsOK() *GetFactionsOK {
	return &GetFactionsOK{}
}

/* GetFactionsOK describes a response with status code 200, with default header values.

An array of factions with historical data
*/
type GetFactionsOK struct {
	Payload []*models.EBGSFactionsPageV5
}

func (o *GetFactionsOK) Error() string {
	return fmt.Sprintf("[GET /factions][%d] getFactionsOK  %+v", 200, o.Payload)
}
func (o *GetFactionsOK) GetPayload() []*models.EBGSFactionsPageV5 {
	return o.Payload
}

func (o *GetFactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
