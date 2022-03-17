// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// StopReplicationOKCode is the HTTP code returned for type StopReplicationOK
const StopReplicationOKCode int = 200

/*StopReplicationOK Success

swagger:response stopReplicationOK
*/
type StopReplicationOK struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewStopReplicationOK creates StopReplicationOK with default headers values
func NewStopReplicationOK() *StopReplicationOK {

	return &StopReplicationOK{}
}

// WithXRequestID adds the xRequestId to the stop replication o k response
func (o *StopReplicationOK) WithXRequestID(xRequestID string) *StopReplicationOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the stop replication o k response
func (o *StopReplicationOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *StopReplicationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// StopReplicationUnauthorizedCode is the HTTP code returned for type StopReplicationUnauthorized
const StopReplicationUnauthorizedCode int = 401

/*StopReplicationUnauthorized Unauthorized

swagger:response stopReplicationUnauthorized
*/
type StopReplicationUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewStopReplicationUnauthorized creates StopReplicationUnauthorized with default headers values
func NewStopReplicationUnauthorized() *StopReplicationUnauthorized {

	return &StopReplicationUnauthorized{}
}

// WithXRequestID adds the xRequestId to the stop replication unauthorized response
func (o *StopReplicationUnauthorized) WithXRequestID(xRequestID string) *StopReplicationUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the stop replication unauthorized response
func (o *StopReplicationUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the stop replication unauthorized response
func (o *StopReplicationUnauthorized) WithPayload(payload *models.Errors) *StopReplicationUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop replication unauthorized response
func (o *StopReplicationUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopReplicationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// StopReplicationForbiddenCode is the HTTP code returned for type StopReplicationForbidden
const StopReplicationForbiddenCode int = 403

/*StopReplicationForbidden Forbidden

swagger:response stopReplicationForbidden
*/
type StopReplicationForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewStopReplicationForbidden creates StopReplicationForbidden with default headers values
func NewStopReplicationForbidden() *StopReplicationForbidden {

	return &StopReplicationForbidden{}
}

// WithXRequestID adds the xRequestId to the stop replication forbidden response
func (o *StopReplicationForbidden) WithXRequestID(xRequestID string) *StopReplicationForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the stop replication forbidden response
func (o *StopReplicationForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the stop replication forbidden response
func (o *StopReplicationForbidden) WithPayload(payload *models.Errors) *StopReplicationForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop replication forbidden response
func (o *StopReplicationForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopReplicationForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// StopReplicationNotFoundCode is the HTTP code returned for type StopReplicationNotFound
const StopReplicationNotFoundCode int = 404

/*StopReplicationNotFound Not found

swagger:response stopReplicationNotFound
*/
type StopReplicationNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewStopReplicationNotFound creates StopReplicationNotFound with default headers values
func NewStopReplicationNotFound() *StopReplicationNotFound {

	return &StopReplicationNotFound{}
}

// WithXRequestID adds the xRequestId to the stop replication not found response
func (o *StopReplicationNotFound) WithXRequestID(xRequestID string) *StopReplicationNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the stop replication not found response
func (o *StopReplicationNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the stop replication not found response
func (o *StopReplicationNotFound) WithPayload(payload *models.Errors) *StopReplicationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop replication not found response
func (o *StopReplicationNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopReplicationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// StopReplicationInternalServerErrorCode is the HTTP code returned for type StopReplicationInternalServerError
const StopReplicationInternalServerErrorCode int = 500

/*StopReplicationInternalServerError Internal server error

swagger:response stopReplicationInternalServerError
*/
type StopReplicationInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewStopReplicationInternalServerError creates StopReplicationInternalServerError with default headers values
func NewStopReplicationInternalServerError() *StopReplicationInternalServerError {

	return &StopReplicationInternalServerError{}
}

// WithXRequestID adds the xRequestId to the stop replication internal server error response
func (o *StopReplicationInternalServerError) WithXRequestID(xRequestID string) *StopReplicationInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the stop replication internal server error response
func (o *StopReplicationInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the stop replication internal server error response
func (o *StopReplicationInternalServerError) WithPayload(payload *models.Errors) *StopReplicationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop replication internal server error response
func (o *StopReplicationInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopReplicationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
