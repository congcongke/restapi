// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetScanAllScheduleOKCode is the HTTP code returned for type GetScanAllScheduleOK
const GetScanAllScheduleOKCode int = 200

/*GetScanAllScheduleOK Get a schedule for the scan all job, which scans all of images in Harbor.

swagger:response getScanAllScheduleOK
*/
type GetScanAllScheduleOK struct {

	/*
	  In: Body
	*/
	Payload *models.Schedule `json:"body,omitempty"`
}

// NewGetScanAllScheduleOK creates GetScanAllScheduleOK with default headers values
func NewGetScanAllScheduleOK() *GetScanAllScheduleOK {

	return &GetScanAllScheduleOK{}
}

// WithPayload adds the payload to the get scan all schedule o k response
func (o *GetScanAllScheduleOK) WithPayload(payload *models.Schedule) *GetScanAllScheduleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get scan all schedule o k response
func (o *GetScanAllScheduleOK) SetPayload(payload *models.Schedule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScanAllScheduleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetScanAllScheduleUnauthorizedCode is the HTTP code returned for type GetScanAllScheduleUnauthorized
const GetScanAllScheduleUnauthorizedCode int = 401

/*GetScanAllScheduleUnauthorized Unauthorized

swagger:response getScanAllScheduleUnauthorized
*/
type GetScanAllScheduleUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetScanAllScheduleUnauthorized creates GetScanAllScheduleUnauthorized with default headers values
func NewGetScanAllScheduleUnauthorized() *GetScanAllScheduleUnauthorized {

	return &GetScanAllScheduleUnauthorized{}
}

// WithXRequestID adds the xRequestId to the get scan all schedule unauthorized response
func (o *GetScanAllScheduleUnauthorized) WithXRequestID(xRequestID string) *GetScanAllScheduleUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get scan all schedule unauthorized response
func (o *GetScanAllScheduleUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get scan all schedule unauthorized response
func (o *GetScanAllScheduleUnauthorized) WithPayload(payload *models.Errors) *GetScanAllScheduleUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get scan all schedule unauthorized response
func (o *GetScanAllScheduleUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScanAllScheduleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetScanAllScheduleForbiddenCode is the HTTP code returned for type GetScanAllScheduleForbidden
const GetScanAllScheduleForbiddenCode int = 403

/*GetScanAllScheduleForbidden Forbidden

swagger:response getScanAllScheduleForbidden
*/
type GetScanAllScheduleForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetScanAllScheduleForbidden creates GetScanAllScheduleForbidden with default headers values
func NewGetScanAllScheduleForbidden() *GetScanAllScheduleForbidden {

	return &GetScanAllScheduleForbidden{}
}

// WithXRequestID adds the xRequestId to the get scan all schedule forbidden response
func (o *GetScanAllScheduleForbidden) WithXRequestID(xRequestID string) *GetScanAllScheduleForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get scan all schedule forbidden response
func (o *GetScanAllScheduleForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get scan all schedule forbidden response
func (o *GetScanAllScheduleForbidden) WithPayload(payload *models.Errors) *GetScanAllScheduleForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get scan all schedule forbidden response
func (o *GetScanAllScheduleForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScanAllScheduleForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetScanAllSchedulePreconditionFailedCode is the HTTP code returned for type GetScanAllSchedulePreconditionFailed
const GetScanAllSchedulePreconditionFailedCode int = 412

/*GetScanAllSchedulePreconditionFailed Precondition failed

swagger:response getScanAllSchedulePreconditionFailed
*/
type GetScanAllSchedulePreconditionFailed struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetScanAllSchedulePreconditionFailed creates GetScanAllSchedulePreconditionFailed with default headers values
func NewGetScanAllSchedulePreconditionFailed() *GetScanAllSchedulePreconditionFailed {

	return &GetScanAllSchedulePreconditionFailed{}
}

// WithXRequestID adds the xRequestId to the get scan all schedule precondition failed response
func (o *GetScanAllSchedulePreconditionFailed) WithXRequestID(xRequestID string) *GetScanAllSchedulePreconditionFailed {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get scan all schedule precondition failed response
func (o *GetScanAllSchedulePreconditionFailed) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get scan all schedule precondition failed response
func (o *GetScanAllSchedulePreconditionFailed) WithPayload(payload *models.Errors) *GetScanAllSchedulePreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get scan all schedule precondition failed response
func (o *GetScanAllSchedulePreconditionFailed) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScanAllSchedulePreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetScanAllScheduleInternalServerErrorCode is the HTTP code returned for type GetScanAllScheduleInternalServerError
const GetScanAllScheduleInternalServerErrorCode int = 500

/*GetScanAllScheduleInternalServerError Internal server error

swagger:response getScanAllScheduleInternalServerError
*/
type GetScanAllScheduleInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetScanAllScheduleInternalServerError creates GetScanAllScheduleInternalServerError with default headers values
func NewGetScanAllScheduleInternalServerError() *GetScanAllScheduleInternalServerError {

	return &GetScanAllScheduleInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get scan all schedule internal server error response
func (o *GetScanAllScheduleInternalServerError) WithXRequestID(xRequestID string) *GetScanAllScheduleInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get scan all schedule internal server error response
func (o *GetScanAllScheduleInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get scan all schedule internal server error response
func (o *GetScanAllScheduleInternalServerError) WithPayload(payload *models.Errors) *GetScanAllScheduleInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get scan all schedule internal server error response
func (o *GetScanAllScheduleInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScanAllScheduleInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
