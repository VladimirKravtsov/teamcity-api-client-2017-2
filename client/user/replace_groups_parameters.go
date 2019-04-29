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

	models "github.com/VladimirKravtsov/teamcity-api-client-2017-2/models"
)

// NewReplaceGroupsParams creates a new ReplaceGroupsParams object
// with the default values initialized.
func NewReplaceGroupsParams() *ReplaceGroupsParams {
	var ()
	return &ReplaceGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceGroupsParamsWithTimeout creates a new ReplaceGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceGroupsParamsWithTimeout(timeout time.Duration) *ReplaceGroupsParams {
	var ()
	return &ReplaceGroupsParams{

		timeout: timeout,
	}
}

// NewReplaceGroupsParamsWithContext creates a new ReplaceGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceGroupsParamsWithContext(ctx context.Context) *ReplaceGroupsParams {
	var ()
	return &ReplaceGroupsParams{

		Context: ctx,
	}
}

// NewReplaceGroupsParamsWithHTTPClient creates a new ReplaceGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceGroupsParamsWithHTTPClient(client *http.Client) *ReplaceGroupsParams {
	var ()
	return &ReplaceGroupsParams{
		HTTPClient: client,
	}
}

/*ReplaceGroupsParams contains all the parameters to send to the API endpoint
for the replace groups operation typically these are written to a http.Request
*/
type ReplaceGroupsParams struct {

	/*Body*/
	Body *models.Groups
	/*Fields*/
	Fields *string
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace groups params
func (o *ReplaceGroupsParams) WithTimeout(timeout time.Duration) *ReplaceGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace groups params
func (o *ReplaceGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace groups params
func (o *ReplaceGroupsParams) WithContext(ctx context.Context) *ReplaceGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace groups params
func (o *ReplaceGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace groups params
func (o *ReplaceGroupsParams) WithHTTPClient(client *http.Client) *ReplaceGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace groups params
func (o *ReplaceGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace groups params
func (o *ReplaceGroupsParams) WithBody(body *models.Groups) *ReplaceGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace groups params
func (o *ReplaceGroupsParams) SetBody(body *models.Groups) {
	o.Body = body
}

// WithFields adds the fields to the replace groups params
func (o *ReplaceGroupsParams) WithFields(fields *string) *ReplaceGroupsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace groups params
func (o *ReplaceGroupsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserLocator adds the userLocator to the replace groups params
func (o *ReplaceGroupsParams) WithUserLocator(userLocator string) *ReplaceGroupsParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the replace groups params
func (o *ReplaceGroupsParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userLocator
	if err := r.SetPathParam("userLocator", o.UserLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
