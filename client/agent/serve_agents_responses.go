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

// ServeAgentsReader is a Reader for the ServeAgents structure.
type ServeAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeAgentsOK creates a ServeAgentsOK with default headers values
func NewServeAgentsOK() *ServeAgentsOK {
	return &ServeAgentsOK{}
}

/*ServeAgentsOK handles this case with default header values.

successful operation
*/
type ServeAgentsOK struct {
	Payload *models.Agents
}

func (o *ServeAgentsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/agents][%d] serveAgentsOK  %+v", 200, o.Payload)
}

func (o *ServeAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Agents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
