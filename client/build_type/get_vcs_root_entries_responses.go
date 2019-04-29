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

// GetVcsRootEntriesReader is a Reader for the GetVcsRootEntries structure.
type GetVcsRootEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVcsRootEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVcsRootEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVcsRootEntriesOK creates a GetVcsRootEntriesOK with default headers values
func NewGetVcsRootEntriesOK() *GetVcsRootEntriesOK {
	return &GetVcsRootEntriesOK{}
}

/*GetVcsRootEntriesOK handles this case with default header values.

successful operation
*/
type GetVcsRootEntriesOK struct {
	Payload *models.VcsRootEntries
}

func (o *GetVcsRootEntriesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/vcs-root-entries][%d] getVcsRootEntriesOK  %+v", 200, o.Payload)
}

func (o *GetVcsRootEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsRootEntries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}