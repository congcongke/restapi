// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetRepositoryOKCode is the HTTP code returned for type GetRepositoryOK
const GetRepositoryOKCode int = 200

/*GetRepositoryOK Success

swagger:response getRepositoryOK
*/
type GetRepositoryOK struct {

	/*
	  In: Body
	*/
	Payload *models.Repository `json:"body,omitempty"`
}

// NewGetRepositoryOK creates GetRepositoryOK with default headers values
func NewGetRepositoryOK() *GetRepositoryOK {

	return &GetRepositoryOK{}
}

// WithPayload adds the payload to the get repository o k response
func (o *GetRepositoryOK) WithPayload(payload *models.Repository) *GetRepositoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repository o k response
func (o *GetRepositoryOK) SetPayload(payload *models.Repository) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRepositoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRepositoryBadRequestCode is the HTTP code returned for type GetRepositoryBadRequest
const GetRepositoryBadRequestCode int = 400

/*GetRepositoryBadRequest Bad request

swagger:response getRepositoryBadRequest
*/
type GetRepositoryBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRepositoryBadRequest creates GetRepositoryBadRequest with default headers values
func NewGetRepositoryBadRequest() *GetRepositoryBadRequest {

	return &GetRepositoryBadRequest{}
}

// WithXRequestID adds the xRequestId to the get repository bad request response
func (o *GetRepositoryBadRequest) WithXRequestID(xRequestID string) *GetRepositoryBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get repository bad request response
func (o *GetRepositoryBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get repository bad request response
func (o *GetRepositoryBadRequest) WithPayload(payload *models.Errors) *GetRepositoryBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repository bad request response
func (o *GetRepositoryBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRepositoryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetRepositoryUnauthorizedCode is the HTTP code returned for type GetRepositoryUnauthorized
const GetRepositoryUnauthorizedCode int = 401

/*GetRepositoryUnauthorized Unauthorized

swagger:response getRepositoryUnauthorized
*/
type GetRepositoryUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRepositoryUnauthorized creates GetRepositoryUnauthorized with default headers values
func NewGetRepositoryUnauthorized() *GetRepositoryUnauthorized {

	return &GetRepositoryUnauthorized{}
}

// WithXRequestID adds the xRequestId to the get repository unauthorized response
func (o *GetRepositoryUnauthorized) WithXRequestID(xRequestID string) *GetRepositoryUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get repository unauthorized response
func (o *GetRepositoryUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get repository unauthorized response
func (o *GetRepositoryUnauthorized) WithPayload(payload *models.Errors) *GetRepositoryUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repository unauthorized response
func (o *GetRepositoryUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRepositoryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetRepositoryForbiddenCode is the HTTP code returned for type GetRepositoryForbidden
const GetRepositoryForbiddenCode int = 403

/*GetRepositoryForbidden Forbidden

swagger:response getRepositoryForbidden
*/
type GetRepositoryForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRepositoryForbidden creates GetRepositoryForbidden with default headers values
func NewGetRepositoryForbidden() *GetRepositoryForbidden {

	return &GetRepositoryForbidden{}
}

// WithXRequestID adds the xRequestId to the get repository forbidden response
func (o *GetRepositoryForbidden) WithXRequestID(xRequestID string) *GetRepositoryForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get repository forbidden response
func (o *GetRepositoryForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get repository forbidden response
func (o *GetRepositoryForbidden) WithPayload(payload *models.Errors) *GetRepositoryForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repository forbidden response
func (o *GetRepositoryForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRepositoryForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetRepositoryNotFoundCode is the HTTP code returned for type GetRepositoryNotFound
const GetRepositoryNotFoundCode int = 404

/*GetRepositoryNotFound Not found

swagger:response getRepositoryNotFound
*/
type GetRepositoryNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRepositoryNotFound creates GetRepositoryNotFound with default headers values
func NewGetRepositoryNotFound() *GetRepositoryNotFound {

	return &GetRepositoryNotFound{}
}

// WithXRequestID adds the xRequestId to the get repository not found response
func (o *GetRepositoryNotFound) WithXRequestID(xRequestID string) *GetRepositoryNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get repository not found response
func (o *GetRepositoryNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get repository not found response
func (o *GetRepositoryNotFound) WithPayload(payload *models.Errors) *GetRepositoryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repository not found response
func (o *GetRepositoryNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRepositoryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetRepositoryInternalServerErrorCode is the HTTP code returned for type GetRepositoryInternalServerError
const GetRepositoryInternalServerErrorCode int = 500

/*GetRepositoryInternalServerError Internal server error

swagger:response getRepositoryInternalServerError
*/
type GetRepositoryInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRepositoryInternalServerError creates GetRepositoryInternalServerError with default headers values
func NewGetRepositoryInternalServerError() *GetRepositoryInternalServerError {

	return &GetRepositoryInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get repository internal server error response
func (o *GetRepositoryInternalServerError) WithXRequestID(xRequestID string) *GetRepositoryInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get repository internal server error response
func (o *GetRepositoryInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get repository internal server error response
func (o *GetRepositoryInternalServerError) WithPayload(payload *models.Errors) *GetRepositoryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repository internal server error response
func (o *GetRepositoryInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRepositoryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
