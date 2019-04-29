// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// NewReplaceProjectsParams creates a new ReplaceProjectsParams object
// with the default values initialized.
func NewReplaceProjectsParams() *ReplaceProjectsParams {
	var ()
	return &ReplaceProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceProjectsParamsWithTimeout creates a new ReplaceProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceProjectsParamsWithTimeout(timeout time.Duration) *ReplaceProjectsParams {
	var ()
	return &ReplaceProjectsParams{

		timeout: timeout,
	}
}

// NewReplaceProjectsParamsWithContext creates a new ReplaceProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceProjectsParamsWithContext(ctx context.Context) *ReplaceProjectsParams {
	var ()
	return &ReplaceProjectsParams{

		Context: ctx,
	}
}

// NewReplaceProjectsParamsWithHTTPClient creates a new ReplaceProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceProjectsParamsWithHTTPClient(client *http.Client) *ReplaceProjectsParams {
	var ()
	return &ReplaceProjectsParams{
		HTTPClient: client,
	}
}

/*ReplaceProjectsParams contains all the parameters to send to the API endpoint
for the replace projects operation typically these are written to a http.Request
*/
type ReplaceProjectsParams struct {

	/*AgentPoolLocator*/
	AgentPoolLocator string
	/*Body*/
	Body *models.Projects

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace projects params
func (o *ReplaceProjectsParams) WithTimeout(timeout time.Duration) *ReplaceProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace projects params
func (o *ReplaceProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace projects params
func (o *ReplaceProjectsParams) WithContext(ctx context.Context) *ReplaceProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace projects params
func (o *ReplaceProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace projects params
func (o *ReplaceProjectsParams) WithHTTPClient(client *http.Client) *ReplaceProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace projects params
func (o *ReplaceProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentPoolLocator adds the agentPoolLocator to the replace projects params
func (o *ReplaceProjectsParams) WithAgentPoolLocator(agentPoolLocator string) *ReplaceProjectsParams {
	o.SetAgentPoolLocator(agentPoolLocator)
	return o
}

// SetAgentPoolLocator adds the agentPoolLocator to the replace projects params
func (o *ReplaceProjectsParams) SetAgentPoolLocator(agentPoolLocator string) {
	o.AgentPoolLocator = agentPoolLocator
}

// WithBody adds the body to the replace projects params
func (o *ReplaceProjectsParams) WithBody(body *models.Projects) *ReplaceProjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace projects params
func (o *ReplaceProjectsParams) SetBody(body *models.Projects) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentPoolLocator
	if err := r.SetPathParam("agentPoolLocator", o.AgentPoolLocator); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
