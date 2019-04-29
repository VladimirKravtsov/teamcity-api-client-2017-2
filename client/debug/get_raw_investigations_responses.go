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

// GetRawInvestigationsReader is a Reader for the GetRawInvestigations structure.
type GetRawInvestigationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRawInvestigationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRawInvestigationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRawInvestigationsOK creates a GetRawInvestigationsOK with default headers values
func NewGetRawInvestigationsOK() *GetRawInvestigationsOK {
	return &GetRawInvestigationsOK{}
}

/*GetRawInvestigationsOK handles this case with default header values.

successful operation
*/
type GetRawInvestigationsOK struct {
	Payload *models.Investigations
}

func (o *GetRawInvestigationsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/debug/investigations][%d] getRawInvestigationsOK  %+v", 200, o.Payload)
}

func (o *GetRawInvestigationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Investigations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
