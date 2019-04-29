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

// NewGetArtifactDepParams creates a new GetArtifactDepParams object
// with the default values initialized.
func NewGetArtifactDepParams() *GetArtifactDepParams {
	var ()
	return &GetArtifactDepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArtifactDepParamsWithTimeout creates a new GetArtifactDepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArtifactDepParamsWithTimeout(timeout time.Duration) *GetArtifactDepParams {
	var ()
	return &GetArtifactDepParams{

		timeout: timeout,
	}
}

// NewGetArtifactDepParamsWithContext creates a new GetArtifactDepParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArtifactDepParamsWithContext(ctx context.Context) *GetArtifactDepParams {
	var ()
	return &GetArtifactDepParams{

		Context: ctx,
	}
}

// NewGetArtifactDepParamsWithHTTPClient creates a new GetArtifactDepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArtifactDepParamsWithHTTPClient(client *http.Client) *GetArtifactDepParams {
	var ()
	return &GetArtifactDepParams{
		HTTPClient: client,
	}
}

/*GetArtifactDepParams contains all the parameters to send to the API endpoint
for the get artifact dep operation typically these are written to a http.Request
*/
type GetArtifactDepParams struct {

	/*ArtifactDepLocator*/
	ArtifactDepLocator string
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get artifact dep params
func (o *GetArtifactDepParams) WithTimeout(timeout time.Duration) *GetArtifactDepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get artifact dep params
func (o *GetArtifactDepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get artifact dep params
func (o *GetArtifactDepParams) WithContext(ctx context.Context) *GetArtifactDepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get artifact dep params
func (o *GetArtifactDepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get artifact dep params
func (o *GetArtifactDepParams) WithHTTPClient(client *http.Client) *GetArtifactDepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get artifact dep params
func (o *GetArtifactDepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactDepLocator adds the artifactDepLocator to the get artifact dep params
func (o *GetArtifactDepParams) WithArtifactDepLocator(artifactDepLocator string) *GetArtifactDepParams {
	o.SetArtifactDepLocator(artifactDepLocator)
	return o
}

// SetArtifactDepLocator adds the artifactDepLocator to the get artifact dep params
func (o *GetArtifactDepParams) SetArtifactDepLocator(artifactDepLocator string) {
	o.ArtifactDepLocator = artifactDepLocator
}

// WithBtLocator adds the btLocator to the get artifact dep params
func (o *GetArtifactDepParams) WithBtLocator(btLocator string) *GetArtifactDepParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get artifact dep params
func (o *GetArtifactDepParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get artifact dep params
func (o *GetArtifactDepParams) WithFields(fields *string) *GetArtifactDepParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get artifact dep params
func (o *GetArtifactDepParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetArtifactDepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param artifactDepLocator
	if err := r.SetPathParam("artifactDepLocator", o.ArtifactDepLocator); err != nil {
		return err
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
