// Code generated by go-swagger; DO NOT EDIT.

package agent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// GetEnabledInfoReader is a Reader for the GetEnabledInfo structure.
type GetEnabledInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnabledInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEnabledInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEnabledInfoOK creates a GetEnabledInfoOK with default headers values
func NewGetEnabledInfoOK() *GetEnabledInfoOK {
	return &GetEnabledInfoOK{}
}

/*GetEnabledInfoOK handles this case with default header values.

successful operation
*/
type GetEnabledInfoOK struct {
	Payload *models.EnabledInfo
}

func (o *GetEnabledInfoOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/agents/{agentLocator}/enabledInfo][%d] getEnabledInfoOK  %+v", 200, o.Payload)
}

func (o *GetEnabledInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnabledInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
