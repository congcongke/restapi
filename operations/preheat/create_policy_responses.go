// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// CreatePolicyCreatedCode is the HTTP code returned for type CreatePolicyCreated
const CreatePolicyCreatedCode int = 201

/*CreatePolicyCreated Created

swagger:response createPolicyCreated
*/
type CreatePolicyCreated struct {
	/*The location of the resource

	 */
	Location string `json:"Location"`
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewCreatePolicyCreated creates CreatePolicyCreated with default headers values
func NewCreatePolicyCreated() *CreatePolicyCreated {

	return &CreatePolicyCreated{}
}

// WithLocation adds the location to the create policy created response
func (o *CreatePolicyCreated) WithLocation(location string) *CreatePolicyCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the create policy created response
func (o *CreatePolicyCreated) SetLocation(location string) {
	o.Location = location
}

// WithXRequestID adds the xRequestId to the create policy created response
func (o *CreatePolicyCreated) WithXRequestID(xRequestID string) *CreatePolicyCreated {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create policy created response
func (o *CreatePolicyCreated) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *CreatePolicyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// CreatePolicyBadRequestCode is the HTTP code returned for type CreatePolicyBadRequest
const CreatePolicyBadRequestCode int = 400

/*CreatePolicyBadRequest Bad request

swagger:response createPolicyBadRequest
*/
type CreatePolicyBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreatePolicyBadRequest creates CreatePolicyBadRequest with default headers values
func NewCreatePolicyBadRequest() *CreatePolicyBadRequest {

	return &CreatePolicyBadRequest{}
}

// WithXRequestID adds the xRequestId to the create policy bad request response
func (o *CreatePolicyBadRequest) WithXRequestID(xRequestID string) *CreatePolicyBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create policy bad request response
func (o *CreatePolicyBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create policy bad request response
func (o *CreatePolicyBadRequest) WithPayload(payload *models.Errors) *CreatePolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create policy bad request response
func (o *CreatePolicyBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreatePolicyUnauthorizedCode is the HTTP code returned for type CreatePolicyUnauthorized
const CreatePolicyUnauthorizedCode int = 401

/*CreatePolicyUnauthorized Unauthorized

swagger:response createPolicyUnauthorized
*/
type CreatePolicyUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreatePolicyUnauthorized creates CreatePolicyUnauthorized with default headers values
func NewCreatePolicyUnauthorized() *CreatePolicyUnauthorized {

	return &CreatePolicyUnauthorized{}
}

// WithXRequestID adds the xRequestId to the create policy unauthorized response
func (o *CreatePolicyUnauthorized) WithXRequestID(xRequestID string) *CreatePolicyUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create policy unauthorized response
func (o *CreatePolicyUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create policy unauthorized response
func (o *CreatePolicyUnauthorized) WithPayload(payload *models.Errors) *CreatePolicyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create policy unauthorized response
func (o *CreatePolicyUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreatePolicyForbiddenCode is the HTTP code returned for type CreatePolicyForbidden
const CreatePolicyForbiddenCode int = 403

/*CreatePolicyForbidden Forbidden

swagger:response createPolicyForbidden
*/
type CreatePolicyForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreatePolicyForbidden creates CreatePolicyForbidden with default headers values
func NewCreatePolicyForbidden() *CreatePolicyForbidden {

	return &CreatePolicyForbidden{}
}

// WithXRequestID adds the xRequestId to the create policy forbidden response
func (o *CreatePolicyForbidden) WithXRequestID(xRequestID string) *CreatePolicyForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create policy forbidden response
func (o *CreatePolicyForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create policy forbidden response
func (o *CreatePolicyForbidden) WithPayload(payload *models.Errors) *CreatePolicyForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create policy forbidden response
func (o *CreatePolicyForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePolicyForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreatePolicyConflictCode is the HTTP code returned for type CreatePolicyConflict
const CreatePolicyConflictCode int = 409

/*CreatePolicyConflict Conflict

swagger:response createPolicyConflict
*/
type CreatePolicyConflict struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreatePolicyConflict creates CreatePolicyConflict with default headers values
func NewCreatePolicyConflict() *CreatePolicyConflict {

	return &CreatePolicyConflict{}
}

// WithXRequestID adds the xRequestId to the create policy conflict response
func (o *CreatePolicyConflict) WithXRequestID(xRequestID string) *CreatePolicyConflict {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create policy conflict response
func (o *CreatePolicyConflict) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create policy conflict response
func (o *CreatePolicyConflict) WithPayload(payload *models.Errors) *CreatePolicyConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create policy conflict response
func (o *CreatePolicyConflict) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePolicyConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreatePolicyInternalServerErrorCode is the HTTP code returned for type CreatePolicyInternalServerError
const CreatePolicyInternalServerErrorCode int = 500

/*CreatePolicyInternalServerError Internal server error

swagger:response createPolicyInternalServerError
*/
type CreatePolicyInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreatePolicyInternalServerError creates CreatePolicyInternalServerError with default headers values
func NewCreatePolicyInternalServerError() *CreatePolicyInternalServerError {

	return &CreatePolicyInternalServerError{}
}

// WithXRequestID adds the xRequestId to the create policy internal server error response
func (o *CreatePolicyInternalServerError) WithXRequestID(xRequestID string) *CreatePolicyInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create policy internal server error response
func (o *CreatePolicyInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create policy internal server error response
func (o *CreatePolicyInternalServerError) WithPayload(payload *models.Errors) *CreatePolicyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create policy internal server error response
func (o *CreatePolicyInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreatePolicyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
