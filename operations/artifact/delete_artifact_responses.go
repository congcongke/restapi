// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// DeleteArtifactOKCode is the HTTP code returned for type DeleteArtifactOK
const DeleteArtifactOKCode int = 200

/*DeleteArtifactOK Success

swagger:response deleteArtifactOK
*/
type DeleteArtifactOK struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewDeleteArtifactOK creates DeleteArtifactOK with default headers values
func NewDeleteArtifactOK() *DeleteArtifactOK {

	return &DeleteArtifactOK{}
}

// WithXRequestID adds the xRequestId to the delete artifact o k response
func (o *DeleteArtifactOK) WithXRequestID(xRequestID string) *DeleteArtifactOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete artifact o k response
func (o *DeleteArtifactOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *DeleteArtifactOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteArtifactUnauthorizedCode is the HTTP code returned for type DeleteArtifactUnauthorized
const DeleteArtifactUnauthorizedCode int = 401

/*DeleteArtifactUnauthorized Unauthorized

swagger:response deleteArtifactUnauthorized
*/
type DeleteArtifactUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteArtifactUnauthorized creates DeleteArtifactUnauthorized with default headers values
func NewDeleteArtifactUnauthorized() *DeleteArtifactUnauthorized {

	return &DeleteArtifactUnauthorized{}
}

// WithXRequestID adds the xRequestId to the delete artifact unauthorized response
func (o *DeleteArtifactUnauthorized) WithXRequestID(xRequestID string) *DeleteArtifactUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete artifact unauthorized response
func (o *DeleteArtifactUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete artifact unauthorized response
func (o *DeleteArtifactUnauthorized) WithPayload(payload *models.Errors) *DeleteArtifactUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete artifact unauthorized response
func (o *DeleteArtifactUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteArtifactUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// DeleteArtifactForbiddenCode is the HTTP code returned for type DeleteArtifactForbidden
const DeleteArtifactForbiddenCode int = 403

/*DeleteArtifactForbidden Forbidden

swagger:response deleteArtifactForbidden
*/
type DeleteArtifactForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteArtifactForbidden creates DeleteArtifactForbidden with default headers values
func NewDeleteArtifactForbidden() *DeleteArtifactForbidden {

	return &DeleteArtifactForbidden{}
}

// WithXRequestID adds the xRequestId to the delete artifact forbidden response
func (o *DeleteArtifactForbidden) WithXRequestID(xRequestID string) *DeleteArtifactForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete artifact forbidden response
func (o *DeleteArtifactForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete artifact forbidden response
func (o *DeleteArtifactForbidden) WithPayload(payload *models.Errors) *DeleteArtifactForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete artifact forbidden response
func (o *DeleteArtifactForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteArtifactForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// DeleteArtifactNotFoundCode is the HTTP code returned for type DeleteArtifactNotFound
const DeleteArtifactNotFoundCode int = 404

/*DeleteArtifactNotFound Not found

swagger:response deleteArtifactNotFound
*/
type DeleteArtifactNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteArtifactNotFound creates DeleteArtifactNotFound with default headers values
func NewDeleteArtifactNotFound() *DeleteArtifactNotFound {

	return &DeleteArtifactNotFound{}
}

// WithXRequestID adds the xRequestId to the delete artifact not found response
func (o *DeleteArtifactNotFound) WithXRequestID(xRequestID string) *DeleteArtifactNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete artifact not found response
func (o *DeleteArtifactNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete artifact not found response
func (o *DeleteArtifactNotFound) WithPayload(payload *models.Errors) *DeleteArtifactNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete artifact not found response
func (o *DeleteArtifactNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteArtifactNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// DeleteArtifactInternalServerErrorCode is the HTTP code returned for type DeleteArtifactInternalServerError
const DeleteArtifactInternalServerErrorCode int = 500

/*DeleteArtifactInternalServerError Internal server error

swagger:response deleteArtifactInternalServerError
*/
type DeleteArtifactInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewDeleteArtifactInternalServerError creates DeleteArtifactInternalServerError with default headers values
func NewDeleteArtifactInternalServerError() *DeleteArtifactInternalServerError {

	return &DeleteArtifactInternalServerError{}
}

// WithXRequestID adds the xRequestId to the delete artifact internal server error response
func (o *DeleteArtifactInternalServerError) WithXRequestID(xRequestID string) *DeleteArtifactInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the delete artifact internal server error response
func (o *DeleteArtifactInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the delete artifact internal server error response
func (o *DeleteArtifactInternalServerError) WithPayload(payload *models.Errors) *DeleteArtifactInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete artifact internal server error response
func (o *DeleteArtifactInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteArtifactInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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