// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ServeBuildFieldShortReader is a Reader for the ServeBuildFieldShort structure.
type ServeBuildFieldShortReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeBuildFieldShortReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeBuildFieldShortOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeBuildFieldShortOK creates a ServeBuildFieldShortOK with default headers values
func NewServeBuildFieldShortOK() *ServeBuildFieldShortOK {
	return &ServeBuildFieldShortOK{}
}

/*ServeBuildFieldShortOK handles this case with default header values.

successful operation
*/
type ServeBuildFieldShortOK struct {
	Payload string
}

func (o *ServeBuildFieldShortOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/{projectLocator}/{btLocator}/{buildLocator}/{field}][%d] serveBuildFieldShortOK  %+v", 200, o.Payload)
}

func (o *ServeBuildFieldShortOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
