// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// ServeBuildStatisticValuesReader is a Reader for the ServeBuildStatisticValues structure.
type ServeBuildStatisticValuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeBuildStatisticValuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeBuildStatisticValuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeBuildStatisticValuesOK creates a ServeBuildStatisticValuesOK with default headers values
func NewServeBuildStatisticValuesOK() *ServeBuildStatisticValuesOK {
	return &ServeBuildStatisticValuesOK{}
}

/*ServeBuildStatisticValuesOK handles this case with default header values.

successful operation
*/
type ServeBuildStatisticValuesOK struct {
	Payload *models.Properties
}

func (o *ServeBuildStatisticValuesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}/statistics][%d] serveBuildStatisticValuesOK  %+v", 200, o.Payload)
}

func (o *ServeBuildStatisticValuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Properties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
