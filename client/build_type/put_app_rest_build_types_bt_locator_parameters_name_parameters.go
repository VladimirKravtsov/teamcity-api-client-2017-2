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

// NewPutAppRestBuildTypesBtLocatorParametersNameParams creates a new PutAppRestBuildTypesBtLocatorParametersNameParams object
// with the default values initialized.
func NewPutAppRestBuildTypesBtLocatorParametersNameParams() *PutAppRestBuildTypesBtLocatorParametersNameParams {
	var ()
	return &PutAppRestBuildTypesBtLocatorParametersNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAppRestBuildTypesBtLocatorParametersNameParamsWithTimeout creates a new PutAppRestBuildTypesBtLocatorParametersNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAppRestBuildTypesBtLocatorParametersNameParamsWithTimeout(timeout time.Duration) *PutAppRestBuildTypesBtLocatorParametersNameParams {
	var ()
	return &PutAppRestBuildTypesBtLocatorParametersNameParams{

		timeout: timeout,
	}
}

// NewPutAppRestBuildTypesBtLocatorParametersNameParamsWithContext creates a new PutAppRestBuildTypesBtLocatorParametersNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAppRestBuildTypesBtLocatorParametersNameParamsWithContext(ctx context.Context) *PutAppRestBuildTypesBtLocatorParametersNameParams {
	var ()
	return &PutAppRestBuildTypesBtLocatorParametersNameParams{

		Context: ctx,
	}
}

// NewPutAppRestBuildTypesBtLocatorParametersNameParamsWithHTTPClient creates a new PutAppRestBuildTypesBtLocatorParametersNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAppRestBuildTypesBtLocatorParametersNameParamsWithHTTPClient(client *http.Client) *PutAppRestBuildTypesBtLocatorParametersNameParams {
	var ()
	return &PutAppRestBuildTypesBtLocatorParametersNameParams{
		HTTPClient: client,
	}
}

/*PutAppRestBuildTypesBtLocatorParametersNameParams contains all the parameters to send to the API endpoint
for the put app rest build types bt locator parameters name operation typically these are written to a http.Request
*/
type PutAppRestBuildTypesBtLocatorParametersNameParams struct {

	/*Body*/
	Body *models.Property
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) WithTimeout(timeout time.Duration) *PutAppRestBuildTypesBtLocatorParametersNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) WithContext(ctx context.Context) *PutAppRestBuildTypesBtLocatorParametersNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) WithHTTPClient(client *http.Client) *PutAppRestBuildTypesBtLocatorParametersNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) WithBody(body *models.Property) *PutAppRestBuildTypesBtLocatorParametersNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) SetBody(body *models.Property) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) WithBtLocator(btLocator string) *PutAppRestBuildTypesBtLocatorParametersNameParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) WithFields(fields *string) *PutAppRestBuildTypesBtLocatorParametersNameParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithName adds the name to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) WithName(name string) *PutAppRestBuildTypesBtLocatorParametersNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put app rest build types bt locator parameters name params
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PutAppRestBuildTypesBtLocatorParametersNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
