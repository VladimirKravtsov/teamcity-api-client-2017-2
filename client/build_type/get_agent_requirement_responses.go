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

// GetAgentRequirementReader is a Reader for the GetAgentRequirement structure.
type GetAgentRequirementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAgentRequirementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAgentRequirementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAgentRequirementOK creates a GetAgentRequirementOK with default headers values
func NewGetAgentRequirementOK() *GetAgentRequirementOK {
	return &GetAgentRequirementOK{}
}

/*GetAgentRequirementOK handles this case with default header values.

successful operation
*/
type GetAgentRequirementOK struct {
	Payload *models.AgentRequirement
}

func (o *GetAgentRequirementOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/agent-requirements/{agentRequirementLocator}][%d] getAgentRequirementOK  %+v", 200, o.Payload)
}

func (o *GetAgentRequirementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentRequirement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
