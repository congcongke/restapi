// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetRobotByIDV1OKCode is the HTTP code returned for type GetRobotByIDV1OK
const GetRobotByIDV1OKCode int = 200

/*GetRobotByIDV1OK Return matched robot information.

swagger:response getRobotByIdV1OK
*/
type GetRobotByIDV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.Robot `json:"body,omitempty"`
}

// NewGetRobotByIDV1OK creates GetRobotByIDV1OK with default headers values
func NewGetRobotByIDV1OK() *GetRobotByIDV1OK {

	return &GetRobotByIDV1OK{}
}

// WithPayload adds the payload to the get robot by Id v1 o k response
func (o *GetRobotByIDV1OK) WithPayload(payload *models.Robot) *GetRobotByIDV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get robot by Id v1 o k response
func (o *GetRobotByIDV1OK) SetPayload(payload *models.Robot) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRobotByIDV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRobotByIDV1UnauthorizedCode is the HTTP code returned for type GetRobotByIDV1Unauthorized
const GetRobotByIDV1UnauthorizedCode int = 401

/*GetRobotByIDV1Unauthorized Unauthorized

swagger:response getRobotByIdV1Unauthorized
*/
type GetRobotByIDV1Unauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRobotByIDV1Unauthorized creates GetRobotByIDV1Unauthorized with default headers values
func NewGetRobotByIDV1Unauthorized() *GetRobotByIDV1Unauthorized {

	return &GetRobotByIDV1Unauthorized{}
}

// WithXRequestID adds the xRequestId to the get robot by Id v1 unauthorized response
func (o *GetRobotByIDV1Unauthorized) WithXRequestID(xRequestID string) *GetRobotByIDV1Unauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get robot by Id v1 unauthorized response
func (o *GetRobotByIDV1Unauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get robot by Id v1 unauthorized response
func (o *GetRobotByIDV1Unauthorized) WithPayload(payload *models.Errors) *GetRobotByIDV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get robot by Id v1 unauthorized response
func (o *GetRobotByIDV1Unauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRobotByIDV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetRobotByIDV1ForbiddenCode is the HTTP code returned for type GetRobotByIDV1Forbidden
const GetRobotByIDV1ForbiddenCode int = 403

/*GetRobotByIDV1Forbidden Forbidden

swagger:response getRobotByIdV1Forbidden
*/
type GetRobotByIDV1Forbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRobotByIDV1Forbidden creates GetRobotByIDV1Forbidden with default headers values
func NewGetRobotByIDV1Forbidden() *GetRobotByIDV1Forbidden {

	return &GetRobotByIDV1Forbidden{}
}

// WithXRequestID adds the xRequestId to the get robot by Id v1 forbidden response
func (o *GetRobotByIDV1Forbidden) WithXRequestID(xRequestID string) *GetRobotByIDV1Forbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get robot by Id v1 forbidden response
func (o *GetRobotByIDV1Forbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get robot by Id v1 forbidden response
func (o *GetRobotByIDV1Forbidden) WithPayload(payload *models.Errors) *GetRobotByIDV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get robot by Id v1 forbidden response
func (o *GetRobotByIDV1Forbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRobotByIDV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetRobotByIDV1NotFoundCode is the HTTP code returned for type GetRobotByIDV1NotFound
const GetRobotByIDV1NotFoundCode int = 404

/*GetRobotByIDV1NotFound Not found

swagger:response getRobotByIdV1NotFound
*/
type GetRobotByIDV1NotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRobotByIDV1NotFound creates GetRobotByIDV1NotFound with default headers values
func NewGetRobotByIDV1NotFound() *GetRobotByIDV1NotFound {

	return &GetRobotByIDV1NotFound{}
}

// WithXRequestID adds the xRequestId to the get robot by Id v1 not found response
func (o *GetRobotByIDV1NotFound) WithXRequestID(xRequestID string) *GetRobotByIDV1NotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get robot by Id v1 not found response
func (o *GetRobotByIDV1NotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get robot by Id v1 not found response
func (o *GetRobotByIDV1NotFound) WithPayload(payload *models.Errors) *GetRobotByIDV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get robot by Id v1 not found response
func (o *GetRobotByIDV1NotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRobotByIDV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetRobotByIDV1InternalServerErrorCode is the HTTP code returned for type GetRobotByIDV1InternalServerError
const GetRobotByIDV1InternalServerErrorCode int = 500

/*GetRobotByIDV1InternalServerError Internal server error

swagger:response getRobotByIdV1InternalServerError
*/
type GetRobotByIDV1InternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetRobotByIDV1InternalServerError creates GetRobotByIDV1InternalServerError with default headers values
func NewGetRobotByIDV1InternalServerError() *GetRobotByIDV1InternalServerError {

	return &GetRobotByIDV1InternalServerError{}
}

// WithXRequestID adds the xRequestId to the get robot by Id v1 internal server error response
func (o *GetRobotByIDV1InternalServerError) WithXRequestID(xRequestID string) *GetRobotByIDV1InternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get robot by Id v1 internal server error response
func (o *GetRobotByIDV1InternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get robot by Id v1 internal server error response
func (o *GetRobotByIDV1InternalServerError) WithPayload(payload *models.Errors) *GetRobotByIDV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get robot by Id v1 internal server error response
func (o *GetRobotByIDV1InternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRobotByIDV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
