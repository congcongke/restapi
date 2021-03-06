// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// ManualPreheatCreatedCode is the HTTP code returned for type ManualPreheatCreated
const ManualPreheatCreatedCode int = 201

/*ManualPreheatCreated Created

swagger:response manualPreheatCreated
*/
type ManualPreheatCreated struct {
	/*The location of the resource

	 */
	Location string `json:"Location"`
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewManualPreheatCreated creates ManualPreheatCreated with default headers values
func NewManualPreheatCreated() *ManualPreheatCreated {

	return &ManualPreheatCreated{}
}

// WithLocation adds the location to the manual preheat created response
func (o *ManualPreheatCreated) WithLocation(location string) *ManualPreheatCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the manual preheat created response
func (o *ManualPreheatCreated) SetLocation(location string) {
	o.Location = location
}

// WithXRequestID adds the xRequestId to the manual preheat created response
func (o *ManualPreheatCreated) WithXRequestID(xRequestID string) *ManualPreheatCreated {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the manual preheat created response
func (o *ManualPreheatCreated) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *ManualPreheatCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ManualPreheatBadRequestCode is the HTTP code returned for type ManualPreheatBadRequest
const ManualPreheatBadRequestCode int = 400

/*ManualPreheatBadRequest Bad request

swagger:response manualPreheatBadRequest
*/
type ManualPreheatBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewManualPreheatBadRequest creates ManualPreheatBadRequest with default headers values
func NewManualPreheatBadRequest() *ManualPreheatBadRequest {

	return &ManualPreheatBadRequest{}
}

// WithXRequestID adds the xRequestId to the manual preheat bad request response
func (o *ManualPreheatBadRequest) WithXRequestID(xRequestID string) *ManualPreheatBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the manual preheat bad request response
func (o *ManualPreheatBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the manual preheat bad request response
func (o *ManualPreheatBadRequest) WithPayload(payload *models.Errors) *ManualPreheatBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the manual preheat bad request response
func (o *ManualPreheatBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ManualPreheatBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ManualPreheatUnauthorizedCode is the HTTP code returned for type ManualPreheatUnauthorized
const ManualPreheatUnauthorizedCode int = 401

/*ManualPreheatUnauthorized Unauthorized

swagger:response manualPreheatUnauthorized
*/
type ManualPreheatUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewManualPreheatUnauthorized creates ManualPreheatUnauthorized with default headers values
func NewManualPreheatUnauthorized() *ManualPreheatUnauthorized {

	return &ManualPreheatUnauthorized{}
}

// WithXRequestID adds the xRequestId to the manual preheat unauthorized response
func (o *ManualPreheatUnauthorized) WithXRequestID(xRequestID string) *ManualPreheatUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the manual preheat unauthorized response
func (o *ManualPreheatUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the manual preheat unauthorized response
func (o *ManualPreheatUnauthorized) WithPayload(payload *models.Errors) *ManualPreheatUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the manual preheat unauthorized response
func (o *ManualPreheatUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ManualPreheatUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ManualPreheatForbiddenCode is the HTTP code returned for type ManualPreheatForbidden
const ManualPreheatForbiddenCode int = 403

/*ManualPreheatForbidden Forbidden

swagger:response manualPreheatForbidden
*/
type ManualPreheatForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewManualPreheatForbidden creates ManualPreheatForbidden with default headers values
func NewManualPreheatForbidden() *ManualPreheatForbidden {

	return &ManualPreheatForbidden{}
}

// WithXRequestID adds the xRequestId to the manual preheat forbidden response
func (o *ManualPreheatForbidden) WithXRequestID(xRequestID string) *ManualPreheatForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the manual preheat forbidden response
func (o *ManualPreheatForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the manual preheat forbidden response
func (o *ManualPreheatForbidden) WithPayload(payload *models.Errors) *ManualPreheatForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the manual preheat forbidden response
func (o *ManualPreheatForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ManualPreheatForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ManualPreheatNotFoundCode is the HTTP code returned for type ManualPreheatNotFound
const ManualPreheatNotFoundCode int = 404

/*ManualPreheatNotFound Not found

swagger:response manualPreheatNotFound
*/
type ManualPreheatNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewManualPreheatNotFound creates ManualPreheatNotFound with default headers values
func NewManualPreheatNotFound() *ManualPreheatNotFound {

	return &ManualPreheatNotFound{}
}

// WithXRequestID adds the xRequestId to the manual preheat not found response
func (o *ManualPreheatNotFound) WithXRequestID(xRequestID string) *ManualPreheatNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the manual preheat not found response
func (o *ManualPreheatNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the manual preheat not found response
func (o *ManualPreheatNotFound) WithPayload(payload *models.Errors) *ManualPreheatNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the manual preheat not found response
func (o *ManualPreheatNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ManualPreheatNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ManualPreheatInternalServerErrorCode is the HTTP code returned for type ManualPreheatInternalServerError
const ManualPreheatInternalServerErrorCode int = 500

/*ManualPreheatInternalServerError Internal server error

swagger:response manualPreheatInternalServerError
*/
type ManualPreheatInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewManualPreheatInternalServerError creates ManualPreheatInternalServerError with default headers values
func NewManualPreheatInternalServerError() *ManualPreheatInternalServerError {

	return &ManualPreheatInternalServerError{}
}

// WithXRequestID adds the xRequestId to the manual preheat internal server error response
func (o *ManualPreheatInternalServerError) WithXRequestID(xRequestID string) *ManualPreheatInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the manual preheat internal server error response
func (o *ManualPreheatInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the manual preheat internal server error response
func (o *ManualPreheatInternalServerError) WithPayload(payload *models.Errors) *ManualPreheatInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the manual preheat internal server error response
func (o *ManualPreheatInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ManualPreheatInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
