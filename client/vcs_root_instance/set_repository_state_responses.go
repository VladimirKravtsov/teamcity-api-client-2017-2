// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// SetRepositoryStateReader is a Reader for the SetRepositoryState structure.
type SetRepositoryStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRepositoryStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetRepositoryStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetRepositoryStateOK creates a SetRepositoryStateOK with default headers values
func NewSetRepositoryStateOK() *SetRepositoryStateOK {
	return &SetRepositoryStateOK{}
}

/*SetRepositoryStateOK handles this case with default header values.

successful operation
*/
type SetRepositoryStateOK struct {
	Payload *models.Entries
}

func (o *SetRepositoryStateOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/repositoryState][%d] setRepositoryStateOK  %+v", 200, o.Payload)
}

func (o *SetRepositoryStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Entries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}