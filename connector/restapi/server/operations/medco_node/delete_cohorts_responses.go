// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteCohortsOKCode is the HTTP code returned for type DeleteCohortsOK
const DeleteCohortsOKCode int = 200

/*DeleteCohortsOK Deleted cohort

swagger:response deleteCohortsOK
*/
type DeleteCohortsOK struct {
}

// NewDeleteCohortsOK creates DeleteCohortsOK with default headers values
func NewDeleteCohortsOK() *DeleteCohortsOK {

	return &DeleteCohortsOK{}
}

// WriteResponse to the client
func (o *DeleteCohortsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*DeleteCohortsDefault Error response.

swagger:response deleteCohortsDefault
*/
type DeleteCohortsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *DeleteCohortsDefaultBody `json:"body,omitempty"`
}

// NewDeleteCohortsDefault creates DeleteCohortsDefault with default headers values
func NewDeleteCohortsDefault(code int) *DeleteCohortsDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteCohortsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete cohorts default response
func (o *DeleteCohortsDefault) WithStatusCode(code int) *DeleteCohortsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete cohorts default response
func (o *DeleteCohortsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete cohorts default response
func (o *DeleteCohortsDefault) WithPayload(payload *DeleteCohortsDefaultBody) *DeleteCohortsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete cohorts default response
func (o *DeleteCohortsDefault) SetPayload(payload *DeleteCohortsDefaultBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCohortsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
