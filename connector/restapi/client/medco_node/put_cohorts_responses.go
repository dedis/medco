// Code generated by go-swagger; DO NOT EDIT.

package medco_node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PutCohortsReader is a Reader for the PutCohorts structure.
type PutCohortsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutCohortsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutCohortsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutCohortsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutCohortsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutCohortsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutCohortsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutCohortsOK creates a PutCohortsOK with default headers values
func NewPutCohortsOK() *PutCohortsOK {
	return &PutCohortsOK{}
}

/*PutCohortsOK handles this case with default header values.

Updated cohort
*/
type PutCohortsOK struct {
}

func (o *PutCohortsOK) Error() string {
	return fmt.Sprintf("[PUT /node/explore/cohorts/{name}][%d] putCohortsOK ", 200)
}

func (o *PutCohortsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutCohortsBadRequest creates a PutCohortsBadRequest with default headers values
func NewPutCohortsBadRequest() *PutCohortsBadRequest {
	return &PutCohortsBadRequest{}
}

/*PutCohortsBadRequest handles this case with default header values.

Bad user input in request.
*/
type PutCohortsBadRequest struct {
	Payload *PutCohortsBadRequestBody
}

func (o *PutCohortsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /node/explore/cohorts/{name}][%d] putCohortsBadRequest  %+v", 400, o.Payload)
}

func (o *PutCohortsBadRequest) GetPayload() *PutCohortsBadRequestBody {
	return o.Payload
}

func (o *PutCohortsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCohortsBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCohortsNotFound creates a PutCohortsNotFound with default headers values
func NewPutCohortsNotFound() *PutCohortsNotFound {
	return &PutCohortsNotFound{}
}

/*PutCohortsNotFound handles this case with default header values.

Not found.
*/
type PutCohortsNotFound struct {
	Payload *PutCohortsNotFoundBody
}

func (o *PutCohortsNotFound) Error() string {
	return fmt.Sprintf("[PUT /node/explore/cohorts/{name}][%d] putCohortsNotFound  %+v", 404, o.Payload)
}

func (o *PutCohortsNotFound) GetPayload() *PutCohortsNotFoundBody {
	return o.Payload
}

func (o *PutCohortsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCohortsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCohortsConflict creates a PutCohortsConflict with default headers values
func NewPutCohortsConflict() *PutCohortsConflict {
	return &PutCohortsConflict{}
}

/*PutCohortsConflict handles this case with default header values.

Conflict with resource's state.
*/
type PutCohortsConflict struct {
	Payload *PutCohortsConflictBody
}

func (o *PutCohortsConflict) Error() string {
	return fmt.Sprintf("[PUT /node/explore/cohorts/{name}][%d] putCohortsConflict  %+v", 409, o.Payload)
}

func (o *PutCohortsConflict) GetPayload() *PutCohortsConflictBody {
	return o.Payload
}

func (o *PutCohortsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCohortsConflictBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutCohortsDefault creates a PutCohortsDefault with default headers values
func NewPutCohortsDefault(code int) *PutCohortsDefault {
	return &PutCohortsDefault{
		_statusCode: code,
	}
}

/*PutCohortsDefault handles this case with default header values.

Error response.
*/
type PutCohortsDefault struct {
	_statusCode int

	Payload *PutCohortsDefaultBody
}

// Code gets the status code for the put cohorts default response
func (o *PutCohortsDefault) Code() int {
	return o._statusCode
}

func (o *PutCohortsDefault) Error() string {
	return fmt.Sprintf("[PUT /node/explore/cohorts/{name}][%d] putCohorts default  %+v", o._statusCode, o.Payload)
}

func (o *PutCohortsDefault) GetPayload() *PutCohortsDefaultBody {
	return o.Payload
}

func (o *PutCohortsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PutCohortsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PutCohortsBadRequestBody put cohorts bad request body
swagger:model PutCohortsBadRequestBody
*/
type PutCohortsBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this put cohorts bad request body
func (o *PutCohortsBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCohortsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCohortsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutCohortsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutCohortsBody put cohorts body
swagger:model PutCohortsBody
*/
type PutCohortsBody struct {

	// creation date
	// Required: true
	CreationDate *string `json:"creationDate"`

	// patient set ID
	// Required: true
	PatientSetID *int64 `json:"patientSetID"`

	// update date
	// Required: true
	UpdateDate *string `json:"updateDate"`
}

// Validate validates this put cohorts body
func (o *PutCohortsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePatientSetID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutCohortsBody) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("cohortsRequest"+"."+"creationDate", "body", o.CreationDate); err != nil {
		return err
	}

	return nil
}

func (o *PutCohortsBody) validatePatientSetID(formats strfmt.Registry) error {

	if err := validate.Required("cohortsRequest"+"."+"patientSetID", "body", o.PatientSetID); err != nil {
		return err
	}

	return nil
}

func (o *PutCohortsBody) validateUpdateDate(formats strfmt.Registry) error {

	if err := validate.Required("cohortsRequest"+"."+"updateDate", "body", o.UpdateDate); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutCohortsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCohortsBody) UnmarshalBinary(b []byte) error {
	var res PutCohortsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutCohortsConflictBody put cohorts conflict body
swagger:model PutCohortsConflictBody
*/
type PutCohortsConflictBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this put cohorts conflict body
func (o *PutCohortsConflictBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCohortsConflictBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCohortsConflictBody) UnmarshalBinary(b []byte) error {
	var res PutCohortsConflictBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutCohortsDefaultBody put cohorts default body
swagger:model PutCohortsDefaultBody
*/
type PutCohortsDefaultBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this put cohorts default body
func (o *PutCohortsDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCohortsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCohortsDefaultBody) UnmarshalBinary(b []byte) error {
	var res PutCohortsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PutCohortsNotFoundBody put cohorts not found body
swagger:model PutCohortsNotFoundBody
*/
type PutCohortsNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this put cohorts not found body
func (o *PutCohortsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutCohortsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutCohortsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutCohortsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
