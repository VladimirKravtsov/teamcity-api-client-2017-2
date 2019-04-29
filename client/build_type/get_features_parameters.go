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

// NewGetFeaturesParams creates a new GetFeaturesParams object
// with the default values initialized.
func NewGetFeaturesParams() *GetFeaturesParams {
	var ()
	return &GetFeaturesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFeaturesParamsWithTimeout creates a new GetFeaturesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFeaturesParamsWithTimeout(timeout time.Duration) *GetFeaturesParams {
	var ()
	return &GetFeaturesParams{

		timeout: timeout,
	}
}

// NewGetFeaturesParamsWithContext creates a new GetFeaturesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFeaturesParamsWithContext(ctx context.Context) *GetFeaturesParams {
	var ()
	return &GetFeaturesParams{

		Context: ctx,
	}
}

// NewGetFeaturesParamsWithHTTPClient creates a new GetFeaturesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFeaturesParamsWithHTTPClient(client *http.Client) *GetFeaturesParams {
	var ()
	return &GetFeaturesParams{
		HTTPClient: client,
	}
}

/*GetFeaturesParams contains all the parameters to send to the API endpoint
for the get features operation typically these are written to a http.Request
*/
type GetFeaturesParams struct {

	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get features params
func (o *GetFeaturesParams) WithTimeout(timeout time.Duration) *GetFeaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get features params
func (o *GetFeaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get features params
func (o *GetFeaturesParams) WithContext(ctx context.Context) *GetFeaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get features params
func (o *GetFeaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get features params
func (o *GetFeaturesParams) WithHTTPClient(client *http.Client) *GetFeaturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get features params
func (o *GetFeaturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get features params
func (o *GetFeaturesParams) WithBtLocator(btLocator string) *GetFeaturesParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get features params
func (o *GetFeaturesParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get features params
func (o *GetFeaturesParams) WithFields(fields *string) *GetFeaturesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get features params
func (o *GetFeaturesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetFeaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

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
