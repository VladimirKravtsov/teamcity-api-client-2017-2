// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PutAppRestUsersUserLocatorRolesRoleIDScope put app rest users user locator roles role ID scope API
*/
func (a *Client) PutAppRestUsersUserLocatorRolesRoleIDScope(params *PutAppRestUsersUserLocatorRolesRoleIDScopeParams, authInfo runtime.ClientAuthInfoWriter) (*PutAppRestUsersUserLocatorRolesRoleIDScopeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAppRestUsersUserLocatorRolesRoleIDScopeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutAppRestUsersUserLocatorRolesRoleIDScope",
		Method:             "PUT",
		PathPattern:        "/app/rest/users/{userLocator}/roles/{roleId}/{scope}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutAppRestUsersUserLocatorRolesRoleIDScopeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutAppRestUsersUserLocatorRolesRoleIDScopeOK), nil

}

/*
AddGroup add group API
*/
func (a *Client) AddGroup(params *AddGroupParams, authInfo runtime.ClientAuthInfoWriter) (*AddGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addGroup",
		Method:             "POST",
		PathPattern:        "/app/rest/users/{userLocator}/groups",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddGroupOK), nil

}

/*
AddRole add role API
*/
func (a *Client) AddRole(params *AddRoleParams, authInfo runtime.ClientAuthInfoWriter) (*AddRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addRole",
		Method:             "POST",
		PathPattern:        "/app/rest/users/{userLocator}/roles",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddRoleOK), nil

}

/*
AddRoleSimplePost add role simple post API
*/
func (a *Client) AddRoleSimplePost(params *AddRoleSimplePostParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRoleSimplePostParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addRoleSimplePost",
		Method:             "POST",
		PathPattern:        "/app/rest/users/{userLocator}/roles/{roleId}/{scope}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddRoleSimplePostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
CreateUser create user API
*/
func (a *Client) CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUser",
		Method:             "POST",
		PathPattern:        "/app/rest/users",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserOK), nil

}

