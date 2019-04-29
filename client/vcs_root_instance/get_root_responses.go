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

// GetRootReader is a Reader for the GetRoot structure.
type GetRootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRootOK creates a GetRootOK with default headers values
func NewGetRootOK() *GetRootOK {
	return &GetRootOK{}
}

/*GetRootOK handles this case with default header values.

successful operation
*/
type GetRootOK struct {
	Payload *models.Files
}

func (o *GetRootOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/files/latest][%d] getRootOK  %+v", 200, o.Payload)
}

func (o *GetRootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Files)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
