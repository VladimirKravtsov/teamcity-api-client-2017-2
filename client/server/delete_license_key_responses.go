// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteLicenseKeyReader is a Reader for the DeleteLicenseKey structure.
type DeleteLicenseKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLicenseKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteLicenseKeyDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteLicenseKeyDefault creates a DeleteLicenseKeyDefault with default headers values
func NewDeleteLicenseKeyDefault(code int) *DeleteLicenseKeyDefault {
	return &DeleteLicenseKeyDefault{
		_statusCode: code,
	}
}

/*DeleteLicenseKeyDefault handles this case with default header values.

successful operation
*/
type DeleteLicenseKeyDefault struct {
	_statusCode int
}

// Code gets the status code for the delete license key default response
func (o *DeleteLicenseKeyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLicenseKeyDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/server/licensingData/licenseKeys/{licenseKey}][%d] deleteLicenseKey default ", o._statusCode)
}

func (o *DeleteLicenseKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
