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

// ReplaceFeaturesReader is a Reader for the ReplaceFeatures structure.
type ReplaceFeaturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceFeaturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceFeaturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceFeaturesOK creates a ReplaceFeaturesOK with default headers values
func NewReplaceFeaturesOK() *ReplaceFeaturesOK {
	return &ReplaceFeaturesOK{}
}

/*ReplaceFeaturesOK handles this case with default header values.

successful operation
*/
type ReplaceFeaturesOK struct {
	Payload *models.Features
}

func (o *ReplaceFeaturesOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/features][%d] replaceFeaturesOK  %+v", 200, o.Payload)
}

func (o *ReplaceFeaturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Features)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
