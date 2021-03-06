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

// NewAddStepParams creates a new AddStepParams object
// with the default values initialized.
func NewAddStepParams() *AddStepParams {
	var ()
	return &AddStepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddStepParamsWithTimeout creates a new AddStepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddStepParamsWithTimeout(timeout time.Duration) *AddStepParams {
	var ()
	return &AddStepParams{

		timeout: timeout,
	}
}

// NewAddStepParamsWithContext creates a new AddStepParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddStepParamsWithContext(ctx context.Context) *AddStepParams {
	var ()
	return &AddStepParams{

		Context: ctx,
	}
}

// NewAddStepParamsWithHTTPClient creates a new AddStepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddStepParamsWithHTTPClient(client *http.Client) *AddStepParams {
	var ()
	return &AddStepParams{
		HTTPClient: client,
	}
}

/*AddStepParams contains all the parameters to send to the API endpoint
for the add step operation typically these are written to a http.Request
*/
type AddStepParams struct {

	/*Body*/
	Body *models.Step
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add step params
func (o *AddStepParams) WithTimeout(timeout time.Duration) *AddStepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add step params
func (o *AddStepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add step params
func (o *AddStepParams) WithContext(ctx context.Context) *AddStepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add step params
func (o *AddStepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add step params
func (o *AddStepParams) WithHTTPClient(client *http.Client) *AddStepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add step params
func (o *AddStepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add step params
func (o *AddStepParams) WithBody(body *models.Step) *AddStepParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add step params
func (o *AddStepParams) SetBody(body *models.Step) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the add step params
func (o *AddStepParams) WithBtLocator(btLocator string) *AddStepParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the add step params
func (o *AddStepParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the add step params
func (o *AddStepParams) WithFields(fields *string) *AddStepParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add step params
func (o *AddStepParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddStepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
