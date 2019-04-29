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

// GetFeatureParametersReader is a Reader for the GetFeatureParameters structure.
type GetFeatureParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFeatureParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFeatureParametersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFeatureParametersOK creates a GetFeatureParametersOK with default headers values
func NewGetFeatureParametersOK() *GetFeatureParametersOK {
	return &GetFeatureParametersOK{}
}

/*GetFeatureParametersOK handles this case with default header values.

successful operation
*/
type GetFeatureParametersOK struct {
	Payload *models.Properties
}

func (o *GetFeatureParametersOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/features/{featureId}/parameters][%d] getFeatureParametersOK  %+v", 200, o.Payload)
}

func (o *GetFeatureParametersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Properties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}