// Code generated by go-swagger; DO NOT EDIT.

package change

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// GetRelatedBuildTypesReader is a Reader for the GetRelatedBuildTypes structure.
type GetRelatedBuildTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRelatedBuildTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRelatedBuildTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRelatedBuildTypesOK creates a GetRelatedBuildTypesOK with default headers values
func NewGetRelatedBuildTypesOK() *GetRelatedBuildTypesOK {
	return &GetRelatedBuildTypesOK{}
}

/*GetRelatedBuildTypesOK handles this case with default header values.

successful operation
*/
type GetRelatedBuildTypesOK struct {
	Payload *models.BuildTypes
}

func (o *GetRelatedBuildTypesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/changes/{changeLocator}/buildTypes][%d] getRelatedBuildTypesOK  %+v", 200, o.Payload)
}

func (o *GetRelatedBuildTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
