// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// GetSystemPropertiesReader is a Reader for the GetSystemProperties structure.
type GetSystemPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSystemPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSystemPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSystemPropertiesOK creates a GetSystemPropertiesOK with default headers values
func NewGetSystemPropertiesOK() *GetSystemPropertiesOK {
	return &GetSystemPropertiesOK{}
}

/*GetSystemPropertiesOK handles this case with default header values.

successful operation
*/
type GetSystemPropertiesOK struct {
	Payload *models.Properties
}

func (o *GetSystemPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/debug/jvm/systemProperties][%d] getSystemPropertiesOK  %+v", 200, o.Payload)
}

func (o *GetSystemPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Properties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
