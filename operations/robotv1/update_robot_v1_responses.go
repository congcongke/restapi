// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// UpdateRobotV1OKCode is the HTTP code returned for type UpdateRobotV1OK
const UpdateRobotV1OKCode int = 200

/*UpdateRobotV1OK Success

swagger:response updateRobotV1OK
*/
type UpdateRobotV1OK struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewUpdateRobotV1OK creates UpdateRobotV1OK with default headers values
func NewUpdateRobotV1OK() *UpdateRobotV1OK {

	return &UpdateRobotV1OK{}
}

// WithXRequestID adds the xRequestId to the update robot v1 o k response
func (o *UpdateRobotV1OK) WithXRequestID(xRequestID string) *UpdateRobotV1OK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update robot v1 o k response
func (o *UpdateRobotV1OK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *UpdateRobotV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateRobotV1BadRequestCode is the HTTP code returned for type UpdateRobotV1BadRequest
const UpdateRobotV1BadRequestCode int = 400

/*UpdateRobotV1BadRequest Bad request

swagger:response updateRobotV1BadRequest
*/
type UpdateRobotV1BadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateRobotV1BadRequest creates UpdateRobotV1BadRequest with default headers values
func NewUpdateRobotV1BadRequest() *UpdateRobotV1BadRequest {

	return &UpdateRobotV1BadRequest{}
}

// WithXRequestID adds the xRequestId to the update robot v1 bad request response
func (o *UpdateRobotV1BadRequest) WithXRequestID(xRequestID string) *UpdateRobotV1BadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update robot v1 bad request response
func (o *UpdateRobotV1BadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update robot v1 bad request response
func (o *UpdateRobotV1BadRequest) WithPayload(payload *models.Errors) *UpdateRobotV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update robot v1 bad request response
func (o *UpdateRobotV1BadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRobotV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateRobotV1UnauthorizedCode is the HTTP code returned for type UpdateRobotV1Unauthorized
const UpdateRobotV1UnauthorizedCode int = 401

/*UpdateRobotV1Unauthorized Unauthorized

swagger:response updateRobotV1Unauthorized
*/
type UpdateRobotV1Unauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateRobotV1Unauthorized creates UpdateRobotV1Unauthorized with default headers values
func NewUpdateRobotV1Unauthorized() *UpdateRobotV1Unauthorized {

	return &UpdateRobotV1Unauthorized{}
}

// WithXRequestID adds the xRequestId to the update robot v1 unauthorized response
func (o *UpdateRobotV1Unauthorized) WithXRequestID(xRequestID string) *UpdateRobotV1Unauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update robot v1 unauthorized response
func (o *UpdateRobotV1Unauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update robot v1 unauthorized response
func (o *UpdateRobotV1Unauthorized) WithPayload(payload *models.Errors) *UpdateRobotV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update robot v1 unauthorized response
func (o *UpdateRobotV1Unauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRobotV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateRobotV1ForbiddenCode is the HTTP code returned for type UpdateRobotV1Forbidden
const UpdateRobotV1ForbiddenCode int = 403

/*UpdateRobotV1Forbidden Forbidden

swagger:response updateRobotV1Forbidden
*/
type UpdateRobotV1Forbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateRobotV1Forbidden creates UpdateRobotV1Forbidden with default headers values
func NewUpdateRobotV1Forbidden() *UpdateRobotV1Forbidden {

	return &UpdateRobotV1Forbidden{}
}

// WithXRequestID adds the xRequestId to the update robot v1 forbidden response
func (o *UpdateRobotV1Forbidden) WithXRequestID(xRequestID string) *UpdateRobotV1Forbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update robot v1 forbidden response
func (o *UpdateRobotV1Forbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update robot v1 forbidden response
func (o *UpdateRobotV1Forbidden) WithPayload(payload *models.Errors) *UpdateRobotV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update robot v1 forbidden response
func (o *UpdateRobotV1Forbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRobotV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateRobotV1NotFoundCode is the HTTP code returned for type UpdateRobotV1NotFound
const UpdateRobotV1NotFoundCode int = 404

/*UpdateRobotV1NotFound Not found

swagger:response updateRobotV1NotFound
*/
type UpdateRobotV1NotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateRobotV1NotFound creates UpdateRobotV1NotFound with default headers values
func NewUpdateRobotV1NotFound() *UpdateRobotV1NotFound {

	return &UpdateRobotV1NotFound{}
}

// WithXRequestID adds the xRequestId to the update robot v1 not found response
func (o *UpdateRobotV1NotFound) WithXRequestID(xRequestID string) *UpdateRobotV1NotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update robot v1 not found response
func (o *UpdateRobotV1NotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update robot v1 not found response
func (o *UpdateRobotV1NotFound) WithPayload(payload *models.Errors) *UpdateRobotV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update robot v1 not found response
func (o *UpdateRobotV1NotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRobotV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// UpdateRobotV1ConflictCode is the HTTP code returned for type UpdateRobotV1Conflict
const UpdateRobotV1ConflictCode int = 409

/*UpdateRobotV1Conflict Conflict

swagger:response updateRobotV1Conflict
*/
type UpdateRobotV1Conflict struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateRobotV1Conflict creates UpdateRobotV1Conflict with default headers values
func NewUpdateRobotV1Conflict() *UpdateRobotV1Conflict {

	return &UpdateRobotV1Conflict{}
}

// WithXRequestID adds the xRequestId to the update robot v1 conflict response
func (o *UpdateRobotV1Conflict) WithXRequestID(xRequestID string) *UpdateRobotV1Conflict {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update robot v1 conflict response
func (o *UpdateRobotV1Conflict) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update robot v1 conflict response
func (o *UpdateRobotV1Conflict) WithPayload(payload *models.Errors) *UpdateRobotV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update robot v1 conflict response
func (o *UpdateRobotV1Conflict) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRobotV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateRobotV1InternalServerErrorCode is the HTTP code returned for type UpdateRobotV1InternalServerError
const UpdateRobotV1InternalServerErrorCode int = 500

/*UpdateRobotV1InternalServerError Internal server error

swagger:response updateRobotV1InternalServerError
*/
type UpdateRobotV1InternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewUpdateRobotV1InternalServerError creates UpdateRobotV1InternalServerError with default headers values
func NewUpdateRobotV1InternalServerError() *UpdateRobotV1InternalServerError {

	return &UpdateRobotV1InternalServerError{}
}

// WithXRequestID adds the xRequestId to the update robot v1 internal server error response
func (o *UpdateRobotV1InternalServerError) WithXRequestID(xRequestID string) *UpdateRobotV1InternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the update robot v1 internal server error response
func (o *UpdateRobotV1InternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the update robot v1 internal server error response
func (o *UpdateRobotV1InternalServerError) WithPayload(payload *models.Errors) *UpdateRobotV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update robot v1 internal server error response
func (o *UpdateRobotV1InternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateRobotV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
