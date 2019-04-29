// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetScrambledReader is a Reader for the GetScrambled structure.
type GetScrambledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScrambledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetScrambledOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScrambledOK creates a GetScrambledOK with default headers values
func NewGetScrambledOK() *GetScrambledOK {
	return &GetScrambledOK{}
}

/*GetScrambledOK handles this case with default header values.

successful operation
*/
type GetScrambledOK struct {
	Payload string
}

func (o *GetScrambledOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/debug/values/password/scrambled][%d] getScrambledOK  %+v", 200, o.Payload)
}

func (o *GetScrambledOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
