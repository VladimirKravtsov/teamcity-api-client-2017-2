// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetCurrentSessionMaxInactiveIntervalReader is a Reader for the SetCurrentSessionMaxInactiveInterval structure.
type SetCurrentSessionMaxInactiveIntervalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetCurrentSessionMaxInactiveIntervalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetCurrentSessionMaxInactiveIntervalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetCurrentSessionMaxInactiveIntervalOK creates a SetCurrentSessionMaxInactiveIntervalOK with default headers values
func NewSetCurrentSessionMaxInactiveIntervalOK() *SetCurrentSessionMaxInactiveIntervalOK {
	return &SetCurrentSessionMaxInactiveIntervalOK{}
}

/*SetCurrentSessionMaxInactiveIntervalOK handles this case with default header values.

successful operation
*/
type SetCurrentSessionMaxInactiveIntervalOK struct {
	Payload string
}

func (o *SetCurrentSessionMaxInactiveIntervalOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/debug/currentRequest/session/maxInactiveSeconds][%d] setCurrentSessionMaxInactiveIntervalOK  %+v", 200, o.Payload)
}

func (o *SetCurrentSessionMaxInactiveIntervalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}