// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// GetExampleNewProjectDescriptionReader is a Reader for the GetExampleNewProjectDescription structure.
type GetExampleNewProjectDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExampleNewProjectDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetExampleNewProjectDescriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExampleNewProjectDescriptionOK creates a GetExampleNewProjectDescriptionOK with default headers values
func NewGetExampleNewProjectDescriptionOK() *GetExampleNewProjectDescriptionOK {
	return &GetExampleNewProjectDescriptionOK{}
}

/*GetExampleNewProjectDescriptionOK handles this case with default header values.

successful operation
*/
type GetExampleNewProjectDescriptionOK struct {
	Payload *models.NewProjectDescription
}

func (o *GetExampleNewProjectDescriptionOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}/example/newProjectDescription][%d] getExampleNewProjectDescriptionOK  %+v", 200, o.Payload)
}

func (o *GetExampleNewProjectDescriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewProjectDescription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
