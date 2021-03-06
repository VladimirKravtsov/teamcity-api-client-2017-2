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

// NewGetParameterValueLongParams creates a new GetParameterValueLongParams object
// with the default values initialized.
func NewGetParameterValueLongParams() *GetParameterValueLongParams {
	var ()
	return &GetParameterValueLongParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetParameterValueLongParamsWithTimeout creates a new GetParameterValueLongParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetParameterValueLongParamsWithTimeout(timeout time.Duration) *GetParameterValueLongParams {
	var ()
	return &GetParameterValueLongParams{

		timeout: timeout,
	}
}

// NewGetParameterValueLongParamsWithContext creates a new GetParameterValueLongParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetParameterValueLongParamsWithContext(ctx context.Context) *GetParameterValueLongParams {
	var ()
	return &GetParameterValueLongParams{

		Context: ctx,
	}
}

// NewGetParameterValueLongParamsWithHTTPClient creates a new GetParameterValueLongParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetParameterValueLongParamsWithHTTPClient(client *http.Client) *GetParameterValueLongParams {
	var ()
	return &GetParameterValueLongParams{
		HTTPClient: client,
	}
}

/*GetParameterValueLongParams contains all the parameters to send to the API endpoint
for the get parameter value long operation typically these are written to a http.Request
*/
type GetParameterValueLongParams struct {

	/*FeatureLocator*/
	FeatureLocator string
	/*Fields*/
	Fields *string
	/*Name*/
	Name string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get parameter value long params
func (o *GetParameterValueLongParams) WithTimeout(timeout time.Duration) *GetParameterValueLongParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get parameter value long params
func (o *GetParameterValueLongParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get parameter value long params
func (o *GetParameterValueLongParams) WithContext(ctx context.Context) *GetParameterValueLongParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get parameter value long params
func (o *GetParameterValueLongParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get parameter value long params
func (o *GetParameterValueLongParams) WithHTTPClient(client *http.Client) *GetParameterValueLongParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get parameter value long params
func (o *GetParameterValueLongParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeatureLocator adds the featureLocator to the get parameter value long params
func (o *GetParameterValueLongParams) WithFeatureLocator(featureLocator string) *GetParameterValueLongParams {
	o.SetFeatureLocator(featureLocator)
	return o
}

// SetFeatureLocator adds the featureLocator to the get parameter value long params
func (o *GetParameterValueLongParams) SetFeatureLocator(featureLocator string) {
	o.FeatureLocator = featureLocator
}

// WithFields adds the fields to the get parameter value long params
func (o *GetParameterValueLongParams) WithFields(fields *string) *GetParameterValueLongParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get parameter value long params
func (o *GetParameterValueLongParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithName adds the name to the get parameter value long params
func (o *GetParameterValueLongParams) WithName(name string) *GetParameterValueLongParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get parameter value long params
func (o *GetParameterValueLongParams) SetName(name string) {
	o.Name = name
}

// WithProjectLocator adds the projectLocator to the get parameter value long params
func (o *GetParameterValueLongParams) WithProjectLocator(projectLocator string) *GetParameterValueLongParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the get parameter value long params
func (o *GetParameterValueLongParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetParameterValueLongParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param featureLocator
	if err := r.SetPathParam("featureLocator", o.FeatureLocator); err != nil {
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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
