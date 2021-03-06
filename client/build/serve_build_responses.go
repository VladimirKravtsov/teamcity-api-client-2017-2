// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// ServeBuildReader is a Reader for the ServeBuild structure.
type ServeBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeBuildOK creates a ServeBuildOK with default headers values
func NewServeBuildOK() *ServeBuildOK {
	return &ServeBuildOK{}
}

/*ServeBuildOK handles this case with default header values.

successful operation
*/
type ServeBuildOK struct {
	Payload *models.Build
}

func (o *ServeBuildOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}][%d] serveBuildOK  %+v", 200, o.Payload)
}

func (o *ServeBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Build)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
