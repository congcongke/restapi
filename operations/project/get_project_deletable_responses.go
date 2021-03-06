// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetProjectDeletableOKCode is the HTTP code returned for type GetProjectDeletableOK
const GetProjectDeletableOKCode int = 200

/*GetProjectDeletableOK Return deletable status of the project.

swagger:response getProjectDeletableOK
*/
type GetProjectDeletableOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProjectDeletable `json:"body,omitempty"`
}

// NewGetProjectDeletableOK creates GetProjectDeletableOK with default headers values
func NewGetProjectDeletableOK() *GetProjectDeletableOK {

	return &GetProjectDeletableOK{}
}

// WithPayload adds the payload to the get project deletable o k response
func (o *GetProjectDeletableOK) WithPayload(payload *models.ProjectDeletable) *GetProjectDeletableOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project deletable o k response
func (o *GetProjectDeletableOK) SetPayload(payload *models.ProjectDeletable) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectDeletableOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProjectDeletableUnauthorizedCode is the HTTP code returned for type GetProjectDeletableUnauthorized
const GetProjectDeletableUnauthorizedCode int = 401

/*GetProjectDeletableUnauthorized Unauthorized

swagger:response getProjectDeletableUnauthorized
*/
type GetProjectDeletableUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectDeletableUnauthorized creates GetProjectDeletableUnauthorized with default headers values
func NewGetProjectDeletableUnauthorized() *GetProjectDeletableUnauthorized {

	return &GetProjectDeletableUnauthorized{}
}

// WithXRequestID adds the xRequestId to the get project deletable unauthorized response
func (o *GetProjectDeletableUnauthorized) WithXRequestID(xRequestID string) *GetProjectDeletableUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project deletable unauthorized response
func (o *GetProjectDeletableUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project deletable unauthorized response
func (o *GetProjectDeletableUnauthorized) WithPayload(payload *models.Errors) *GetProjectDeletableUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project deletable unauthorized response
func (o *GetProjectDeletableUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectDeletableUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectDeletableForbiddenCode is the HTTP code returned for type GetProjectDeletableForbidden
const GetProjectDeletableForbiddenCode int = 403

/*GetProjectDeletableForbidden Forbidden

swagger:response getProjectDeletableForbidden
*/
type GetProjectDeletableForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectDeletableForbidden creates GetProjectDeletableForbidden with default headers values
func NewGetProjectDeletableForbidden() *GetProjectDeletableForbidden {

	return &GetProjectDeletableForbidden{}
}

// WithXRequestID adds the xRequestId to the get project deletable forbidden response
func (o *GetProjectDeletableForbidden) WithXRequestID(xRequestID string) *GetProjectDeletableForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project deletable forbidden response
func (o *GetProjectDeletableForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project deletable forbidden response
func (o *GetProjectDeletableForbidden) WithPayload(payload *models.Errors) *GetProjectDeletableForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project deletable forbidden response
func (o *GetProjectDeletableForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectDeletableForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectDeletableNotFoundCode is the HTTP code returned for type GetProjectDeletableNotFound
const GetProjectDeletableNotFoundCode int = 404

/*GetProjectDeletableNotFound Not found

swagger:response getProjectDeletableNotFound
*/
type GetProjectDeletableNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectDeletableNotFound creates GetProjectDeletableNotFound with default headers values
func NewGetProjectDeletableNotFound() *GetProjectDeletableNotFound {

	return &GetProjectDeletableNotFound{}
}

// WithXRequestID adds the xRequestId to the get project deletable not found response
func (o *GetProjectDeletableNotFound) WithXRequestID(xRequestID string) *GetProjectDeletableNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project deletable not found response
func (o *GetProjectDeletableNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project deletable not found response
func (o *GetProjectDeletableNotFound) WithPayload(payload *models.Errors) *GetProjectDeletableNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project deletable not found response
func (o *GetProjectDeletableNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectDeletableNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetProjectDeletableInternalServerErrorCode is the HTTP code returned for type GetProjectDeletableInternalServerError
const GetProjectDeletableInternalServerErrorCode int = 500

/*GetProjectDeletableInternalServerError Internal server error

swagger:response getProjectDeletableInternalServerError
*/
type GetProjectDeletableInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetProjectDeletableInternalServerError creates GetProjectDeletableInternalServerError with default headers values
func NewGetProjectDeletableInternalServerError() *GetProjectDeletableInternalServerError {

	return &GetProjectDeletableInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get project deletable internal server error response
func (o *GetProjectDeletableInternalServerError) WithXRequestID(xRequestID string) *GetProjectDeletableInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get project deletable internal server error response
func (o *GetProjectDeletableInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get project deletable internal server error response
func (o *GetProjectDeletableInternalServerError) WithPayload(payload *models.Errors) *GetProjectDeletableInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project deletable internal server error response
func (o *GetProjectDeletableInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectDeletableInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
