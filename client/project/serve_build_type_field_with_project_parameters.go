// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewServeBuildTypeFieldWithProjectParams creates a new ServeBuildTypeFieldWithProjectParams object
// with the default values initialized.
func NewServeBuildTypeFieldWithProjectParams() *ServeBuildTypeFieldWithProjectParams {
	var ()
	return &ServeBuildTypeFieldWithProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildTypeFieldWithProjectParamsWithTimeout creates a new ServeBuildTypeFieldWithProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildTypeFieldWithProjectParamsWithTimeout(timeout time.Duration) *ServeBuildTypeFieldWithProjectParams {
	var ()
	return &ServeBuildTypeFieldWithProjectParams{

		timeout: timeout,
	}
}

// NewServeBuildTypeFieldWithProjectParamsWithContext creates a new ServeBuildTypeFieldWithProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildTypeFieldWithProjectParamsWithContext(ctx context.Context) *ServeBuildTypeFieldWithProjectParams {
	var ()
	return &ServeBuildTypeFieldWithProjectParams{

		Context: ctx,
	}
}

// NewServeBuildTypeFieldWithProjectParamsWithHTTPClient creates a new ServeBuildTypeFieldWithProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildTypeFieldWithProjectParamsWithHTTPClient(client *http.Client) *ServeBuildTypeFieldWithProjectParams {
	var ()
	return &ServeBuildTypeFieldWithProjectParams{
		HTTPClient: client,
	}
}

/*ServeBuildTypeFieldWithProjectParams contains all the parameters to send to the API endpoint
for the serve build type field with project operation typically these are written to a http.Request
*/
type ServeBuildTypeFieldWithProjectParams struct {

	/*BtLocator*/
	BtLocator string
	/*Field*/
	Field string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) WithTimeout(timeout time.Duration) *ServeBuildTypeFieldWithProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) WithContext(ctx context.Context) *ServeBuildTypeFieldWithProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) WithHTTPClient(client *http.Client) *ServeBuildTypeFieldWithProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) WithBtLocator(btLocator string) *ServeBuildTypeFieldWithProjectParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithField adds the field to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) WithField(field string) *ServeBuildTypeFieldWithProjectParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) SetField(field string) {
	o.Field = field
}

// WithProjectLocator adds the projectLocator to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) WithProjectLocator(projectLocator string) *ServeBuildTypeFieldWithProjectParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the serve build type field with project params
func (o *ServeBuildTypeFieldWithProjectParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildTypeFieldWithProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
	}

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
