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

// NewAddStepParameterParams creates a new AddStepParameterParams object
// with the default values initialized.
func NewAddStepParameterParams() *AddStepParameterParams {
	var ()
	return &AddStepParameterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddStepParameterParamsWithTimeout creates a new AddStepParameterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddStepParameterParamsWithTimeout(timeout time.Duration) *AddStepParameterParams {
	var ()
	return &AddStepParameterParams{

		timeout: timeout,
	}
}

// NewAddStepParameterParamsWithContext creates a new AddStepParameterParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddStepParameterParamsWithContext(ctx context.Context) *AddStepParameterParams {
	var ()
	return &AddStepParameterParams{

		Context: ctx,
	}
}

// NewAddStepParameterParamsWithHTTPClient creates a new AddStepParameterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddStepParameterParamsWithHTTPClient(client *http.Client) *AddStepParameterParams {
	var ()
	return &AddStepParameterParams{
		HTTPClient: client,
	}
}

/*AddStepParameterParams contains all the parameters to send to the API endpoint
for the add step parameter operation typically these are written to a http.Request
*/
type AddStepParameterParams struct {

	/*Body*/
	Body string
	/*BtLocator*/
	BtLocator string
	/*ParameterName*/
	ParameterName string
	/*StepID*/
	StepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add step parameter params
func (o *AddStepParameterParams) WithTimeout(timeout time.Duration) *AddStepParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add step parameter params
func (o *AddStepParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add step parameter params
func (o *AddStepParameterParams) WithContext(ctx context.Context) *AddStepParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add step parameter params
func (o *AddStepParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add step parameter params
func (o *AddStepParameterParams) WithHTTPClient(client *http.Client) *AddStepParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add step parameter params
func (o *AddStepParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add step parameter params
func (o *AddStepParameterParams) WithBody(body string) *AddStepParameterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add step parameter params
func (o *AddStepParameterParams) SetBody(body string) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the add step parameter params
func (o *AddStepParameterParams) WithBtLocator(btLocator string) *AddStepParameterParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the add step parameter params
func (o *AddStepParameterParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithParameterName adds the parameterName to the add step parameter params
func (o *AddStepParameterParams) WithParameterName(parameterName string) *AddStepParameterParams {
	o.SetParameterName(parameterName)
	return o
}

// SetParameterName adds the parameterName to the add step parameter params
func (o *AddStepParameterParams) SetParameterName(parameterName string) {
	o.ParameterName = parameterName
}

// WithStepID adds the stepID to the add step parameter params
func (o *AddStepParameterParams) WithStepID(stepID string) *AddStepParameterParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the add step parameter params
func (o *AddStepParameterParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *AddStepParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param parameterName
	if err := r.SetPathParam("parameterName", o.ParameterName); err != nil {
		return err
	}

	// path param stepId
	if err := r.SetPathParam("stepId", o.StepID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
