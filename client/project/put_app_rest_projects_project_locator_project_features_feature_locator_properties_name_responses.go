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

// PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameReader is a Reader for the PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesName structure.
type PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK creates a PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK with default headers values
func NewPutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK() *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK {
	return &PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK{}
}

/*PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK handles this case with default header values.

successful operation
*/
type PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK struct {
	Payload *models.Property
}

func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/projects/{projectLocator}/projectFeatures/{featureLocator}/properties/{name}][%d] putAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK  %+v", 200, o.Payload)
}

func (o *PutAppRestProjectsProjectLocatorProjectFeaturesFeatureLocatorPropertiesNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Property)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
