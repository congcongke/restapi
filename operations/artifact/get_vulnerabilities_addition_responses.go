// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetVulnerabilitiesAdditionOKCode is the HTTP code returned for type GetVulnerabilitiesAdditionOK
const GetVulnerabilitiesAdditionOKCode int = 200

/*GetVulnerabilitiesAdditionOK Success

swagger:response getVulnerabilitiesAdditionOK
*/
type GetVulnerabilitiesAdditionOK struct {
	/*The content type of the vulnerabilities addition

	 */
	ContentType string `json:"Content-Type"`

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetVulnerabilitiesAdditionOK creates GetVulnerabilitiesAdditionOK with default headers values
func NewGetVulnerabilitiesAdditionOK() *GetVulnerabilitiesAdditionOK {

	return &GetVulnerabilitiesAdditionOK{}
}

// WithContentType adds the contentType to the get vulnerabilities addition o k response
func (o *GetVulnerabilitiesAdditionOK) WithContentType(contentType string) *GetVulnerabilitiesAdditionOK {
	o.ContentType = contentType
	return o
}

// SetContentType sets the contentType to the get vulnerabilities addition o k response
func (o *GetVulnerabilitiesAdditionOK) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithPayload adds the payload to the get vulnerabilities addition o k response
func (o *GetVulnerabilitiesAdditionOK) WithPayload(payload string) *GetVulnerabilitiesAdditionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get vulnerabilities addition o k response
func (o *GetVulnerabilitiesAdditionOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVulnerabilitiesAdditionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Type

	contentType := o.ContentType
	if contentType != "" {
		rw.Header().Set("Content-Type", contentType)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetVulnerabilitiesAdditionBadRequestCode is the HTTP code returned for type GetVulnerabilitiesAdditionBadRequest
const GetVulnerabilitiesAdditionBadRequestCode int = 400

/*GetVulnerabilitiesAdditionBadRequest Bad request

swagger:response getVulnerabilitiesAdditionBadRequest
*/
type GetVulnerabilitiesAdditionBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetVulnerabilitiesAdditionBadRequest creates GetVulnerabilitiesAdditionBadRequest with default headers values
func NewGetVulnerabilitiesAdditionBadRequest() *GetVulnerabilitiesAdditionBadRequest {

	return &GetVulnerabilitiesAdditionBadRequest{}
}

// WithXRequestID adds the xRequestId to the get vulnerabilities addition bad request response
func (o *GetVulnerabilitiesAdditionBadRequest) WithXRequestID(xRequestID string) *GetVulnerabilitiesAdditionBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get vulnerabilities addition bad request response
func (o *GetVulnerabilitiesAdditionBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get vulnerabilities addition bad request response
func (o *GetVulnerabilitiesAdditionBadRequest) WithPayload(payload *models.Errors) *GetVulnerabilitiesAdditionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get vulnerabilities addition bad request response
func (o *GetVulnerabilitiesAdditionBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVulnerabilitiesAdditionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetVulnerabilitiesAdditionUnauthorizedCode is the HTTP code returned for type GetVulnerabilitiesAdditionUnauthorized
const GetVulnerabilitiesAdditionUnauthorizedCode int = 401

/*GetVulnerabilitiesAdditionUnauthorized Unauthorized

swagger:response getVulnerabilitiesAdditionUnauthorized
*/
type GetVulnerabilitiesAdditionUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetVulnerabilitiesAdditionUnauthorized creates GetVulnerabilitiesAdditionUnauthorized with default headers values
func NewGetVulnerabilitiesAdditionUnauthorized() *GetVulnerabilitiesAdditionUnauthorized {

	return &GetVulnerabilitiesAdditionUnauthorized{}
}

// WithXRequestID adds the xRequestId to the get vulnerabilities addition unauthorized response
func (o *GetVulnerabilitiesAdditionUnauthorized) WithXRequestID(xRequestID string) *GetVulnerabilitiesAdditionUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get vulnerabilities addition unauthorized response
func (o *GetVulnerabilitiesAdditionUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get vulnerabilities addition unauthorized response
func (o *GetVulnerabilitiesAdditionUnauthorized) WithPayload(payload *models.Errors) *GetVulnerabilitiesAdditionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get vulnerabilities addition unauthorized response
func (o *GetVulnerabilitiesAdditionUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVulnerabilitiesAdditionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetVulnerabilitiesAdditionForbiddenCode is the HTTP code returned for type GetVulnerabilitiesAdditionForbidden
const GetVulnerabilitiesAdditionForbiddenCode int = 403

/*GetVulnerabilitiesAdditionForbidden Forbidden

swagger:response getVulnerabilitiesAdditionForbidden
*/
type GetVulnerabilitiesAdditionForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetVulnerabilitiesAdditionForbidden creates GetVulnerabilitiesAdditionForbidden with default headers values
func NewGetVulnerabilitiesAdditionForbidden() *GetVulnerabilitiesAdditionForbidden {

	return &GetVulnerabilitiesAdditionForbidden{}
}

// WithXRequestID adds the xRequestId to the get vulnerabilities addition forbidden response
func (o *GetVulnerabilitiesAdditionForbidden) WithXRequestID(xRequestID string) *GetVulnerabilitiesAdditionForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get vulnerabilities addition forbidden response
func (o *GetVulnerabilitiesAdditionForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get vulnerabilities addition forbidden response
func (o *GetVulnerabilitiesAdditionForbidden) WithPayload(payload *models.Errors) *GetVulnerabilitiesAdditionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get vulnerabilities addition forbidden response
func (o *GetVulnerabilitiesAdditionForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVulnerabilitiesAdditionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetVulnerabilitiesAdditionNotFoundCode is the HTTP code returned for type GetVulnerabilitiesAdditionNotFound
const GetVulnerabilitiesAdditionNotFoundCode int = 404

/*GetVulnerabilitiesAdditionNotFound Not found

swagger:response getVulnerabilitiesAdditionNotFound
*/
type GetVulnerabilitiesAdditionNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetVulnerabilitiesAdditionNotFound creates GetVulnerabilitiesAdditionNotFound with default headers values
func NewGetVulnerabilitiesAdditionNotFound() *GetVulnerabilitiesAdditionNotFound {

	return &GetVulnerabilitiesAdditionNotFound{}
}

// WithXRequestID adds the xRequestId to the get vulnerabilities addition not found response
func (o *GetVulnerabilitiesAdditionNotFound) WithXRequestID(xRequestID string) *GetVulnerabilitiesAdditionNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get vulnerabilities addition not found response
func (o *GetVulnerabilitiesAdditionNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get vulnerabilities addition not found response
func (o *GetVulnerabilitiesAdditionNotFound) WithPayload(payload *models.Errors) *GetVulnerabilitiesAdditionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get vulnerabilities addition not found response
func (o *GetVulnerabilitiesAdditionNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVulnerabilitiesAdditionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetVulnerabilitiesAdditionInternalServerErrorCode is the HTTP code returned for type GetVulnerabilitiesAdditionInternalServerError
const GetVulnerabilitiesAdditionInternalServerErrorCode int = 500

/*GetVulnerabilitiesAdditionInternalServerError Internal server error

swagger:response getVulnerabilitiesAdditionInternalServerError
*/
type GetVulnerabilitiesAdditionInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetVulnerabilitiesAdditionInternalServerError creates GetVulnerabilitiesAdditionInternalServerError with default headers values
func NewGetVulnerabilitiesAdditionInternalServerError() *GetVulnerabilitiesAdditionInternalServerError {

	return &GetVulnerabilitiesAdditionInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get vulnerabilities addition internal server error response
func (o *GetVulnerabilitiesAdditionInternalServerError) WithXRequestID(xRequestID string) *GetVulnerabilitiesAdditionInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get vulnerabilities addition internal server error response
func (o *GetVulnerabilitiesAdditionInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get vulnerabilities addition internal server error response
func (o *GetVulnerabilitiesAdditionInternalServerError) WithPayload(payload *models.Errors) *GetVulnerabilitiesAdditionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get vulnerabilities addition internal server error response
func (o *GetVulnerabilitiesAdditionInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVulnerabilitiesAdditionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
