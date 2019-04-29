// Code generated by go-swagger; DO NOT EDIT.

package debug

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

// NewPostRequestDetailsParams creates a new PostRequestDetailsParams object
// with the default values initialized.
func NewPostRequestDetailsParams() *PostRequestDetailsParams {
	var ()
	return &PostRequestDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRequestDetailsParamsWithTimeout creates a new PostRequestDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRequestDetailsParamsWithTimeout(timeout time.Duration) *PostRequestDetailsParams {
	var ()
	return &PostRequestDetailsParams{

		timeout: timeout,
	}
}

// NewPostRequestDetailsParamsWithContext creates a new PostRequestDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRequestDetailsParamsWithContext(ctx context.Context) *PostRequestDetailsParams {
	var ()
	return &PostRequestDetailsParams{

		Context: ctx,
	}
}

// NewPostRequestDetailsParamsWithHTTPClient creates a new PostRequestDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRequestDetailsParamsWithHTTPClient(client *http.Client) *PostRequestDetailsParams {
	var ()
	return &PostRequestDetailsParams{
		HTTPClient: client,
	}
}

/*PostRequestDetailsParams contains all the parameters to send to the API endpoint
for the post request details operation typically these are written to a http.Request
*/
type PostRequestDetailsParams struct {

	/*Extra*/
	Extra string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post request details params
func (o *PostRequestDetailsParams) WithTimeout(timeout time.Duration) *PostRequestDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post request details params
func (o *PostRequestDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post request details params
func (o *PostRequestDetailsParams) WithContext(ctx context.Context) *PostRequestDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post request details params
func (o *PostRequestDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post request details params
func (o *PostRequestDetailsParams) WithHTTPClient(client *http.Client) *PostRequestDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post request details params
func (o *PostRequestDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExtra adds the extra to the post request details params
func (o *PostRequestDetailsParams) WithExtra(extra string) *PostRequestDetailsParams {
	o.SetExtra(extra)
	return o
}

// SetExtra adds the extra to the post request details params
func (o *PostRequestDetailsParams) SetExtra(extra string) {
	o.Extra = extra
}

// WriteToRequest writes these params to a swagger request
func (o *PostRequestDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param extra
	if err := r.SetPathParam("extra", o.Extra); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
