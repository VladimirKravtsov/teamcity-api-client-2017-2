// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// GetLicensingDataReader is a Reader for the GetLicensingData structure.
type GetLicensingDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicensingDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLicensingDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLicensingDataOK creates a GetLicensingDataOK with default headers values
func NewGetLicensingDataOK() *GetLicensingDataOK {
	return &GetLicensingDataOK{}
}

/*GetLicensingDataOK handles this case with default header values.

successful operation
*/
type GetLicensingDataOK struct {
	Payload *models.LicensingData
}

func (o *GetLicensingDataOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/server/licensingData][%d] getLicensingDataOK  %+v", 200, o.Payload)
}

func (o *GetLicensingDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicensingData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
