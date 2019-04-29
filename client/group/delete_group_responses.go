// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteGroupReader is a Reader for the DeleteGroup structure.
type DeleteGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteGroupDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteGroupDefault creates a DeleteGroupDefault with default headers values
func NewDeleteGroupDefault(code int) *DeleteGroupDefault {
	return &DeleteGroupDefault{
		_statusCode: code,
	}
}

/*DeleteGroupDefault handles this case with default header values.

successful operation
*/
type DeleteGroupDefault struct {
	_statusCode int
}

// Code gets the status code for the delete group default response
func (o *DeleteGroupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteGroupDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/userGroups/{groupLocator}][%d] deleteGroup default ", o._statusCode)
}

func (o *DeleteGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}