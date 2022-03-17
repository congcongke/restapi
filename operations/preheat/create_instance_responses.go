// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// CreateInstanceCreatedCode is the HTTP code returned for type CreateInstanceCreated
const CreateInstanceCreatedCode int = 201

/*CreateInstanceCreated Created

swagger:response createInstanceCreated
*/
type CreateInstanceCreated struct {
	/*The location of the resource

	 */
	Location string `json:"Location"`
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewCreateInstanceCreated creates CreateInstanceCreated with default headers values
func NewCreateInstanceCreated() *CreateInstanceCreated {

	return &CreateInstanceCreated{}
}

// WithLocation adds the location to the create instance created response
func (o *CreateInstanceCreated) WithLocation(location string) *CreateInstanceCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the create instance created response
func (o *CreateInstanceCreated) SetLocation(location string) {
	o.Location = location
}

// WithXRequestID adds the xRequestId to the create instance created response
func (o *CreateInstanceCreated) WithXRequestID(xRequestID string) *CreateInstanceCreated {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create instance created response
func (o *CreateInstanceCreated) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *CreateInstanceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateInstanceBadRequestCode is the HTTP code returned for type CreateInstanceBadRequest
const CreateInstanceBadRequestCode int = 400

/*CreateInstanceBadRequest Bad request

swagger:response createInstanceBadRequest
*/
type CreateInstanceBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateInstanceBadRequest creates CreateInstanceBadRequest with default headers values
func NewCreateInstanceBadRequest() *CreateInstanceBadRequest {

	return &CreateInstanceBadRequest{}
}

// WithXRequestID adds the xRequestId to the create instance bad request response
func (o *CreateInstanceBadRequest) WithXRequestID(xRequestID string) *CreateInstanceBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create instance bad request response
func (o *CreateInstanceBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create instance bad request response
func (o *CreateInstanceBadRequest) WithPayload(payload *models.Errors) *CreateInstanceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create instance bad request response
func (o *CreateInstanceBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateInstanceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateInstanceUnauthorizedCode is the HTTP code returned for type CreateInstanceUnauthorized
const CreateInstanceUnauthorizedCode int = 401

/*CreateInstanceUnauthorized Unauthorized

swagger:response createInstanceUnauthorized
*/
type CreateInstanceUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateInstanceUnauthorized creates CreateInstanceUnauthorized with default headers values
func NewCreateInstanceUnauthorized() *CreateInstanceUnauthorized {

	return &CreateInstanceUnauthorized{}
}

// WithXRequestID adds the xRequestId to the create instance unauthorized response
func (o *CreateInstanceUnauthorized) WithXRequestID(xRequestID string) *CreateInstanceUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create instance unauthorized response
func (o *CreateInstanceUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create instance unauthorized response
func (o *CreateInstanceUnauthorized) WithPayload(payload *models.Errors) *CreateInstanceUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create instance unauthorized response
func (o *CreateInstanceUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateInstanceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateInstanceForbiddenCode is the HTTP code returned for type CreateInstanceForbidden
const CreateInstanceForbiddenCode int = 403

/*CreateInstanceForbidden Forbidden

swagger:response createInstanceForbidden
*/
type CreateInstanceForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateInstanceForbidden creates CreateInstanceForbidden with default headers values
func NewCreateInstanceForbidden() *CreateInstanceForbidden {

	return &CreateInstanceForbidden{}
}

// WithXRequestID adds the xRequestId to the create instance forbidden response
func (o *CreateInstanceForbidden) WithXRequestID(xRequestID string) *CreateInstanceForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create instance forbidden response
func (o *CreateInstanceForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create instance forbidden response
func (o *CreateInstanceForbidden) WithPayload(payload *models.Errors) *CreateInstanceForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create instance forbidden response
func (o *CreateInstanceForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateInstanceForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateInstanceNotFoundCode is the HTTP code returned for type CreateInstanceNotFound
const CreateInstanceNotFoundCode int = 404

/*CreateInstanceNotFound Not found

swagger:response createInstanceNotFound
*/
type CreateInstanceNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateInstanceNotFound creates CreateInstanceNotFound with default headers values
func NewCreateInstanceNotFound() *CreateInstanceNotFound {

	return &CreateInstanceNotFound{}
}

// WithXRequestID adds the xRequestId to the create instance not found response
func (o *CreateInstanceNotFound) WithXRequestID(xRequestID string) *CreateInstanceNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create instance not found response
func (o *CreateInstanceNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create instance not found response
func (o *CreateInstanceNotFound) WithPayload(payload *models.Errors) *CreateInstanceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create instance not found response
func (o *CreateInstanceNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateInstanceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateInstanceConflictCode is the HTTP code returned for type CreateInstanceConflict
const CreateInstanceConflictCode int = 409

/*CreateInstanceConflict Conflict

swagger:response createInstanceConflict
*/
type CreateInstanceConflict struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateInstanceConflict creates CreateInstanceConflict with default headers values
func NewCreateInstanceConflict() *CreateInstanceConflict {

	return &CreateInstanceConflict{}
}

// WithXRequestID adds the xRequestId to the create instance conflict response
func (o *CreateInstanceConflict) WithXRequestID(xRequestID string) *CreateInstanceConflict {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create instance conflict response
func (o *CreateInstanceConflict) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create instance conflict response
func (o *CreateInstanceConflict) WithPayload(payload *models.Errors) *CreateInstanceConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create instance conflict response
func (o *CreateInstanceConflict) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateInstanceConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateInstanceInternalServerErrorCode is the HTTP code returned for type CreateInstanceInternalServerError
const CreateInstanceInternalServerErrorCode int = 500

/*CreateInstanceInternalServerError Internal server error

swagger:response createInstanceInternalServerError
*/
type CreateInstanceInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateInstanceInternalServerError creates CreateInstanceInternalServerError with default headers values
func NewCreateInstanceInternalServerError() *CreateInstanceInternalServerError {

	return &CreateInstanceInternalServerError{}
}

// WithXRequestID adds the xRequestId to the create instance internal server error response
func (o *CreateInstanceInternalServerError) WithXRequestID(xRequestID string) *CreateInstanceInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create instance internal server error response
func (o *CreateInstanceInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create instance internal server error response
func (o *CreateInstanceInternalServerError) WithPayload(payload *models.Errors) *CreateInstanceInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create instance internal server error response
func (o *CreateInstanceInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateInstanceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
