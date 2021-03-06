// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// PutAppRestUsersUserLocatorRolesRoleIDScopeReader is a Reader for the PutAppRestUsersUserLocatorRolesRoleIDScope structure.
type PutAppRestUsersUserLocatorRolesRoleIDScopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAppRestUsersUserLocatorRolesRoleIDScopeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAppRestUsersUserLocatorRolesRoleIDScopeOK creates a PutAppRestUsersUserLocatorRolesRoleIDScopeOK with default headers values
func NewPutAppRestUsersUserLocatorRolesRoleIDScopeOK() *PutAppRestUsersUserLocatorRolesRoleIDScopeOK {
	return &PutAppRestUsersUserLocatorRolesRoleIDScopeOK{}
}

/*PutAppRestUsersUserLocatorRolesRoleIDScopeOK handles this case with default header values.

successful operation
*/
type PutAppRestUsersUserLocatorRolesRoleIDScopeOK struct {
	Payload *models.Role
}

func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/users/{userLocator}/roles/{roleId}/{scope}][%d] putAppRestUsersUserLocatorRolesRoleIdScopeOK  %+v", 200, o.Payload)
}

func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
