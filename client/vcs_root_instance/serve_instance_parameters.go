// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

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

// NewServeInstanceParams creates a new ServeInstanceParams object
// with the default values initialized.
func NewServeInstanceParams() *ServeInstanceParams {
	var ()
	return &ServeInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeInstanceParamsWithTimeout creates a new ServeInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeInstanceParamsWithTimeout(timeout time.Duration) *ServeInstanceParams {
	var ()
	return &ServeInstanceParams{

		timeout: timeout,
	}
}

// NewServeInstanceParamsWithContext creates a new ServeInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeInstanceParamsWithContext(ctx context.Context) *ServeInstanceParams {
	var ()
	return &ServeInstanceParams{

		Context: ctx,
	}
}

// NewServeInstanceParamsWithHTTPClient creates a new ServeInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeInstanceParamsWithHTTPClient(client *http.Client) *ServeInstanceParams {
	var ()
	return &ServeInstanceParams{
		HTTPClient: client,
	}
}

/*ServeInstanceParams contains all the parameters to send to the API endpoint
for the serve instance operation typically these are written to a http.Request
*/
type ServeInstanceParams struct {

	/*Fields*/
	Fields *string
	/*VcsRootInstanceLocator*/
	VcsRootInstanceLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve instance params
func (o *ServeInstanceParams) WithTimeout(timeout time.Duration) *ServeInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve instance params
func (o *ServeInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve instance params
func (o *ServeInstanceParams) WithContext(ctx context.Context) *ServeInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve instance params
func (o *ServeInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve instance params
func (o *ServeInstanceParams) WithHTTPClient(client *http.Client) *ServeInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve instance params
func (o *ServeInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve instance params
func (o *ServeInstanceParams) WithFields(fields *string) *ServeInstanceParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve instance params
func (o *ServeInstanceParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithVcsRootInstanceLocator adds the vcsRootInstanceLocator to the serve instance params
func (o *ServeInstanceParams) WithVcsRootInstanceLocator(vcsRootInstanceLocator string) *ServeInstanceParams {
	o.SetVcsRootInstanceLocator(vcsRootInstanceLocator)
	return o
}

// SetVcsRootInstanceLocator adds the vcsRootInstanceLocator to the serve instance params
func (o *ServeInstanceParams) SetVcsRootInstanceLocator(vcsRootInstanceLocator string) {
	o.VcsRootInstanceLocator = vcsRootInstanceLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param vcsRootInstanceLocator
	if err := r.SetPathParam("vcsRootInstanceLocator", o.VcsRootInstanceLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
