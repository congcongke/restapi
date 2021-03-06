// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// OperateRetentionExecutionOKCode is the HTTP code returned for type OperateRetentionExecutionOK
const OperateRetentionExecutionOKCode int = 200

/*OperateRetentionExecutionOK Stop a Retention job successfully.

swagger:response operateRetentionExecutionOK
*/
type OperateRetentionExecutionOK struct {
}

// NewOperateRetentionExecutionOK creates OperateRetentionExecutionOK with default headers values
func NewOperateRetentionExecutionOK() *OperateRetentionExecutionOK {

	return &OperateRetentionExecutionOK{}
}

// WriteResponse to the client
func (o *OperateRetentionExecutionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// OperateRetentionExecutionUnauthorizedCode is the HTTP code returned for type OperateRetentionExecutionUnauthorized
const OperateRetentionExecutionUnauthorizedCode int = 401

/*OperateRetentionExecutionUnauthorized Unauthorized

swagger:response operateRetentionExecutionUnauthorized
*/
type OperateRetentionExecutionUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewOperateRetentionExecutionUnauthorized creates OperateRetentionExecutionUnauthorized with default headers values
func NewOperateRetentionExecutionUnauthorized() *OperateRetentionExecutionUnauthorized {

	return &OperateRetentionExecutionUnauthorized{}
}

// WithXRequestID adds the xRequestId to the operate retention execution unauthorized response
func (o *OperateRetentionExecutionUnauthorized) WithXRequestID(xRequestID string) *OperateRetentionExecutionUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the operate retention execution unauthorized response
func (o *OperateRetentionExecutionUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the operate retention execution unauthorized response
func (o *OperateRetentionExecutionUnauthorized) WithPayload(payload *models.Errors) *OperateRetentionExecutionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the operate retention execution unauthorized response
func (o *OperateRetentionExecutionUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OperateRetentionExecutionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// OperateRetentionExecutionForbiddenCode is the HTTP code returned for type OperateRetentionExecutionForbidden
const OperateRetentionExecutionForbiddenCode int = 403

/*OperateRetentionExecutionForbidden Forbidden

swagger:response operateRetentionExecutionForbidden
*/
type OperateRetentionExecutionForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewOperateRetentionExecutionForbidden creates OperateRetentionExecutionForbidden with default headers values
func NewOperateRetentionExecutionForbidden() *OperateRetentionExecutionForbidden {

	return &OperateRetentionExecutionForbidden{}
}

// WithXRequestID adds the xRequestId to the operate retention execution forbidden response
func (o *OperateRetentionExecutionForbidden) WithXRequestID(xRequestID string) *OperateRetentionExecutionForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the operate retention execution forbidden response
func (o *OperateRetentionExecutionForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the operate retention execution forbidden response
func (o *OperateRetentionExecutionForbidden) WithPayload(payload *models.Errors) *OperateRetentionExecutionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the operate retention execution forbidden response
func (o *OperateRetentionExecutionForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OperateRetentionExecutionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// OperateRetentionExecutionInternalServerErrorCode is the HTTP code returned for type OperateRetentionExecutionInternalServerError
const OperateRetentionExecutionInternalServerErrorCode int = 500

/*OperateRetentionExecutionInternalServerError Internal server error

swagger:response operateRetentionExecutionInternalServerError
*/
type OperateRetentionExecutionInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewOperateRetentionExecutionInternalServerError creates OperateRetentionExecutionInternalServerError with default headers values
func NewOperateRetentionExecutionInternalServerError() *OperateRetentionExecutionInternalServerError {

	return &OperateRetentionExecutionInternalServerError{}
}

// WithXRequestID adds the xRequestId to the operate retention execution internal server error response
func (o *OperateRetentionExecutionInternalServerError) WithXRequestID(xRequestID string) *OperateRetentionExecutionInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the operate retention execution internal server error response
func (o *OperateRetentionExecutionInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the operate retention execution internal server error response
func (o *OperateRetentionExecutionInternalServerError) WithPayload(payload *models.Errors) *OperateRetentionExecutionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the operate retention execution internal server error response
func (o *OperateRetentionExecutionInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *OperateRetentionExecutionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
