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
)

// NewRemoveUserPropertyParams creates a new RemoveUserPropertyParams object
// with the default values initialized.
func NewRemoveUserPropertyParams() *RemoveUserPropertyParams {
	var ()
	return &RemoveUserPropertyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveUserPropertyParamsWithTimeout creates a new RemoveUserPropertyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveUserPropertyParamsWithTimeout(timeout time.Duration) *RemoveUserPropertyParams {
	var ()
	return &RemoveUserPropertyParams{

		timeout: timeout,
	}
}

// NewRemoveUserPropertyParamsWithContext creates a new RemoveUserPropertyParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveUserPropertyParamsWithContext(ctx context.Context) *RemoveUserPropertyParams {
	var ()
	return &RemoveUserPropertyParams{

		Context: ctx,
	}
}

// NewRemoveUserPropertyParamsWithHTTPClient creates a new RemoveUserPropertyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveUserPropertyParamsWithHTTPClient(client *http.Client) *RemoveUserPropertyParams {
	var ()
	return &RemoveUserPropertyParams{
		HTTPClient: client,
	}
}

/*RemoveUserPropertyParams contains all the parameters to send to the API endpoint
for the remove user property operation typically these are written to a http.Request
*/
type RemoveUserPropertyParams struct {

	/*Name*/
	Name string
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove user property params
func (o *RemoveUserPropertyParams) WithTimeout(timeout time.Duration) *RemoveUserPropertyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove user property params
func (o *RemoveUserPropertyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove user property params
func (o *RemoveUserPropertyParams) WithContext(ctx context.Context) *RemoveUserPropertyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove user property params
func (o *RemoveUserPropertyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove user property params
func (o *RemoveUserPropertyParams) WithHTTPClient(client *http.Client) *RemoveUserPropertyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove user property params
func (o *RemoveUserPropertyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the remove user property params
func (o *RemoveUserPropertyParams) WithName(name string) *RemoveUserPropertyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the remove user property params
func (o *RemoveUserPropertyParams) SetName(name string) {
	o.Name = name
}

// WithUserLocator adds the userLocator to the remove user property params
func (o *RemoveUserPropertyParams) WithUserLocator(userLocator string) *RemoveUserPropertyParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the remove user property params
func (o *RemoveUserPropertyParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveUserPropertyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
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
