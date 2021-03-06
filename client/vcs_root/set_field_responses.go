// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetFieldReader is a Reader for the SetField structure.
type SetFieldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetFieldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetFieldOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetFieldOK creates a SetFieldOK with default headers values
func NewSetFieldOK() *SetFieldOK {
	return &SetFieldOK{}
}

/*SetFieldOK handles this case with default header values.

successful operation
*/
type SetFieldOK struct {
	Payload string
}

func (o *SetFieldOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/vcs-roots/{vcsRootLocator}/{field}][%d] setFieldOK  %+v", 200, o.Payload)
}

func (o *SetFieldOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
