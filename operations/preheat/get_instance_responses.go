// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetInstanceOKCode is the HTTP code returned for type GetInstanceOK
const GetInstanceOKCode int = 200

/*GetInstanceOK Success

swagger:response getInstanceOK
*/
type GetInstanceOK struct {

	/*
	  In: Body
	*/
	Payload *models.Instance `json:"body,omitempty"`
}

// NewGetInstanceOK creates GetInstanceOK with default headers values
func NewGetInstanceOK() *GetInstanceOK {

	return &GetInstanceOK{}
}

// WithPayload adds the payload to the get instance o k response
func (o *GetInstanceOK) WithPayload(payload *models.Instance) *GetInstanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get instance o k response
func (o *GetInstanceOK) SetPayload(payload *models.Instance) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetInstanceBadRequestCode is the HTTP code returned for type GetInstanceBadRequest
const GetInstanceBadRequestCode int = 400

/*GetInstanceBadRequest Bad request

swagger:response getInstanceBadRequest
*/
type GetInstanceBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetInstanceBadRequest creates GetInstanceBadRequest with default headers values
func NewGetInstanceBadRequest() *GetInstanceBadRequest {

	return &GetInstanceBadRequest{}
}

// WithXRequestID adds the xRequestId to the get instance bad request response
func (o *GetInstanceBadRequest) WithXRequestID(xRequestID string) *GetInstanceBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get instance bad request response
func (o *GetInstanceBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get instance bad request response
func (o *GetInstanceBadRequest) WithPayload(payload *models.Errors) *GetInstanceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get instance bad request response
func (o *GetInstanceBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstanceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetInstanceUnauthorizedCode is the HTTP code returned for type GetInstanceUnauthorized
const GetInstanceUnauthorizedCode int = 401

/*GetInstanceUnauthorized Unauthorized

swagger:response getInstanceUnauthorized
*/
type GetInstanceUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetInstanceUnauthorized creates GetInstanceUnauthorized with default headers values
func NewGetInstanceUnauthorized() *GetInstanceUnauthorized {

	return &GetInstanceUnauthorized{}
}

// WithXRequestID adds the xRequestId to the get instance unauthorized response
func (o *GetInstanceUnauthorized) WithXRequestID(xRequestID string) *GetInstanceUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get instance unauthorized response
func (o *GetInstanceUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get instance unauthorized response
func (o *GetInstanceUnauthorized) WithPayload(payload *models.Errors) *GetInstanceUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get instance unauthorized response
func (o *GetInstanceUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstanceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetInstanceForbiddenCode is the HTTP code returned for type GetInstanceForbidden
const GetInstanceForbiddenCode int = 403

/*GetInstanceForbidden Forbidden

swagger:response getInstanceForbidden
*/
type GetInstanceForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetInstanceForbidden creates GetInstanceForbidden with default headers values
func NewGetInstanceForbidden() *GetInstanceForbidden {

	return &GetInstanceForbidden{}
}

// WithXRequestID adds the xRequestId to the get instance forbidden response
func (o *GetInstanceForbidden) WithXRequestID(xRequestID string) *GetInstanceForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get instance forbidden response
func (o *GetInstanceForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get instance forbidden response
func (o *GetInstanceForbidden) WithPayload(payload *models.Errors) *GetInstanceForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get instance forbidden response
func (o *GetInstanceForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstanceForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetInstanceNotFoundCode is the HTTP code returned for type GetInstanceNotFound
const GetInstanceNotFoundCode int = 404

/*GetInstanceNotFound Not found

swagger:response getInstanceNotFound
*/
type GetInstanceNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetInstanceNotFound creates GetInstanceNotFound with default headers values
func NewGetInstanceNotFound() *GetInstanceNotFound {

	return &GetInstanceNotFound{}
}

// WithXRequestID adds the xRequestId to the get instance not found response
func (o *GetInstanceNotFound) WithXRequestID(xRequestID string) *GetInstanceNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get instance not found response
func (o *GetInstanceNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get instance not found response
func (o *GetInstanceNotFound) WithPayload(payload *models.Errors) *GetInstanceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get instance not found response
func (o *GetInstanceNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstanceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetInstanceInternalServerErrorCode is the HTTP code returned for type GetInstanceInternalServerError
const GetInstanceInternalServerErrorCode int = 500

/*GetInstanceInternalServerError Internal server error

swagger:response getInstanceInternalServerError
*/
type GetInstanceInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetInstanceInternalServerError creates GetInstanceInternalServerError with default headers values
func NewGetInstanceInternalServerError() *GetInstanceInternalServerError {

	return &GetInstanceInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get instance internal server error response
func (o *GetInstanceInternalServerError) WithXRequestID(xRequestID string) *GetInstanceInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get instance internal server error response
func (o *GetInstanceInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get instance internal server error response
func (o *GetInstanceInternalServerError) WithPayload(payload *models.Errors) *GetInstanceInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get instance internal server error response
func (o *GetInstanceInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInstanceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
