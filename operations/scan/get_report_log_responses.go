// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetReportLogOKCode is the HTTP code returned for type GetReportLogOK
const GetReportLogOKCode int = 200

/*GetReportLogOK Successfully get scan log file

swagger:response getReportLogOK
*/
type GetReportLogOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetReportLogOK creates GetReportLogOK with default headers values
func NewGetReportLogOK() *GetReportLogOK {

	return &GetReportLogOK{}
}

// WithPayload adds the payload to the get report log o k response
func (o *GetReportLogOK) WithPayload(payload string) *GetReportLogOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report log o k response
func (o *GetReportLogOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportLogOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetReportLogUnauthorizedCode is the HTTP code returned for type GetReportLogUnauthorized
const GetReportLogUnauthorizedCode int = 401

/*GetReportLogUnauthorized Unauthorized

swagger:response getReportLogUnauthorized
*/
type GetReportLogUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetReportLogUnauthorized creates GetReportLogUnauthorized with default headers values
func NewGetReportLogUnauthorized() *GetReportLogUnauthorized {

	return &GetReportLogUnauthorized{}
}

// WithXRequestID adds the xRequestId to the get report log unauthorized response
func (o *GetReportLogUnauthorized) WithXRequestID(xRequestID string) *GetReportLogUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get report log unauthorized response
func (o *GetReportLogUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get report log unauthorized response
func (o *GetReportLogUnauthorized) WithPayload(payload *models.Errors) *GetReportLogUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report log unauthorized response
func (o *GetReportLogUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportLogUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetReportLogForbiddenCode is the HTTP code returned for type GetReportLogForbidden
const GetReportLogForbiddenCode int = 403

/*GetReportLogForbidden Forbidden

swagger:response getReportLogForbidden
*/
type GetReportLogForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetReportLogForbidden creates GetReportLogForbidden with default headers values
func NewGetReportLogForbidden() *GetReportLogForbidden {

	return &GetReportLogForbidden{}
}

// WithXRequestID adds the xRequestId to the get report log forbidden response
func (o *GetReportLogForbidden) WithXRequestID(xRequestID string) *GetReportLogForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get report log forbidden response
func (o *GetReportLogForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get report log forbidden response
func (o *GetReportLogForbidden) WithPayload(payload *models.Errors) *GetReportLogForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report log forbidden response
func (o *GetReportLogForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportLogForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetReportLogNotFoundCode is the HTTP code returned for type GetReportLogNotFound
const GetReportLogNotFoundCode int = 404

/*GetReportLogNotFound Not found

swagger:response getReportLogNotFound
*/
type GetReportLogNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetReportLogNotFound creates GetReportLogNotFound with default headers values
func NewGetReportLogNotFound() *GetReportLogNotFound {

	return &GetReportLogNotFound{}
}

// WithXRequestID adds the xRequestId to the get report log not found response
func (o *GetReportLogNotFound) WithXRequestID(xRequestID string) *GetReportLogNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get report log not found response
func (o *GetReportLogNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get report log not found response
func (o *GetReportLogNotFound) WithPayload(payload *models.Errors) *GetReportLogNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report log not found response
func (o *GetReportLogNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportLogNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetReportLogInternalServerErrorCode is the HTTP code returned for type GetReportLogInternalServerError
const GetReportLogInternalServerErrorCode int = 500

/*GetReportLogInternalServerError Internal server error

swagger:response getReportLogInternalServerError
*/
type GetReportLogInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetReportLogInternalServerError creates GetReportLogInternalServerError with default headers values
func NewGetReportLogInternalServerError() *GetReportLogInternalServerError {

	return &GetReportLogInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get report log internal server error response
func (o *GetReportLogInternalServerError) WithXRequestID(xRequestID string) *GetReportLogInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get report log internal server error response
func (o *GetReportLogInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get report log internal server error response
func (o *GetReportLogInternalServerError) WithPayload(payload *models.Errors) *GetReportLogInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report log internal server error response
func (o *GetReportLogInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportLogInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
