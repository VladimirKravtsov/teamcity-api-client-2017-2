// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewServePluginInfoParams creates a new ServePluginInfoParams object
// with the default values initialized.
func NewServePluginInfoParams() *ServePluginInfoParams {
	var ()
	return &ServePluginInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServePluginInfoParamsWithTimeout creates a new ServePluginInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServePluginInfoParamsWithTimeout(timeout time.Duration) *ServePluginInfoParams {
	var ()
	return &ServePluginInfoParams{

		timeout: timeout,
	}
}

// NewServePluginInfoParamsWithContext creates a new ServePluginInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewServePluginInfoParamsWithContext(ctx context.Context) *ServePluginInfoParams {
	var ()
	return &ServePluginInfoParams{

		Context: ctx,
	}
}

// NewServePluginInfoParamsWithHTTPClient creates a new ServePluginInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServePluginInfoParamsWithHTTPClient(client *http.Client) *ServePluginInfoParams {
	var ()
	return &ServePluginInfoParams{
		HTTPClient: client,
	}
}

/*ServePluginInfoParams contains all the parameters to send to the API endpoint
for the serve plugin info operation typically these are written to a http.Request
*/
type ServePluginInfoParams struct {

	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve plugin info params
func (o *ServePluginInfoParams) WithTimeout(timeout time.Duration) *ServePluginInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve plugin info params
func (o *ServePluginInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve plugin info params
func (o *ServePluginInfoParams) WithContext(ctx context.Context) *ServePluginInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve plugin info params
func (o *ServePluginInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve plugin info params
func (o *ServePluginInfoParams) WithHTTPClient(client *http.Client) *ServePluginInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve plugin info params
func (o *ServePluginInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve plugin info params
func (o *ServePluginInfoParams) WithFields(fields *string) *ServePluginInfoParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve plugin info params
func (o *ServePluginInfoParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ServePluginInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
