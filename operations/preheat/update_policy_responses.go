// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// UpdatePolicyOKCode is the HTTP code returned for type UpdatePolicyOK
const UpdatePolicyOKCode int = 200

/*UpdatePolicyOK Success

swagger:response updatePolicyOK
*/
type UpdatePolicyOK struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewUpdatePolicyOK creates UpdatePolicyOK with default headers values
func NewUpdatePolicyOK() *UpdatePolicyOK {

	return &UpdatePolicyOK{}
}

// WithXRequestID adds the xRequestId to the update policy o k response
func (o *UpdatePolicyOK) WithXRequestID(xRequestID string) *UpdatePolicyOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update policy o k response
func (o *UpdatePolicyOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *UpdatePolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdatePolicyBadRequestCode is the HTTP code returned for type UpdatePolicyBadRequest
const UpdatePolicyBadRequestCode int = 400

/*UpdatePolicyBadRequest Bad request

swagger:response updatePolicyBadRequest
*/
type UpdatePolicyBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdatePolicyBadRequest creates UpdatePolicyBadRequest with default headers values
func NewUpdatePolicyBadRequest() *UpdatePolicyBadRequest {

	return &UpdatePolicyBadRequest{}
}

// WithXRequestID adds the xRequestId to the update policy bad request response
func (o *UpdatePolicyBadRequest) WithXRequestID(xRequestID string) *UpdatePolicyBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update policy bad request response
func (o *UpdatePolicyBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update policy bad request response
func (o *UpdatePolicyBadRequest) WithPayload(payload *models.Errors) *UpdatePolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy bad request response
func (o *UpdatePolicyBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyUnauthorizedCode is the HTTP code returned for type UpdatePolicyUnauthorized
const UpdatePolicyUnauthorizedCode int = 401

/*UpdatePolicyUnauthorized Unauthorized

swagger:response updatePolicyUnauthorized
*/
type UpdatePolicyUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdatePolicyUnauthorized creates UpdatePolicyUnauthorized with default headers values
func NewUpdatePolicyUnauthorized() *UpdatePolicyUnauthorized {

	return &UpdatePolicyUnauthorized{}
}

// WithXRequestID adds the xRequestId to the update policy unauthorized response
func (o *UpdatePolicyUnauthorized) WithXRequestID(xRequestID string) *UpdatePolicyUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update policy unauthorized response
func (o *UpdatePolicyUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update policy unauthorized response
func (o *UpdatePolicyUnauthorized) WithPayload(payload *models.Errors) *UpdatePolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy unauthorized response
func (o *UpdatePolicyUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyForbiddenCode is the HTTP code returned for type UpdatePolicyForbidden
const UpdatePolicyForbiddenCode int = 403

/*UpdatePolicyForbidden Forbidden

swagger:response updatePolicyForbidden
*/
type UpdatePolicyForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdatePolicyForbidden creates UpdatePolicyForbidden with default headers values
func NewUpdatePolicyForbidden() *UpdatePolicyForbidden {

	return &UpdatePolicyForbidden{}
}

// WithXRequestID adds the xRequestId to the update policy forbidden response
func (o *UpdatePolicyForbidden) WithXRequestID(xRequestID string) *UpdatePolicyForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update policy forbidden response
func (o *UpdatePolicyForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update policy forbidden response
func (o *UpdatePolicyForbidden) WithPayload(payload *models.Errors) *UpdatePolicyForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy forbidden response
func (o *UpdatePolicyForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyNotFoundCode is the HTTP code returned for type UpdatePolicyNotFound
const UpdatePolicyNotFoundCode int = 404

/*UpdatePolicyNotFound Not found

swagger:response updatePolicyNotFound
*/
type UpdatePolicyNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdatePolicyNotFound creates UpdatePolicyNotFound with default headers values
func NewUpdatePolicyNotFound() *UpdatePolicyNotFound {

	return &UpdatePolicyNotFound{}
}

// WithXRequestID adds the xRequestId to the update policy not found response
func (o *UpdatePolicyNotFound) WithXRequestID(xRequestID string) *UpdatePolicyNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update policy not found response
func (o *UpdatePolicyNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update policy not found response
func (o *UpdatePolicyNotFound) WithPayload(payload *models.Errors) *UpdatePolicyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy not found response
func (o *UpdatePolicyNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyConflictCode is the HTTP code returned for type UpdatePolicyConflict
const UpdatePolicyConflictCode int = 409

/*UpdatePolicyConflict Conflict

swagger:response updatePolicyConflict
*/
type UpdatePolicyConflict struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdatePolicyConflict creates UpdatePolicyConflict with default headers values
func NewUpdatePolicyConflict() *UpdatePolicyConflict {

	return &UpdatePolicyConflict{}
}

// WithXRequestID adds the xRequestId to the update policy conflict response
func (o *UpdatePolicyConflict) WithXRequestID(xRequestID string) *UpdatePolicyConflict {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update policy conflict response
func (o *UpdatePolicyConflict) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update policy conflict response
func (o *UpdatePolicyConflict) WithPayload(payload *models.Errors) *UpdatePolicyConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy conflict response
func (o *UpdatePolicyConflict) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePolicyInternalServerErrorCode is the HTTP code returned for type UpdatePolicyInternalServerError
const UpdatePolicyInternalServerErrorCode int = 500

/*UpdatePolicyInternalServerError Internal server error

swagger:response updatePolicyInternalServerError
*/
type UpdatePolicyInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdatePolicyInternalServerError creates UpdatePolicyInternalServerError with default headers values
func NewUpdatePolicyInternalServerError() *UpdatePolicyInternalServerError {

	return &UpdatePolicyInternalServerError{}
}

// WithXRequestID adds the xRequestId to the update policy internal server error response
func (o *UpdatePolicyInternalServerError) WithXRequestID(xRequestID string) *UpdatePolicyInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update policy internal server error response
func (o *UpdatePolicyInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update policy internal server error response
func (o *UpdatePolicyInternalServerError) WithPayload(payload *models.Errors) *UpdatePolicyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update policy internal server error response
func (o *UpdatePolicyInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePolicyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
