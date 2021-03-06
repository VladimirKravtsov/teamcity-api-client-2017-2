// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewServeUserParams creates a new ServeUserParams object
// with the default values initialized.
func NewServeUserParams() *ServeUserParams {
	var ()
	return &ServeUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeUserParamsWithTimeout creates a new ServeUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeUserParamsWithTimeout(timeout time.Duration) *ServeUserParams {
	var ()
	return &ServeUserParams{

		timeout: timeout,
	}
}

// NewServeUserParamsWithContext creates a new ServeUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeUserParamsWithContext(ctx context.Context) *ServeUserParams {
	var ()
	return &ServeUserParams{

		Context: ctx,
	}
}

// NewServeUserParamsWithHTTPClient creates a new ServeUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeUserParamsWithHTTPClient(client *http.Client) *ServeUserParams {
	var ()
	return &ServeUserParams{
		HTTPClient: client,
	}
}

/*ServeUserParams contains all the parameters to send to the API endpoint
for the serve user operation typically these are written to a http.Request
*/
type ServeUserParams struct {

	/*Fields*/
	Fields *string
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve user params
func (o *ServeUserParams) WithTimeout(timeout time.Duration) *ServeUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve user params
func (o *ServeUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve user params
func (o *ServeUserParams) WithContext(ctx context.Context) *ServeUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve user params
func (o *ServeUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve user params
func (o *ServeUserParams) WithHTTPClient(client *http.Client) *ServeUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve user params
func (o *ServeUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve user params
func (o *ServeUserParams) WithFields(fields *string) *ServeUserParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve user params
func (o *ServeUserParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserLocator adds the userLocator to the serve user params
func (o *ServeUserParams) WithUserLocator(userLocator string) *ServeUserParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the serve user params
func (o *ServeUserParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userLocator
	if err := r.SetPathParam("userLocator", o.UserLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
