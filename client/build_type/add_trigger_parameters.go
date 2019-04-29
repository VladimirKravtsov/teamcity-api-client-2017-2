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

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// NewAddTriggerParams creates a new AddTriggerParams object
// with the default values initialized.
func NewAddTriggerParams() *AddTriggerParams {
	var ()
	return &AddTriggerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddTriggerParamsWithTimeout creates a new AddTriggerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddTriggerParamsWithTimeout(timeout time.Duration) *AddTriggerParams {
	var ()
	return &AddTriggerParams{

		timeout: timeout,
	}
}

// NewAddTriggerParamsWithContext creates a new AddTriggerParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddTriggerParamsWithContext(ctx context.Context) *AddTriggerParams {
	var ()
	return &AddTriggerParams{

		Context: ctx,
	}
}

// NewAddTriggerParamsWithHTTPClient creates a new AddTriggerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddTriggerParamsWithHTTPClient(client *http.Client) *AddTriggerParams {
	var ()
	return &AddTriggerParams{
		HTTPClient: client,
	}
}

/*AddTriggerParams contains all the parameters to send to the API endpoint
for the add trigger operation typically these are written to a http.Request
*/
type AddTriggerParams struct {

	/*Body*/
	Body *models.Trigger
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add trigger params
func (o *AddTriggerParams) WithTimeout(timeout time.Duration) *AddTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add trigger params
func (o *AddTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add trigger params
func (o *AddTriggerParams) WithContext(ctx context.Context) *AddTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add trigger params
func (o *AddTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add trigger params
func (o *AddTriggerParams) WithHTTPClient(client *http.Client) *AddTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add trigger params
func (o *AddTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add trigger params
func (o *AddTriggerParams) WithBody(body *models.Trigger) *AddTriggerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add trigger params
func (o *AddTriggerParams) SetBody(body *models.Trigger) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the add trigger params
func (o *AddTriggerParams) WithBtLocator(btLocator string) *AddTriggerParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the add trigger params
func (o *AddTriggerParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the add trigger params
func (o *AddTriggerParams) WithFields(fields *string) *AddTriggerParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add trigger params
func (o *AddTriggerParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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