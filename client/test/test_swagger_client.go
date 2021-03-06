// Code generated by go-swagger; DO NOT EDIT.

package test

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new test API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for test API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetTests get tests API
*/
func (a *Client) GetTests(params *GetTestsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTestsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTests",
		Method:             "GET",
		PathPattern:        "/app/rest/tests",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "application/xml", "text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTestsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTestsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
