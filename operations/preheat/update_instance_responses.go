// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// UpdateInstanceOKCode is the HTTP code returned for type UpdateInstanceOK
const UpdateInstanceOKCode int = 200

/*UpdateInstanceOK Success

swagger:response updateInstanceOK
*/
type UpdateInstanceOK struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewUpdateInstanceOK creates UpdateInstanceOK with default headers values
func NewUpdateInstanceOK() *UpdateInstanceOK {

	return &UpdateInstanceOK{}
}

// WithXRequestID adds the xRequestId to the update instance o k response
func (o *UpdateInstanceOK) WithXRequestID(xRequestID string) *UpdateInstanceOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update instance o k response
func (o *UpdateInstanceOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *UpdateInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateInstanceBadRequestCode is the HTTP code returned for type UpdateInstanceBadRequest
const UpdateInstanceBadRequestCode int = 400

/*UpdateInstanceBadRequest Bad request

swagger:response updateInstanceBadRequest
*/
type UpdateInstanceBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateInstanceBadRequest creates UpdateInstanceBadRequest with default headers values
func NewUpdateInstanceBadRequest() *UpdateInstanceBadRequest {

	return &UpdateInstanceBadRequest{}
}

// WithXRequestID adds the xRequestId to the update instance bad request response
func (o *UpdateInstanceBadRequest) WithXRequestID(xRequestID string) *UpdateInstanceBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update instance bad request response
func (o *UpdateInstanceBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update instance bad request response
func (o *UpdateInstanceBadRequest) WithPayload(payload *models.Errors) *UpdateInstanceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update instance bad request response
func (o *UpdateInstanceBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInstanceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateInstanceUnauthorizedCode is the HTTP code returned for type UpdateInstanceUnauthorized
const UpdateInstanceUnauthorizedCode int = 401

/*UpdateInstanceUnauthorized Unauthorized

swagger:response updateInstanceUnauthorized
*/
type UpdateInstanceUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateInstanceUnauthorized creates UpdateInstanceUnauthorized with default headers values
func NewUpdateInstanceUnauthorized() *UpdateInstanceUnauthorized {

	return &UpdateInstanceUnauthorized{}
}

// WithXRequestID adds the xRequestId to the update instance unauthorized response
func (o *UpdateInstanceUnauthorized) WithXRequestID(xRequestID string) *UpdateInstanceUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update instance unauthorized response
func (o *UpdateInstanceUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update instance unauthorized response
func (o *UpdateInstanceUnauthorized) WithPayload(payload *models.Errors) *UpdateInstanceUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update instance unauthorized response
func (o *UpdateInstanceUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInstanceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateInstanceForbiddenCode is the HTTP code returned for type UpdateInstanceForbidden
const UpdateInstanceForbiddenCode int = 403

/*UpdateInstanceForbidden Forbidden

swagger:response updateInstanceForbidden
*/
type UpdateInstanceForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateInstanceForbidden creates UpdateInstanceForbidden with default headers values
func NewUpdateInstanceForbidden() *UpdateInstanceForbidden {

	return &UpdateInstanceForbidden{}
}

// WithXRequestID adds the xRequestId to the update instance forbidden response
func (o *UpdateInstanceForbidden) WithXRequestID(xRequestID string) *UpdateInstanceForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update instance forbidden response
func (o *UpdateInstanceForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update instance forbidden response
func (o *UpdateInstanceForbidden) WithPayload(payload *models.Errors) *UpdateInstanceForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update instance forbidden response
func (o *UpdateInstanceForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInstanceForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateInstanceNotFoundCode is the HTTP code returned for type UpdateInstanceNotFound
const UpdateInstanceNotFoundCode int = 404

/*UpdateInstanceNotFound Not found

swagger:response updateInstanceNotFound
*/
type UpdateInstanceNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateInstanceNotFound creates UpdateInstanceNotFound with default headers values
func NewUpdateInstanceNotFound() *UpdateInstanceNotFound {

	return &UpdateInstanceNotFound{}
}

// WithXRequestID adds the xRequestId to the update instance not found response
func (o *UpdateInstanceNotFound) WithXRequestID(xRequestID string) *UpdateInstanceNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update instance not found response
func (o *UpdateInstanceNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update instance not found response
func (o *UpdateInstanceNotFound) WithPayload(payload *models.Errors) *UpdateInstanceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update instance not found response
func (o *UpdateInstanceNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInstanceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateInstanceInternalServerErrorCode is the HTTP code returned for type UpdateInstanceInternalServerError
const UpdateInstanceInternalServerErrorCode int = 500

/*UpdateInstanceInternalServerError Internal server error

swagger:response updateInstanceInternalServerError
*/
type UpdateInstanceInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateInstanceInternalServerError creates UpdateInstanceInternalServerError with default headers values
func NewUpdateInstanceInternalServerError() *UpdateInstanceInternalServerError {

	return &UpdateInstanceInternalServerError{}
}

// WithXRequestID adds the xRequestId to the update instance internal server error response
func (o *UpdateInstanceInternalServerError) WithXRequestID(xRequestID string) *UpdateInstanceInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update instance internal server error response
func (o *UpdateInstanceInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update instance internal server error response
func (o *UpdateInstanceInternalServerError) WithPayload(payload *models.Errors) *UpdateInstanceInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update instance internal server error response
func (o *UpdateInstanceInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateInstanceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
