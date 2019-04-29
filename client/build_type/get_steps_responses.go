// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// GetStepsReader is a Reader for the GetSteps structure.
type GetStepsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStepsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStepsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStepsOK creates a GetStepsOK with default headers values
func NewGetStepsOK() *GetStepsOK {
	return &GetStepsOK{}
}

/*GetStepsOK handles this case with default header values.

successful operation
*/
type GetStepsOK struct {
	Payload *models.Steps
}

func (o *GetStepsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/steps][%d] getStepsOK  %+v", 200, o.Payload)
}

func (o *GetStepsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Steps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
