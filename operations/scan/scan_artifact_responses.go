// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// ScanArtifactAcceptedCode is the HTTP code returned for type ScanArtifactAccepted
const ScanArtifactAcceptedCode int = 202

/*ScanArtifactAccepted Accepted

swagger:response scanArtifactAccepted
*/
type ScanArtifactAccepted struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewScanArtifactAccepted creates ScanArtifactAccepted with default headers values
func NewScanArtifactAccepted() *ScanArtifactAccepted {

	return &ScanArtifactAccepted{}
}

// WithXRequestID adds the xRequestId to the scan artifact accepted response
func (o *ScanArtifactAccepted) WithXRequestID(xRequestID string) *ScanArtifactAccepted {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the scan artifact accepted response
func (o *ScanArtifactAccepted) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *ScanArtifactAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// ScanArtifactBadRequestCode is the HTTP code returned for type ScanArtifactBadRequest
const ScanArtifactBadRequestCode int = 400

/*ScanArtifactBadRequest Bad request

swagger:response scanArtifactBadRequest
*/
type ScanArtifactBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewScanArtifactBadRequest creates ScanArtifactBadRequest with default headers values
func NewScanArtifactBadRequest() *ScanArtifactBadRequest {

	return &ScanArtifactBadRequest{}
}

// WithXRequestID adds the xRequestId to the scan artifact bad request response
func (o *ScanArtifactBadRequest) WithXRequestID(xRequestID string) *ScanArtifactBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the scan artifact bad request response
func (o *ScanArtifactBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the scan artifact bad request response
func (o *ScanArtifactBadRequest) WithPayload(payload *models.Errors) *ScanArtifactBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the scan artifact bad request response
func (o *ScanArtifactBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ScanArtifactBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ScanArtifactUnauthorizedCode is the HTTP code returned for type ScanArtifactUnauthorized
const ScanArtifactUnauthorizedCode int = 401

/*ScanArtifactUnauthorized Unauthorized

swagger:response scanArtifactUnauthorized
*/
type ScanArtifactUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewScanArtifactUnauthorized creates ScanArtifactUnauthorized with default headers values
func NewScanArtifactUnauthorized() *ScanArtifactUnauthorized {

	return &ScanArtifactUnauthorized{}
}

// WithXRequestID adds the xRequestId to the scan artifact unauthorized response
func (o *ScanArtifactUnauthorized) WithXRequestID(xRequestID string) *ScanArtifactUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the scan artifact unauthorized response
func (o *ScanArtifactUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the scan artifact unauthorized response
func (o *ScanArtifactUnauthorized) WithPayload(payload *models.Errors) *ScanArtifactUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the scan artifact unauthorized response
func (o *ScanArtifactUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ScanArtifactUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ScanArtifactForbiddenCode is the HTTP code returned for type ScanArtifactForbidden
const ScanArtifactForbiddenCode int = 403

/*ScanArtifactForbidden Forbidden

swagger:response scanArtifactForbidden
*/
type ScanArtifactForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewScanArtifactForbidden creates ScanArtifactForbidden with default headers values
func NewScanArtifactForbidden() *ScanArtifactForbidden {

	return &ScanArtifactForbidden{}
}

// WithXRequestID adds the xRequestId to the scan artifact forbidden response
func (o *ScanArtifactForbidden) WithXRequestID(xRequestID string) *ScanArtifactForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the scan artifact forbidden response
func (o *ScanArtifactForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the scan artifact forbidden response
func (o *ScanArtifactForbidden) WithPayload(payload *models.Errors) *ScanArtifactForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the scan artifact forbidden response
func (o *ScanArtifactForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ScanArtifactForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ScanArtifactNotFoundCode is the HTTP code returned for type ScanArtifactNotFound
const ScanArtifactNotFoundCode int = 404

/*ScanArtifactNotFound Not found

swagger:response scanArtifactNotFound
*/
type ScanArtifactNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewScanArtifactNotFound creates ScanArtifactNotFound with default headers values
func NewScanArtifactNotFound() *ScanArtifactNotFound {

	return &ScanArtifactNotFound{}
}

// WithXRequestID adds the xRequestId to the scan artifact not found response
func (o *ScanArtifactNotFound) WithXRequestID(xRequestID string) *ScanArtifactNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the scan artifact not found response
func (o *ScanArtifactNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the scan artifact not found response
func (o *ScanArtifactNotFound) WithPayload(payload *models.Errors) *ScanArtifactNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the scan artifact not found response
func (o *ScanArtifactNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ScanArtifactNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ScanArtifactInternalServerErrorCode is the HTTP code returned for type ScanArtifactInternalServerError
const ScanArtifactInternalServerErrorCode int = 500

/*ScanArtifactInternalServerError Internal server error

swagger:response scanArtifactInternalServerError
*/
type ScanArtifactInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewScanArtifactInternalServerError creates ScanArtifactInternalServerError with default headers values
func NewScanArtifactInternalServerError() *ScanArtifactInternalServerError {

	return &ScanArtifactInternalServerError{}
}

// WithXRequestID adds the xRequestId to the scan artifact internal server error response
func (o *ScanArtifactInternalServerError) WithXRequestID(xRequestID string) *ScanArtifactInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the scan artifact internal server error response
func (o *ScanArtifactInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the scan artifact internal server error response
func (o *ScanArtifactInternalServerError) WithPayload(payload *models.Errors) *ScanArtifactInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the scan artifact internal server error response
func (o *ScanArtifactInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ScanArtifactInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
