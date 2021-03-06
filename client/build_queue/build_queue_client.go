// Code generated by go-swagger; DO NOT EDIT.

package build_queue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new build queue API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for build queue API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostAppRestBuildQueueQueuedBuildLocator post app rest build queue queued build locator API
*/
func (a *Client) PostAppRestBuildQueueQueuedBuildLocator(params *PostAppRestBuildQueueQueuedBuildLocatorParams, authInfo runtime.ClientAuthInfoWriter) (*PostAppRestBuildQueueQueuedBuildLocatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAppRestBuildQueueQueuedBuildLocatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAppRestBuildQueueQueuedBuildLocator",
		Method:             "POST",
		PathPattern:        "/app/rest/buildQueue/{queuedBuildLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAppRestBuildQueueQueuedBuildLocatorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAppRestBuildQueueQueuedBuildLocatorOK), nil

}

/*
DeleteBuildsExperimental delete builds experimental API
*/
func (a *Client) DeleteBuildsExperimental(params *DeleteBuildsExperimentalParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBuildsExperimentalParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBuildsExperimental",
		Method:             "DELETE",
		PathPattern:        "/app/rest/buildQueue",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBuildsExperimentalReader{formats: a.formats},
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
GetBuild get build API
*/
func (a *Client) GetBuild(params *GetBuildParams, authInfo runtime.ClientAuthInfoWriter) (*GetBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBuild",
		Method:             "GET",
		PathPattern:        "/app/rest/buildQueue/{queuedBuildLocator}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBuildOK), nil

}

/*
GetBuilds get builds API
*/
func (a *Client) GetBuilds(params *GetBuildsParams, authInfo runtime.ClientAuthInfoWriter) (*GetBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBuilds",
		Method:             "GET",
		PathPattern:        "/app/rest/buildQueue",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBuildsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBuildsOK), nil

}

/*
QueueNewBuild queue new build API
*/
func (a *Client) QueueNewBuild(params *QueueNewBuildParams, authInfo runtime.ClientAuthInfoWriter) (*QueueNewBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueueNewBuildParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queueNewBuild",
		Method:             "POST",
		PathPattern:        "/app/rest/buildQueue",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueueNewBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueueNewBuildOK), nil

}

/*
ReplaceBuilds replace builds API
*/
func (a *Client) ReplaceBuilds(params *ReplaceBuildsParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceBuildsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceBuildsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceBuilds",
		Method:             "PUT",
		PathPattern:        "/app/rest/buildQueue",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReplaceBuildsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ReplaceBuildsOK), nil

}

/*
ServeCompatibleAgents serve compatible agents API
*/
func (a *Client) ServeCompatibleAgents(params *ServeCompatibleAgentsParams, authInfo runtime.ClientAuthInfoWriter) (*ServeCompatibleAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeCompatibleAgentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveCompatibleAgents",
		Method:             "GET",
		PathPattern:        "/app/rest/buildQueue/{queuedBuildLocator}/compatibleAgents",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeCompatibleAgentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServeCompatibleAgentsOK), nil

}

/*
SetBuildQueueOrder set build queue order API
*/
func (a *Client) SetBuildQueueOrder(params *SetBuildQueueOrderParams, authInfo runtime.ClientAuthInfoWriter) (*SetBuildQueueOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetBuildQueueOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setBuildQueueOrder",
		Method:             "PUT",
		PathPattern:        "/app/rest/buildQueue/order",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetBuildQueueOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetBuildQueueOrderOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
