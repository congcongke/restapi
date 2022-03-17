// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// CreateScanAllScheduleCreatedCode is the HTTP code returned for type CreateScanAllScheduleCreated
const CreateScanAllScheduleCreatedCode int = 201

/*CreateScanAllScheduleCreated Created

swagger:response createScanAllScheduleCreated
*/
type CreateScanAllScheduleCreated struct {
	/*The location of the resource

	 */
	Location string `json:"Location"`
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`
}

// NewCreateScanAllScheduleCreated creates CreateScanAllScheduleCreated with default headers values
func NewCreateScanAllScheduleCreated() *CreateScanAllScheduleCreated {

	return &CreateScanAllScheduleCreated{}
}

// WithLocation adds the location to the create scan all schedule created response
func (o *CreateScanAllScheduleCreated) WithLocation(location string) *CreateScanAllScheduleCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the create scan all schedule created response
func (o *CreateScanAllScheduleCreated) SetLocation(location string) {
	o.Location = location
}

// WithXRequestID adds the xRequestId to the create scan all schedule created response
func (o *CreateScanAllScheduleCreated) WithXRequestID(xRequestID string) *CreateScanAllScheduleCreated {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create scan all schedule created response
func (o *CreateScanAllScheduleCreated) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WriteResponse to the client
func (o *CreateScanAllScheduleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location

	location := o.Location
	if location != "" {
		rw.Header().Set("Location", location)
	}

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// CreateScanAllScheduleBadRequestCode is the HTTP code returned for type CreateScanAllScheduleBadRequest
const CreateScanAllScheduleBadRequestCode int = 400

/*CreateScanAllScheduleBadRequest Bad request

swagger:response createScanAllScheduleBadRequest
*/
type CreateScanAllScheduleBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateScanAllScheduleBadRequest creates CreateScanAllScheduleBadRequest with default headers values
func NewCreateScanAllScheduleBadRequest() *CreateScanAllScheduleBadRequest {

	return &CreateScanAllScheduleBadRequest{}
}

// WithXRequestID adds the xRequestId to the create scan all schedule bad request response
func (o *CreateScanAllScheduleBadRequest) WithXRequestID(xRequestID string) *CreateScanAllScheduleBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create scan all schedule bad request response
func (o *CreateScanAllScheduleBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create scan all schedule bad request response
func (o *CreateScanAllScheduleBadRequest) WithPayload(payload *models.Errors) *CreateScanAllScheduleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create scan all schedule bad request response
func (o *CreateScanAllScheduleBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateScanAllScheduleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateScanAllScheduleUnauthorizedCode is the HTTP code returned for type CreateScanAllScheduleUnauthorized
const CreateScanAllScheduleUnauthorizedCode int = 401

/*CreateScanAllScheduleUnauthorized Unauthorized

swagger:response createScanAllScheduleUnauthorized
*/
type CreateScanAllScheduleUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateScanAllScheduleUnauthorized creates CreateScanAllScheduleUnauthorized with default headers values
func NewCreateScanAllScheduleUnauthorized() *CreateScanAllScheduleUnauthorized {

	return &CreateScanAllScheduleUnauthorized{}
}

// WithXRequestID adds the xRequestId to the create scan all schedule unauthorized response
func (o *CreateScanAllScheduleUnauthorized) WithXRequestID(xRequestID string) *CreateScanAllScheduleUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create scan all schedule unauthorized response
func (o *CreateScanAllScheduleUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create scan all schedule unauthorized response
func (o *CreateScanAllScheduleUnauthorized) WithPayload(payload *models.Errors) *CreateScanAllScheduleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create scan all schedule unauthorized response
func (o *CreateScanAllScheduleUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateScanAllScheduleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateScanAllScheduleForbiddenCode is the HTTP code returned for type CreateScanAllScheduleForbidden
const CreateScanAllScheduleForbiddenCode int = 403

/*CreateScanAllScheduleForbidden Forbidden

swagger:response createScanAllScheduleForbidden
*/
type CreateScanAllScheduleForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateScanAllScheduleForbidden creates CreateScanAllScheduleForbidden with default headers values
func NewCreateScanAllScheduleForbidden() *CreateScanAllScheduleForbidden {

	return &CreateScanAllScheduleForbidden{}
}

// WithXRequestID adds the xRequestId to the create scan all schedule forbidden response
func (o *CreateScanAllScheduleForbidden) WithXRequestID(xRequestID string) *CreateScanAllScheduleForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create scan all schedule forbidden response
func (o *CreateScanAllScheduleForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create scan all schedule forbidden response
func (o *CreateScanAllScheduleForbidden) WithPayload(payload *models.Errors) *CreateScanAllScheduleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create scan all schedule forbidden response
func (o *CreateScanAllScheduleForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateScanAllScheduleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateScanAllScheduleConflictCode is the HTTP code returned for type CreateScanAllScheduleConflict
const CreateScanAllScheduleConflictCode int = 409

/*CreateScanAllScheduleConflict Conflict

swagger:response createScanAllScheduleConflict
*/
type CreateScanAllScheduleConflict struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateScanAllScheduleConflict creates CreateScanAllScheduleConflict with default headers values
func NewCreateScanAllScheduleConflict() *CreateScanAllScheduleConflict {

	return &CreateScanAllScheduleConflict{}
}

// WithXRequestID adds the xRequestId to the create scan all schedule conflict response
func (o *CreateScanAllScheduleConflict) WithXRequestID(xRequestID string) *CreateScanAllScheduleConflict {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create scan all schedule conflict response
func (o *CreateScanAllScheduleConflict) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create scan all schedule conflict response
func (o *CreateScanAllScheduleConflict) WithPayload(payload *models.Errors) *CreateScanAllScheduleConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create scan all schedule conflict response
func (o *CreateScanAllScheduleConflict) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateScanAllScheduleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateScanAllSchedulePreconditionFailedCode is the HTTP code returned for type CreateScanAllSchedulePreconditionFailed
const CreateScanAllSchedulePreconditionFailedCode int = 412

/*CreateScanAllSchedulePreconditionFailed Precondition failed

swagger:response createScanAllSchedulePreconditionFailed
*/
type CreateScanAllSchedulePreconditionFailed struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateScanAllSchedulePreconditionFailed creates CreateScanAllSchedulePreconditionFailed with default headers values
func NewCreateScanAllSchedulePreconditionFailed() *CreateScanAllSchedulePreconditionFailed {

	return &CreateScanAllSchedulePreconditionFailed{}
}

// WithXRequestID adds the xRequestId to the create scan all schedule precondition failed response
func (o *CreateScanAllSchedulePreconditionFailed) WithXRequestID(xRequestID string) *CreateScanAllSchedulePreconditionFailed {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create scan all schedule precondition failed response
func (o *CreateScanAllSchedulePreconditionFailed) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create scan all schedule precondition failed response
func (o *CreateScanAllSchedulePreconditionFailed) WithPayload(payload *models.Errors) *CreateScanAllSchedulePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create scan all schedule precondition failed response
func (o *CreateScanAllSchedulePreconditionFailed) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateScanAllSchedulePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateScanAllScheduleInternalServerErrorCode is the HTTP code returned for type CreateScanAllScheduleInternalServerError
const CreateScanAllScheduleInternalServerErrorCode int = 500

/*CreateScanAllScheduleInternalServerError Internal server error

swagger:response createScanAllScheduleInternalServerError
*/
type CreateScanAllScheduleInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewCreateScanAllScheduleInternalServerError creates CreateScanAllScheduleInternalServerError with default headers values
func NewCreateScanAllScheduleInternalServerError() *CreateScanAllScheduleInternalServerError {

	return &CreateScanAllScheduleInternalServerError{}
}

// WithXRequestID adds the xRequestId to the create scan all schedule internal server error response
func (o *CreateScanAllScheduleInternalServerError) WithXRequestID(xRequestID string) *CreateScanAllScheduleInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the create scan all schedule internal server error response
func (o *CreateScanAllScheduleInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the create scan all schedule internal server error response
func (o *CreateScanAllScheduleInternalServerError) WithPayload(payload *models.Errors) *CreateScanAllScheduleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create scan all schedule internal server error response
func (o *CreateScanAllScheduleInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateScanAllScheduleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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