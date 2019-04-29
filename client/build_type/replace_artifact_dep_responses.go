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

// ReplaceArtifactDepReader is a Reader for the ReplaceArtifactDep structure.
type ReplaceArtifactDepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceArtifactDepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceArtifactDepOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceArtifactDepOK creates a ReplaceArtifactDepOK with default headers values
func NewReplaceArtifactDepOK() *ReplaceArtifactDepOK {
	return &ReplaceArtifactDepOK{}
}

/*ReplaceArtifactDepOK handles this case with default header values.

successful operation
*/
type ReplaceArtifactDepOK struct {
	Payload *models.ArtifactDependency
}

func (o *ReplaceArtifactDepOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator}][%d] replaceArtifactDepOK  %+v", 200, o.Payload)
}

func (o *ReplaceArtifactDepOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ArtifactDependency)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