/*
DeleteRememberMe delete remember me API
*/
func (a *Client) DeleteRememberMe(params *DeleteRememberMeParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRememberMeParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRememberMe",
		Method:             "DELETE",
		PathPattern:        "/app/rest/users/{userLocator}/debug/rememberMe",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRememberMeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteRole delete role API
*/
func (a *Client) DeleteRole(params *DeleteRoleParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRoleParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRole",
		Method:             "DELETE",
		PathPattern:        "/app/rest/users/{userLocator}/roles/{roleId}/{scope}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteUser delete user API
*/
func (a *Client) DeleteUser(params *DeleteUserParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUser",
		Method:             "DELETE",
		PathPattern:        "/app/rest/users/{userLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteUserField delete user field API
*/
func (a *Client) DeleteUserField(params *DeleteUserFieldParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserFieldParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUserField",
		Method:             "DELETE",
		PathPattern:        "/app/rest/users/{userLocator}/{field}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserFieldReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetGroup get group API
*/
func (a *Client) GetGroup(params *GetGroupParams, authInfo runtime.ClientAuthInfoWriter) (*GetGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGroup",
		Method:             "GET",
		PathPattern:        "/app/rest/users/{userLocator}/groups/{groupLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGroupOK), nil

}

/*
GetGroups get groups API
*/
func (a *Client) GetGroups(params *GetGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGroups",
		Method:             "GET",
		PathPattern:        "/app/rest/users/{userLocator}/groups",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGroupsOK), nil

}

/*
GetPermissions get permissions API
*/
func (a *Client) GetPermissions(params *GetPermissionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPermissionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPermissions",
		Method:             "GET",
		PathPattern:        "/app/rest/users/{userLocator}/permissions",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPermissionsOK), nil

}

/*
ListRole list role API
*/
func (a *Client) ListRole(params *ListRoleParams, authInfo runtime.ClientAuthInfoWriter) (*ListRoleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRoleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRole",
		Method:             "GET",
		PathPattern:        "/app/rest/users/{userLocator}/roles/{roleId}/{scope}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRoleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRoleOK), nil

}

/*
ListRoles list roles API
*/
func (a *Client) ListRoles(params *ListRolesParams, authInfo runtime.ClientAuthInfoWriter) (*ListRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRoles",
		Method:             "GET",
		PathPattern:        "/app/rest/users/{userLocator}/roles",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRolesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRolesOK), nil

}

/*
PutUserProperty put user property API
*/
func (a *Client) PutUserProperty(params *PutUserPropertyParams, authInfo runtime.ClientAuthInfoWriter) (*PutUserPropertyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUserPropertyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putUserProperty",
		Method:             "PUT",
		PathPattern:        "/app/rest/users/{userLocator}/properties/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutUserPropertyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutUserPropertyOK), nil

}

/*
RemoveGroup remove group API
*/
func (a *Client) RemoveGroup(params *RemoveGroupParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveGroupParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeGroup",
		Method:             "DELETE",
		PathPattern:        "/app/rest/users/{userLocator}/groups/{groupLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RemoveUserProperty remove user property API
*/
func (a *Client) RemoveUserProperty(params *RemoveUserPropertyParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveUserPropertyParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeUserProperty",
		Method:             "DELETE",
		PathPattern:        "/app/rest/users/{userLocator}/properties/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveUserPropertyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
ReplaceGroups replace groups API
*/
func (a *Client) ReplaceGroups(params *ReplaceGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceGroups",
		Method:             "PUT",
		PathPattern:        "/app/rest/users/{userLocator}/groups",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReplaceGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReplaceGroupsOK), nil

}

/*
ReplaceRoles replace roles API
*/
func (a *Client) ReplaceRoles(params *ReplaceRolesParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceRoles",
		Method:             "PUT",
		PathPattern:        "/app/rest/users/{userLocator}/roles",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReplaceRolesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReplaceRolesOK), nil

}

/*
ServeUser serve user API
*/
func (a *Client) ServeUser(params *ServeUserParams, authInfo runtime.ClientAuthInfoWriter) (*ServeUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveUser",
		Method:             "GET",
		PathPattern:        "/app/rest/users/{userLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeUserOK), nil

}

/*
ServeUserField serve user field API
*/
func (a *Client) ServeUserField(params *ServeUserFieldParams, authInfo runtime.ClientAuthInfoWriter) (*ServeUserFieldOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeUserFieldParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveUserField",
		Method:             "GET",
		PathPattern:        "/app/rest/users/{userLocator}/{field}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeUserFieldReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeUserFieldOK), nil

}

/*
ServeUserProperties serve user properties API
*/
func (a *Client) ServeUserProperties(params *ServeUserPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*ServeUserPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeUserPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveUserProperties",
		Method:             "GET",
		PathPattern:        "/app/rest/users/{userLocator}/properties",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeUserPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeUserPropertiesOK), nil

}

/*
ServeUserProperty serve user property API
*/
func (a *Client) ServeUserProperty(params *ServeUserPropertyParams, authInfo runtime.ClientAuthInfoWriter) (*ServeUserPropertyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeUserPropertyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveUserProperty",
		Method:             "GET",
		PathPattern:        "/app/rest/users/{userLocator}/properties/{name}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeUserPropertyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeUserPropertyOK), nil

}

/*
ServeUsers serve users API
*/
func (a *Client) ServeUsers(params *ServeUsersParams, authInfo runtime.ClientAuthInfoWriter) (*ServeUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveUsers",
		Method:             "GET",
		PathPattern:        "/app/rest/users",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeUsersOK), nil

}

/*
SetUserField set user field API
*/
func (a *Client) SetUserField(params *SetUserFieldParams, authInfo runtime.ClientAuthInfoWriter) (*SetUserFieldOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetUserFieldParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setUserField",
		Method:             "PUT",
		PathPattern:        "/app/rest/users/{userLocator}/{field}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetUserFieldReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetUserFieldOK), nil

}

/*
UpdateUser update user API
*/
func (a *Client) UpdateUser(params *UpdateUserParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUser",
		Method:             "PUT",
		PathPattern:        "/app/rest/users/{userLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateUserOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
