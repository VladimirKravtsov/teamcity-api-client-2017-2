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

// NewSetProjectFieldParams creates a new SetProjectFieldParams object
// with the default values initialized.
func NewSetProjectFieldParams() *SetProjectFieldParams {
	var ()
	return &SetProjectFieldParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetProjectFieldParamsWithTimeout creates a new SetProjectFieldParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetProjectFieldParamsWithTimeout(timeout time.Duration) *SetProjectFieldParams {
	var ()
	return &SetProjectFieldParams{

		timeout: timeout,
	}
}

// NewSetProjectFieldParamsWithContext creates a new SetProjectFieldParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetProjectFieldParamsWithContext(ctx context.Context) *SetProjectFieldParams {
	var ()
	return &SetProjectFieldParams{

		Context: ctx,
	}
}

// NewSetProjectFieldParamsWithHTTPClient creates a new SetProjectFieldParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetProjectFieldParamsWithHTTPClient(client *http.Client) *SetProjectFieldParams {
	var ()
	return &SetProjectFieldParams{
		HTTPClient: client,
	}
}

/*SetProjectFieldParams contains all the parameters to send to the API endpoint
for the set project field operation typically these are written to a http.Request
*/
type SetProjectFieldParams struct {

	/*Body*/
	Body string
	/*Field*/
	Field string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set project field params
func (o *SetProjectFieldParams) WithTimeout(timeout time.Duration) *SetProjectFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set project field params
func (o *SetProjectFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set project field params
func (o *SetProjectFieldParams) WithContext(ctx context.Context) *SetProjectFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set project field params
func (o *SetProjectFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set project field params
func (o *SetProjectFieldParams) WithHTTPClient(client *http.Client) *SetProjectFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set project field params
func (o *SetProjectFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set project field params
func (o *SetProjectFieldParams) WithBody(body string) *SetProjectFieldParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set project field params
func (o *SetProjectFieldParams) SetBody(body string) {
	o.Body = body
}

// WithField adds the field to the set project field params
func (o *SetProjectFieldParams) WithField(field string) *SetProjectFieldParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the set project field params
func (o *SetProjectFieldParams) SetField(field string) {
	o.Field = field
}

// WithProjectLocator adds the projectLocator to the set project field params
func (o *SetProjectFieldParams) WithProjectLocator(projectLocator string) *SetProjectFieldParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the set project field params
func (o *SetProjectFieldParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetProjectFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
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
