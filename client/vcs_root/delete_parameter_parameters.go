// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

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
)

// NewDeleteParameterParams creates a new DeleteParameterParams object
// with the default values initialized.
func NewDeleteParameterParams() *DeleteParameterParams {
	var ()
	return &DeleteParameterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteParameterParamsWithTimeout creates a new DeleteParameterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteParameterParamsWithTimeout(timeout time.Duration) *DeleteParameterParams {
	var ()
	return &DeleteParameterParams{

		timeout: timeout,
	}
}

// NewDeleteParameterParamsWithContext creates a new DeleteParameterParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteParameterParamsWithContext(ctx context.Context) *DeleteParameterParams {
	var ()
	return &DeleteParameterParams{

		Context: ctx,
	}
}

// NewDeleteParameterParamsWithHTTPClient creates a new DeleteParameterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteParameterParamsWithHTTPClient(client *http.Client) *DeleteParameterParams {
	var ()
	return &DeleteParameterParams{
		HTTPClient: client,
	}
}

/*DeleteParameterParams contains all the parameters to send to the API endpoint
for the delete parameter operation typically these are written to a http.Request
*/
type DeleteParameterParams struct {

	/*Name*/
	Name string
	/*VcsRootLocator*/
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete parameter params
func (o *DeleteParameterParams) WithTimeout(timeout time.Duration) *DeleteParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete parameter params
func (o *DeleteParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete parameter params
func (o *DeleteParameterParams) WithContext(ctx context.Context) *DeleteParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete parameter params
func (o *DeleteParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete parameter params
func (o *DeleteParameterParams) WithHTTPClient(client *http.Client) *DeleteParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete parameter params
func (o *DeleteParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete parameter params
func (o *DeleteParameterParams) WithName(name string) *DeleteParameterParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete parameter params
func (o *DeleteParameterParams) SetName(name string) {
	o.Name = name
}

// WithVcsRootLocator adds the vcsRootLocator to the delete parameter params
func (o *DeleteParameterParams) WithVcsRootLocator(vcsRootLocator string) *DeleteParameterParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the delete parameter params
func (o *DeleteParameterParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param vcsRootLocator
	if err := r.SetPathParam("vcsRootLocator", o.VcsRootLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
