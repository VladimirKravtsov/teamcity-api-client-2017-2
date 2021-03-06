// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new agent pool API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for agent pool API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddAgent add agent API
*/
func (a *Client) AddAgent(params *AddAgentParams, authInfo runtime.ClientAuthInfoWriter) (*AddAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addAgent",
		Method:             "POST",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}/agents",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddAgentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddAgentOK), nil

}

/*
AddProject add project API
*/
func (a *Client) AddProject(params *AddProjectParams, authInfo runtime.ClientAuthInfoWriter) (*AddProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addProject",
		Method:             "POST",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}/projects",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddProjectOK), nil

}

/*
CreatePool create pool API
*/
func (a *Client) CreatePool(params *CreatePoolParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPool",
		Method:             "POST",
		PathPattern:        "/app/rest/agentPools",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePoolOK), nil

}

/*
DeletePool delete pool API
*/
func (a *Client) DeletePool(params *DeletePoolParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePoolParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePool",
		Method:             "DELETE",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePoolReader{formats: a.formats},
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
DeletePoolProject delete pool project API
*/
func (a *Client) DeletePoolProject(params *DeletePoolProjectParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePoolProjectParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePoolProject",
		Method:             "DELETE",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}/projects/{projectLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePoolProjectReader{formats: a.formats},
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
DeleteProjects delete projects API
*/
func (a *Client) DeleteProjects(params *DeleteProjectsParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectsParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProjects",
		Method:             "DELETE",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}/projects",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProjectsReader{formats: a.formats},
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
GetField get field API
*/
func (a *Client) GetField(params *GetFieldParams, authInfo runtime.ClientAuthInfoWriter) (*GetFieldOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFieldParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getField",
		Method:             "GET",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}/{field}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFieldReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFieldOK), nil

}

/*
GetPool get pool API
*/
func (a *Client) GetPool(params *GetPoolParams, authInfo runtime.ClientAuthInfoWriter) (*GetPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoolParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPool",
		Method:             "GET",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPoolOK), nil

}

/*
GetPoolAgents get pool agents API
*/
func (a *Client) GetPoolAgents(params *GetPoolAgentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetPoolAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoolAgentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPoolAgents",
		Method:             "GET",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}/agents",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoolAgentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPoolAgentsOK), nil

}

/*
GetPoolProject get pool project API
*/
func (a *Client) GetPoolProject(params *GetPoolProjectParams, authInfo runtime.ClientAuthInfoWriter) (*GetPoolProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoolProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPoolProject",
		Method:             "GET",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}/projects/{projectLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoolProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPoolProjectOK), nil

}

/*
GetPoolProjects get pool projects API
*/
func (a *Client) GetPoolProjects(params *GetPoolProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*GetPoolProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoolProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPoolProjects",
		Method:             "GET",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}/projects",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoolProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPoolProjectsOK), nil

}

/*
GetPools get pools API
*/
func (a *Client) GetPools(params *GetPoolsParams, authInfo runtime.ClientAuthInfoWriter) (*GetPoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoolsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPools",
		Method:             "GET",
		PathPattern:        "/app/rest/agentPools",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPoolsOK), nil

}

/*
ReplaceProjects replace projects API
*/
func (a *Client) ReplaceProjects(params *ReplaceProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceProjects",
		Method:             "PUT",
		PathPattern:        "/app/rest/agentPools/{agentPoolLocator}/projects",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReplaceProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReplaceProjectsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
