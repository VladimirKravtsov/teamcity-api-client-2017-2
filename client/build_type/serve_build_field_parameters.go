// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

// NewServeBuildFieldParams creates a new ServeBuildFieldParams object
// with the default values initialized.
func NewServeBuildFieldParams() *ServeBuildFieldParams {
	var ()
	return &ServeBuildFieldParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildFieldParamsWithTimeout creates a new ServeBuildFieldParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildFieldParamsWithTimeout(timeout time.Duration) *ServeBuildFieldParams {
	var ()
	return &ServeBuildFieldParams{

		timeout: timeout,
	}
}

// NewServeBuildFieldParamsWithContext creates a new ServeBuildFieldParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildFieldParamsWithContext(ctx context.Context) *ServeBuildFieldParams {
	var ()
	return &ServeBuildFieldParams{

		Context: ctx,
	}
}

// NewServeBuildFieldParamsWithHTTPClient creates a new ServeBuildFieldParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildFieldParamsWithHTTPClient(client *http.Client) *ServeBuildFieldParams {
	var ()
	return &ServeBuildFieldParams{
		HTTPClient: client,
	}
}

/*ServeBuildFieldParams contains all the parameters to send to the API endpoint
for the serve build field operation typically these are written to a http.Request
*/
type ServeBuildFieldParams struct {

	/*BtLocator*/
	BtLocator string
	/*BuildLocator*/
	BuildLocator string
	/*Field*/
	Field string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build field params
func (o *ServeBuildFieldParams) WithTimeout(timeout time.Duration) *ServeBuildFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build field params
func (o *ServeBuildFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build field params
func (o *ServeBuildFieldParams) WithContext(ctx context.Context) *ServeBuildFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build field params
func (o *ServeBuildFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build field params
func (o *ServeBuildFieldParams) WithHTTPClient(client *http.Client) *ServeBuildFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build field params
func (o *ServeBuildFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the serve build field params
func (o *ServeBuildFieldParams) WithBtLocator(btLocator string) *ServeBuildFieldParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the serve build field params
func (o *ServeBuildFieldParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithBuildLocator adds the buildLocator to the serve build field params
func (o *ServeBuildFieldParams) WithBuildLocator(buildLocator string) *ServeBuildFieldParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the serve build field params
func (o *ServeBuildFieldParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithField adds the field to the serve build field params
func (o *ServeBuildFieldParams) WithField(field string) *ServeBuildFieldParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the serve build field params
func (o *ServeBuildFieldParams) SetField(field string) {
	o.Field = field
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
