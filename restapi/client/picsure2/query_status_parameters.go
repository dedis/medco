// Code generated by go-swagger; DO NOT EDIT.

package picsure2

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

// NewQueryStatusParams creates a new QueryStatusParams object
// with the default values initialized.
func NewQueryStatusParams() *QueryStatusParams {
	var ()
	return &QueryStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewQueryStatusParamsWithTimeout creates a new QueryStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewQueryStatusParamsWithTimeout(timeout time.Duration) *QueryStatusParams {
	var ()
	return &QueryStatusParams{

		timeout: timeout,
	}
}

// NewQueryStatusParamsWithContext creates a new QueryStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewQueryStatusParamsWithContext(ctx context.Context) *QueryStatusParams {
	var ()
	return &QueryStatusParams{

		Context: ctx,
	}
}

// NewQueryStatusParamsWithHTTPClient creates a new QueryStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewQueryStatusParamsWithHTTPClient(client *http.Client) *QueryStatusParams {
	var ()
	return &QueryStatusParams{
		HTTPClient: client,
	}
}

/*QueryStatusParams contains all the parameters to send to the API endpoint
for the query status operation typically these are written to a http.Request
*/
type QueryStatusParams struct {

	/*Body
	  Credentials to be used.

	*/
	Body QueryStatusBody
	/*QueryID
	  Query ID

	*/
	QueryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the query status params
func (o *QueryStatusParams) WithTimeout(timeout time.Duration) *QueryStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the query status params
func (o *QueryStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the query status params
func (o *QueryStatusParams) WithContext(ctx context.Context) *QueryStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the query status params
func (o *QueryStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the query status params
func (o *QueryStatusParams) WithHTTPClient(client *http.Client) *QueryStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the query status params
func (o *QueryStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the query status params
func (o *QueryStatusParams) WithBody(body QueryStatusBody) *QueryStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the query status params
func (o *QueryStatusParams) SetBody(body QueryStatusBody) {
	o.Body = body
}

// WithQueryID adds the queryID to the query status params
func (o *QueryStatusParams) WithQueryID(queryID string) *QueryStatusParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the query status params
func (o *QueryStatusParams) SetQueryID(queryID string) {
	o.QueryID = queryID
}

// WriteToRequest writes these params to a swagger request
func (o *QueryStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param queryId
	if err := r.SetPathParam("queryId", o.QueryID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
