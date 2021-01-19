// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewExploreQueryParams creates a new ExploreQueryParams object
// with the default values initialized.
func NewExploreQueryParams() *ExploreQueryParams {
	var (
		syncDefault = bool(true)
	)
	return &ExploreQueryParams{
		Sync: &syncDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewExploreQueryParamsWithTimeout creates a new ExploreQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExploreQueryParamsWithTimeout(timeout time.Duration) *ExploreQueryParams {
	var (
		syncDefault = bool(true)
	)
	return &ExploreQueryParams{
		Sync: &syncDefault,

		timeout: timeout,
	}
}

// NewExploreQueryParamsWithContext creates a new ExploreQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewExploreQueryParamsWithContext(ctx context.Context) *ExploreQueryParams {
	var (
		syncDefault = bool(true)
	)
	return &ExploreQueryParams{
		Sync: &syncDefault,

		Context: ctx,
	}
}

// NewExploreQueryParamsWithHTTPClient creates a new ExploreQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExploreQueryParamsWithHTTPClient(client *http.Client) *ExploreQueryParams {
	var (
		syncDefault = bool(true)
	)
	return &ExploreQueryParams{
		Sync:       &syncDefault,
		HTTPClient: client,
	}
}

/*ExploreQueryParams contains all the parameters to send to the API endpoint
for the explore query operation typically these are written to a http.Request
*/
type ExploreQueryParams struct {

	/*QueryRequest
	  MedCo-Explore query request.

	*/
	QueryRequest ExploreQueryBody
	/*Sync
	  Request synchronous query (defaults to true).

	*/
	Sync *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the explore query params
func (o *ExploreQueryParams) WithTimeout(timeout time.Duration) *ExploreQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the explore query params
func (o *ExploreQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the explore query params
func (o *ExploreQueryParams) WithContext(ctx context.Context) *ExploreQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the explore query params
func (o *ExploreQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the explore query params
func (o *ExploreQueryParams) WithHTTPClient(client *http.Client) *ExploreQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the explore query params
func (o *ExploreQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQueryRequest adds the queryRequest to the explore query params
func (o *ExploreQueryParams) WithQueryRequest(queryRequest ExploreQueryBody) *ExploreQueryParams {
	o.SetQueryRequest(queryRequest)
	return o
}

// SetQueryRequest adds the queryRequest to the explore query params
func (o *ExploreQueryParams) SetQueryRequest(queryRequest ExploreQueryBody) {
	o.QueryRequest = queryRequest
}

// WithSync adds the sync to the explore query params
func (o *ExploreQueryParams) WithSync(sync *bool) *ExploreQueryParams {
	o.SetSync(sync)
	return o
}

// SetSync adds the sync to the explore query params
func (o *ExploreQueryParams) SetSync(sync *bool) {
	o.Sync = sync
}

// WriteToRequest writes these params to a swagger request
func (o *ExploreQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.QueryRequest); err != nil {
		return err
	}

	if o.Sync != nil {

		// query param sync
		var qrSync bool
		if o.Sync != nil {
			qrSync = *o.Sync
		}
		qSync := swag.FormatBool(qrSync)
		if qSync != "" {
			if err := r.SetQueryParam("sync", qSync); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
