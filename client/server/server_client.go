// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddLicenseKeys add license keys API
*/
func (a *Client) AddLicenseKeys(params *AddLicenseKeysParams, authInfo runtime.ClientAuthInfoWriter) (*AddLicenseKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddLicenseKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addLicenseKeys",
		Method:             "POST",
		PathPattern:        "/app/rest/server/licensingData/licenseKeys",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddLicenseKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddLicenseKeysOK), nil

}

/*
DeleteLicenseKey delete license key API
*/
func (a *Client) DeleteLicenseKey(params *DeleteLicenseKeyParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLicenseKeyParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLicenseKey",
		Method:             "DELETE",
		PathPattern:        "/app/rest/server/licensingData/licenseKeys/{licenseKey}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLicenseKeyReader{formats: a.formats},
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
GetBackupStatus get backup status API
*/
func (a *Client) GetBackupStatus(params *GetBackupStatusParams, authInfo runtime.ClientAuthInfoWriter) (*GetBackupStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBackupStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBackupStatus",
		Method:             "GET",
		PathPattern:        "/app/rest/server/backup",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBackupStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBackupStatusOK), nil

}

/*
GetLicenseKey get license key API
*/
func (a *Client) GetLicenseKey(params *GetLicenseKeyParams, authInfo runtime.ClientAuthInfoWriter) (*GetLicenseKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLicenseKey",
		Method:             "GET",
		PathPattern:        "/app/rest/server/licensingData/licenseKeys/{licenseKey}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicenseKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLicenseKeyOK), nil

}

/*
GetLicenseKeys get license keys API
*/
func (a *Client) GetLicenseKeys(params *GetLicenseKeysParams, authInfo runtime.ClientAuthInfoWriter) (*GetLicenseKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicenseKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLicenseKeys",
		Method:             "GET",
		PathPattern:        "/app/rest/server/licensingData/licenseKeys",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicenseKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLicenseKeysOK), nil

}

/*
GetLicensingData get licensing data API
*/
func (a *Client) GetLicensingData(params *GetLicensingDataParams, authInfo runtime.ClientAuthInfoWriter) (*GetLicensingDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLicensingDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLicensingData",
		Method:             "GET",
		PathPattern:        "/app/rest/server/licensingData",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLicensingDataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLicensingDataOK), nil

}

/*
ServePlugins serve plugins API
*/
func (a *Client) ServePlugins(params *ServePluginsParams, authInfo runtime.ClientAuthInfoWriter) (*ServePluginsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServePluginsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "servePlugins",
		Method:             "GET",
		PathPattern:        "/app/rest/server/plugins",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServePluginsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServePluginsOK), nil

}

/*
ServeServerInfo serve server info API
*/
func (a *Client) ServeServerInfo(params *ServeServerInfoParams, authInfo runtime.ClientAuthInfoWriter) (*ServeServerInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeServerInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveServerInfo",
		Method:             "GET",
		PathPattern:        "/app/rest/server",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeServerInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeServerInfoOK), nil

}

/*
ServeServerVersion serve server version API
*/
func (a *Client) ServeServerVersion(params *ServeServerVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ServeServerVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeServerVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveServerVersion",
		Method:             "GET",
		PathPattern:        "/app/rest/server/{field}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeServerVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeServerVersionOK), nil

}

/*
StartBackup start backup API
*/
func (a *Client) StartBackup(params *StartBackupParams, authInfo runtime.ClientAuthInfoWriter) (*StartBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartBackupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startBackup",
		Method:             "POST",
		PathPattern:        "/app/rest/server/backup",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartBackupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartBackupOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
