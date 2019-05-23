// Code generated by go-swagger; DO NOT EDIT.

package picsure2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new picsure2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for picsure2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetInfo returns information on how to interact with this p i c s u r e endpoint
*/
func (a *Client) GetInfo(params *GetInfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getInfo",
		Method:             "POST",
		PathPattern:        "/picsure2/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInfoOK), nil

}

/*
Query queries med co node
*/
func (a *Client) Query(params *QueryParams, authInfo runtime.ClientAuthInfoWriter) (*QueryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "query",
		Method:             "POST",
		PathPattern:        "/picsure2/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryOK), nil

}

/*
QueryResult gets result of query
*/
func (a *Client) QueryResult(params *QueryResultParams, authInfo runtime.ClientAuthInfoWriter) (*QueryResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryResultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryResult",
		Method:             "POST",
		PathPattern:        "/picsure2/{queryId}/result",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryResultReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryResultOK), nil

}

/*
QueryStatus gets status of query
*/
func (a *Client) QueryStatus(params *QueryStatusParams, authInfo runtime.ClientAuthInfoWriter) (*QueryStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryStatus",
		Method:             "POST",
		PathPattern:        "/picsure2/{queryId}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QueryStatusOK), nil

}

/*
QuerySync queries med co node synchronously
*/
func (a *Client) QuerySync(params *QuerySyncParams, authInfo runtime.ClientAuthInfoWriter) (*QuerySyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "querySync",
		Method:             "POST",
		PathPattern:        "/picsure2/query/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QuerySyncReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*QuerySyncOK), nil

}

/*
Search searches through the ontology
*/
func (a *Client) Search(params *SearchParams, authInfo runtime.ClientAuthInfoWriter) (*SearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "search",
		Method:             "POST",
		PathPattern:        "/picsure2/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
