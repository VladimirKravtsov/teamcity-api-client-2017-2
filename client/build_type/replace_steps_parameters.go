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

// NewReplaceStepsParams creates a new ReplaceStepsParams object
// with the default values initialized.
func NewReplaceStepsParams() *ReplaceStepsParams {
	var ()
	return &ReplaceStepsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceStepsParamsWithTimeout creates a new ReplaceStepsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceStepsParamsWithTimeout(timeout time.Duration) *ReplaceStepsParams {
	var ()
	return &ReplaceStepsParams{

		timeout: timeout,
	}
}

// NewReplaceStepsParamsWithContext creates a new ReplaceStepsParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceStepsParamsWithContext(ctx context.Context) *ReplaceStepsParams {
	var ()
	return &ReplaceStepsParams{

		Context: ctx,
	}
}

// NewReplaceStepsParamsWithHTTPClient creates a new ReplaceStepsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceStepsParamsWithHTTPClient(client *http.Client) *ReplaceStepsParams {
	var ()
	return &ReplaceStepsParams{
		HTTPClient: client,
	}
}

/*ReplaceStepsParams contains all the parameters to send to the API endpoint
for the replace steps operation typically these are written to a http.Request
*/
type ReplaceStepsParams struct {

	/*Body*/
	Body *models.Steps
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace steps params
func (o *ReplaceStepsParams) WithTimeout(timeout time.Duration) *ReplaceStepsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace steps params
func (o *ReplaceStepsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace steps params
func (o *ReplaceStepsParams) WithContext(ctx context.Context) *ReplaceStepsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace steps params
func (o *ReplaceStepsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace steps params
func (o *ReplaceStepsParams) WithHTTPClient(client *http.Client) *ReplaceStepsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace steps params
func (o *ReplaceStepsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace steps params
func (o *ReplaceStepsParams) WithBody(body *models.Steps) *ReplaceStepsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace steps params
func (o *ReplaceStepsParams) SetBody(body *models.Steps) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the replace steps params
func (o *ReplaceStepsParams) WithBtLocator(btLocator string) *ReplaceStepsParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the replace steps params
func (o *ReplaceStepsParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the replace steps params
func (o *ReplaceStepsParams) WithFields(fields *string) *ReplaceStepsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace steps params
func (o *ReplaceStepsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceStepsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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