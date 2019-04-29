// Code generated by go-swagger; DO NOT EDIT.

package change

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// GetChangeDuplicatesReader is a Reader for the GetChangeDuplicates structure.
type GetChangeDuplicatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChangeDuplicatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChangeDuplicatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetChangeDuplicatesOK creates a GetChangeDuplicatesOK with default headers values
func NewGetChangeDuplicatesOK() *GetChangeDuplicatesOK {
	return &GetChangeDuplicatesOK{}
}

/*GetChangeDuplicatesOK handles this case with default header values.

successful operation
*/
type GetChangeDuplicatesOK struct {
	Payload *models.Changes
}

func (o *GetChangeDuplicatesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/changes/{changeLocator}/duplicates][%d] getChangeDuplicatesOK  %+v", 200, o.Payload)
}

func (o *GetChangeDuplicatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Changes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
