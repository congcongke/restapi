// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// CreateTagCreatedCode is the HTTP code returned for type CreateTagCreated
const CreateTagCreatedCode int = 201

/*CreateTagCreated Created

swagger:response createTagCreated
*/
type CreateTagCreated struct {
	/*The location of the resource

	 */
	Location string `json:"Location"`
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewCreateTagCreated creates CreateTagCreated with default headers values
func NewCreateTagCreated() *CreateTagCreated {

	return &CreateTagCreated{}
}

// WithLocation adds the location to the create tag created response
func (o *CreateTagCreated) WithLocation(location string) *CreateTagCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the create tag created response
func (o *CreateTagCreated) SetLocation(location string) {
	o.Location = location
}

// WithXRequestID adds the xRequestId to the create tag created response
func (o *CreateTagCreated) WithXRequestID(xRequestID string) *CreateTagCreated {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create tag created response
func (o *CreateTagCreated) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *CreateTagCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTagBadRequestCode is the HTTP code returned for type CreateTagBadRequest
const CreateTagBadRequestCode int = 400

/*CreateTagBadRequest Bad request

swagger:response createTagBadRequest
*/
type CreateTagBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateTagBadRequest creates CreateTagBadRequest with default headers values
func NewCreateTagBadRequest() *CreateTagBadRequest {

	return &CreateTagBadRequest{}
}

// WithXRequestID adds the xRequestId to the create tag bad request response
func (o *CreateTagBadRequest) WithXRequestID(xRequestID string) *CreateTagBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create tag bad request response
func (o *CreateTagBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create tag bad request response
func (o *CreateTagBadRequest) WithPayload(payload *models.Errors) *CreateTagBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create tag bad request response
func (o *CreateTagBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTagBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTagUnauthorizedCode is the HTTP code returned for type CreateTagUnauthorized
const CreateTagUnauthorizedCode int = 401

/*CreateTagUnauthorized Unauthorized

swagger:response createTagUnauthorized
*/
type CreateTagUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateTagUnauthorized creates CreateTagUnauthorized with default headers values
func NewCreateTagUnauthorized() *CreateTagUnauthorized {

	return &CreateTagUnauthorized{}
}

// WithXRequestID adds the xRequestId to the create tag unauthorized response
func (o *CreateTagUnauthorized) WithXRequestID(xRequestID string) *CreateTagUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create tag unauthorized response
func (o *CreateTagUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create tag unauthorized response
func (o *CreateTagUnauthorized) WithPayload(payload *models.Errors) *CreateTagUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create tag unauthorized response
func (o *CreateTagUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTagUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTagForbiddenCode is the HTTP code returned for type CreateTagForbidden
const CreateTagForbiddenCode int = 403

/*CreateTagForbidden Forbidden

swagger:response createTagForbidden
*/
type CreateTagForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateTagForbidden creates CreateTagForbidden with default headers values
func NewCreateTagForbidden() *CreateTagForbidden {

	return &CreateTagForbidden{}
}

// WithXRequestID adds the xRequestId to the create tag forbidden response
func (o *CreateTagForbidden) WithXRequestID(xRequestID string) *CreateTagForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create tag forbidden response
func (o *CreateTagForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create tag forbidden response
func (o *CreateTagForbidden) WithPayload(payload *models.Errors) *CreateTagForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create tag forbidden response
func (o *CreateTagForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTagForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTagNotFoundCode is the HTTP code returned for type CreateTagNotFound
const CreateTagNotFoundCode int = 404

/*CreateTagNotFound Not found

swagger:response createTagNotFound
*/
type CreateTagNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateTagNotFound creates CreateTagNotFound with default headers values
func NewCreateTagNotFound() *CreateTagNotFound {

	return &CreateTagNotFound{}
}

// WithXRequestID adds the xRequestId to the create tag not found response
func (o *CreateTagNotFound) WithXRequestID(xRequestID string) *CreateTagNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create tag not found response
func (o *CreateTagNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create tag not found response
func (o *CreateTagNotFound) WithPayload(payload *models.Errors) *CreateTagNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create tag not found response
func (o *CreateTagNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTagNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTagMethodNotAllowedCode is the HTTP code returned for type CreateTagMethodNotAllowed
const CreateTagMethodNotAllowedCode int = 405

/*CreateTagMethodNotAllowed Method not allowed

swagger:response createTagMethodNotAllowed
*/
type CreateTagMethodNotAllowed struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateTagMethodNotAllowed creates CreateTagMethodNotAllowed with default headers values
func NewCreateTagMethodNotAllowed() *CreateTagMethodNotAllowed {

	return &CreateTagMethodNotAllowed{}
}

// WithXRequestID adds the xRequestId to the create tag method not allowed response
func (o *CreateTagMethodNotAllowed) WithXRequestID(xRequestID string) *CreateTagMethodNotAllowed {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create tag method not allowed response
func (o *CreateTagMethodNotAllowed) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create tag method not allowed response
func (o *CreateTagMethodNotAllowed) WithPayload(payload *models.Errors) *CreateTagMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create tag method not allowed response
func (o *CreateTagMethodNotAllowed) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTagMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateTagConflictCode is the HTTP code returned for type CreateTagConflict
const CreateTagConflictCode int = 409

/*CreateTagConflict Conflict

swagger:response createTagConflict
*/
type CreateTagConflict struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateTagConflict creates CreateTagConflict with default headers values
func NewCreateTagConflict() *CreateTagConflict {

	return &CreateTagConflict{}
}

// WithXRequestID adds the xRequestId to the create tag conflict response
func (o *CreateTagConflict) WithXRequestID(xRequestID string) *CreateTagConflict {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create tag conflict response
func (o *CreateTagConflict) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create tag conflict response
func (o *CreateTagConflict) WithPayload(payload *models.Errors) *CreateTagConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create tag conflict response
func (o *CreateTagConflict) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTagConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateTagInternalServerErrorCode is the HTTP code returned for type CreateTagInternalServerError
const CreateTagInternalServerErrorCode int = 500

/*CreateTagInternalServerError Internal server error

swagger:response createTagInternalServerError
*/
type CreateTagInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateTagInternalServerError creates CreateTagInternalServerError with default headers values
func NewCreateTagInternalServerError() *CreateTagInternalServerError {

	return &CreateTagInternalServerError{}
}

// WithXRequestID adds the xRequestId to the create tag internal server error response
func (o *CreateTagInternalServerError) WithXRequestID(xRequestID string) *CreateTagInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create tag internal server error response
func (o *CreateTagInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create tag internal server error response
func (o *CreateTagInternalServerError) WithPayload(payload *models.Errors) *CreateTagInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create tag internal server error response
func (o *CreateTagInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTagInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
