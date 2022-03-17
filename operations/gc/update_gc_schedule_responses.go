// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// UpdateGCScheduleOKCode is the HTTP code returned for type UpdateGCScheduleOK
const UpdateGCScheduleOKCode int = 200

/*UpdateGCScheduleOK Updated gc's schedule successfully.

swagger:response updateGcScheduleOK
*/
type UpdateGCScheduleOK struct {
}

// NewUpdateGCScheduleOK creates UpdateGCScheduleOK with default headers values
func NewUpdateGCScheduleOK() *UpdateGCScheduleOK {

	return &UpdateGCScheduleOK{}
}

// WriteResponse to the client
func (o *UpdateGCScheduleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateGCScheduleBadRequestCode is the HTTP code returned for type UpdateGCScheduleBadRequest
const UpdateGCScheduleBadRequestCode int = 400

/*UpdateGCScheduleBadRequest Bad request

swagger:response updateGcScheduleBadRequest
*/
type UpdateGCScheduleBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateGCScheduleBadRequest creates UpdateGCScheduleBadRequest with default headers values
func NewUpdateGCScheduleBadRequest() *UpdateGCScheduleBadRequest {

	return &UpdateGCScheduleBadRequest{}
}

// WithXRequestID adds the xRequestId to the update Gc schedule bad request response
func (o *UpdateGCScheduleBadRequest) WithXRequestID(xRequestID string) *UpdateGCScheduleBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update Gc schedule bad request response
func (o *UpdateGCScheduleBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update Gc schedule bad request response
func (o *UpdateGCScheduleBadRequest) WithPayload(payload *models.Errors) *UpdateGCScheduleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Gc schedule bad request response
func (o *UpdateGCScheduleBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGCScheduleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateGCScheduleUnauthorizedCode is the HTTP code returned for type UpdateGCScheduleUnauthorized
const UpdateGCScheduleUnauthorizedCode int = 401

/*UpdateGCScheduleUnauthorized Unauthorized

swagger:response updateGcScheduleUnauthorized
*/
type UpdateGCScheduleUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateGCScheduleUnauthorized creates UpdateGCScheduleUnauthorized with default headers values
func NewUpdateGCScheduleUnauthorized() *UpdateGCScheduleUnauthorized {

	return &UpdateGCScheduleUnauthorized{}
}

// WithXRequestID adds the xRequestId to the update Gc schedule unauthorized response
func (o *UpdateGCScheduleUnauthorized) WithXRequestID(xRequestID string) *UpdateGCScheduleUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update Gc schedule unauthorized response
func (o *UpdateGCScheduleUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update Gc schedule unauthorized response
func (o *UpdateGCScheduleUnauthorized) WithPayload(payload *models.Errors) *UpdateGCScheduleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Gc schedule unauthorized response
func (o *UpdateGCScheduleUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGCScheduleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateGCScheduleForbiddenCode is the HTTP code returned for type UpdateGCScheduleForbidden
const UpdateGCScheduleForbiddenCode int = 403

/*UpdateGCScheduleForbidden Forbidden

swagger:response updateGcScheduleForbidden
*/
type UpdateGCScheduleForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateGCScheduleForbidden creates UpdateGCScheduleForbidden with default headers values
func NewUpdateGCScheduleForbidden() *UpdateGCScheduleForbidden {

	return &UpdateGCScheduleForbidden{}
}

// WithXRequestID adds the xRequestId to the update Gc schedule forbidden response
func (o *UpdateGCScheduleForbidden) WithXRequestID(xRequestID string) *UpdateGCScheduleForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update Gc schedule forbidden response
func (o *UpdateGCScheduleForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update Gc schedule forbidden response
func (o *UpdateGCScheduleForbidden) WithPayload(payload *models.Errors) *UpdateGCScheduleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Gc schedule forbidden response
func (o *UpdateGCScheduleForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGCScheduleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateGCScheduleInternalServerErrorCode is the HTTP code returned for type UpdateGCScheduleInternalServerError
const UpdateGCScheduleInternalServerErrorCode int = 500

/*UpdateGCScheduleInternalServerError Internal server error

swagger:response updateGcScheduleInternalServerError
*/
type UpdateGCScheduleInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateGCScheduleInternalServerError creates UpdateGCScheduleInternalServerError with default headers values
func NewUpdateGCScheduleInternalServerError() *UpdateGCScheduleInternalServerError {

	return &UpdateGCScheduleInternalServerError{}
}

// WithXRequestID adds the xRequestId to the update Gc schedule internal server error response
func (o *UpdateGCScheduleInternalServerError) WithXRequestID(xRequestID string) *UpdateGCScheduleInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update Gc schedule internal server error response
func (o *UpdateGCScheduleInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update Gc schedule internal server error response
func (o *UpdateGCScheduleInternalServerError) WithPayload(payload *models.Errors) *UpdateGCScheduleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update Gc schedule internal server error response
func (o *UpdateGCScheduleInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGCScheduleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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