// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// ServeInstanceReader is a Reader for the ServeInstance structure.
type ServeInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeInstanceOK creates a ServeInstanceOK with default headers values
func NewServeInstanceOK() *ServeInstanceOK {
	return &ServeInstanceOK{}
}

/*ServeInstanceOK handles this case with default header values.

successful operation
*/
type ServeInstanceOK struct {
	Payload *models.VcsRootInstance
}

func (o *ServeInstanceOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/vcs-root-instances/{vcsRootInstanceLocator}][%d] serveInstanceOK  %+v", 200, o.Payload)
}

func (o *ServeInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsRootInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
