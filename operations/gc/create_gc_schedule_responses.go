// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// CreateGCScheduleCreatedCode is the HTTP code returned for type CreateGCScheduleCreated
const CreateGCScheduleCreatedCode int = 201

/*CreateGCScheduleCreated Created

swagger:response createGcScheduleCreated
*/
type CreateGCScheduleCreated struct {
	/*The location of the resource

	 */
	Location string `json:"Location"`
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewCreateGCScheduleCreated creates CreateGCScheduleCreated with default headers values
func NewCreateGCScheduleCreated() *CreateGCScheduleCreated {

	return &CreateGCScheduleCreated{}
}

// WithLocation adds the location to the create Gc schedule created response
func (o *CreateGCScheduleCreated) WithLocation(location string) *CreateGCScheduleCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the create Gc schedule created response
func (o *CreateGCScheduleCreated) SetLocation(location string) {
	o.Location = location
}

// WithXRequestID adds the xRequestId to the create Gc schedule created response
func (o *CreateGCScheduleCreated) WithXRequestID(xRequestID string) *CreateGCScheduleCreated {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create Gc schedule created response
func (o *CreateGCScheduleCreated) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *CreateGCScheduleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateGCScheduleBadRequestCode is the HTTP code returned for type CreateGCScheduleBadRequest
const CreateGCScheduleBadRequestCode int = 400

/*CreateGCScheduleBadRequest Bad request

swagger:response createGcScheduleBadRequest
*/
type CreateGCScheduleBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateGCScheduleBadRequest creates CreateGCScheduleBadRequest with default headers values
func NewCreateGCScheduleBadRequest() *CreateGCScheduleBadRequest {

	return &CreateGCScheduleBadRequest{}
}

// WithXRequestID adds the xRequestId to the create Gc schedule bad request response
func (o *CreateGCScheduleBadRequest) WithXRequestID(xRequestID string) *CreateGCScheduleBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create Gc schedule bad request response
func (o *CreateGCScheduleBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create Gc schedule bad request response
func (o *CreateGCScheduleBadRequest) WithPayload(payload *models.Errors) *CreateGCScheduleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Gc schedule bad request response
func (o *CreateGCScheduleBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGCScheduleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateGCScheduleUnauthorizedCode is the HTTP code returned for type CreateGCScheduleUnauthorized
const CreateGCScheduleUnauthorizedCode int = 401

/*CreateGCScheduleUnauthorized Unauthorized

swagger:response createGcScheduleUnauthorized
*/
type CreateGCScheduleUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateGCScheduleUnauthorized creates CreateGCScheduleUnauthorized with default headers values
func NewCreateGCScheduleUnauthorized() *CreateGCScheduleUnauthorized {

	return &CreateGCScheduleUnauthorized{}
}

// WithXRequestID adds the xRequestId to the create Gc schedule unauthorized response
func (o *CreateGCScheduleUnauthorized) WithXRequestID(xRequestID string) *CreateGCScheduleUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create Gc schedule unauthorized response
func (o *CreateGCScheduleUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create Gc schedule unauthorized response
func (o *CreateGCScheduleUnauthorized) WithPayload(payload *models.Errors) *CreateGCScheduleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Gc schedule unauthorized response
func (o *CreateGCScheduleUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGCScheduleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateGCScheduleForbiddenCode is the HTTP code returned for type CreateGCScheduleForbidden
const CreateGCScheduleForbiddenCode int = 403

/*CreateGCScheduleForbidden Forbidden

swagger:response createGcScheduleForbidden
*/
type CreateGCScheduleForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateGCScheduleForbidden creates CreateGCScheduleForbidden with default headers values
func NewCreateGCScheduleForbidden() *CreateGCScheduleForbidden {

	return &CreateGCScheduleForbidden{}
}

// WithXRequestID adds the xRequestId to the create Gc schedule forbidden response
func (o *CreateGCScheduleForbidden) WithXRequestID(xRequestID string) *CreateGCScheduleForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create Gc schedule forbidden response
func (o *CreateGCScheduleForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create Gc schedule forbidden response
func (o *CreateGCScheduleForbidden) WithPayload(payload *models.Errors) *CreateGCScheduleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Gc schedule forbidden response
func (o *CreateGCScheduleForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGCScheduleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateGCScheduleConflictCode is the HTTP code returned for type CreateGCScheduleConflict
const CreateGCScheduleConflictCode int = 409

/*CreateGCScheduleConflict Conflict

swagger:response createGcScheduleConflict
*/
type CreateGCScheduleConflict struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateGCScheduleConflict creates CreateGCScheduleConflict with default headers values
func NewCreateGCScheduleConflict() *CreateGCScheduleConflict {

	return &CreateGCScheduleConflict{}
}

// WithXRequestID adds the xRequestId to the create Gc schedule conflict response
func (o *CreateGCScheduleConflict) WithXRequestID(xRequestID string) *CreateGCScheduleConflict {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create Gc schedule conflict response
func (o *CreateGCScheduleConflict) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create Gc schedule conflict response
func (o *CreateGCScheduleConflict) WithPayload(payload *models.Errors) *CreateGCScheduleConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Gc schedule conflict response
func (o *CreateGCScheduleConflict) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGCScheduleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateGCScheduleInternalServerErrorCode is the HTTP code returned for type CreateGCScheduleInternalServerError
const CreateGCScheduleInternalServerErrorCode int = 500

/*CreateGCScheduleInternalServerError Internal server error

swagger:response createGcScheduleInternalServerError
*/
type CreateGCScheduleInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateGCScheduleInternalServerError creates CreateGCScheduleInternalServerError with default headers values
func NewCreateGCScheduleInternalServerError() *CreateGCScheduleInternalServerError {

	return &CreateGCScheduleInternalServerError{}
}

// WithXRequestID adds the xRequestId to the create Gc schedule internal server error response
func (o *CreateGCScheduleInternalServerError) WithXRequestID(xRequestID string) *CreateGCScheduleInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create Gc schedule internal server error response
func (o *CreateGCScheduleInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create Gc schedule internal server error response
func (o *CreateGCScheduleInternalServerError) WithPayload(payload *models.Errors) *CreateGCScheduleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Gc schedule internal server error response
func (o *CreateGCScheduleInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGCScheduleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
