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

// NewAddBuildTypeParams creates a new AddBuildTypeParams object
// with the default values initialized.
func NewAddBuildTypeParams() *AddBuildTypeParams {
	var ()
	return &AddBuildTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddBuildTypeParamsWithTimeout creates a new AddBuildTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddBuildTypeParamsWithTimeout(timeout time.Duration) *AddBuildTypeParams {
	var ()
	return &AddBuildTypeParams{

		timeout: timeout,
	}
}

// NewAddBuildTypeParamsWithContext creates a new AddBuildTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddBuildTypeParamsWithContext(ctx context.Context) *AddBuildTypeParams {
	var ()
	return &AddBuildTypeParams{

		Context: ctx,
	}
}

// NewAddBuildTypeParamsWithHTTPClient creates a new AddBuildTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddBuildTypeParamsWithHTTPClient(client *http.Client) *AddBuildTypeParams {
	var ()
	return &AddBuildTypeParams{
		HTTPClient: client,
	}
}

/*AddBuildTypeParams contains all the parameters to send to the API endpoint
for the add build type operation typically these are written to a http.Request
*/
type AddBuildTypeParams struct {

	/*Body*/
	Body *models.BuildType
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add build type params
func (o *AddBuildTypeParams) WithTimeout(timeout time.Duration) *AddBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add build type params
func (o *AddBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add build type params
func (o *AddBuildTypeParams) WithContext(ctx context.Context) *AddBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add build type params
func (o *AddBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add build type params
func (o *AddBuildTypeParams) WithHTTPClient(client *http.Client) *AddBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add build type params
func (o *AddBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add build type params
func (o *AddBuildTypeParams) WithBody(body *models.BuildType) *AddBuildTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add build type params
func (o *AddBuildTypeParams) SetBody(body *models.BuildType) {
	o.Body = body
}

// WithFields adds the fields to the add build type params
func (o *AddBuildTypeParams) WithFields(fields *string) *AddBuildTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add build type params
func (o *AddBuildTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
