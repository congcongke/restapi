// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// ListArtifactsOKCode is the HTTP code returned for type ListArtifactsOK
const ListArtifactsOKCode int = 200

/*ListArtifactsOK Success

swagger:response listArtifactsOK
*/
type ListArtifactsOK struct {
	/*Link refers to the previous page and next page

	 */
	Link string `json:"Link"`
	/*The total count of artifacts

	 */
	XTotalCount int64 `json:"X-Total-Count"`

	/*
	  In: Body
	*/
	Payload []*models.Artifact `json:"body,omitempty"`
}

// NewListArtifactsOK creates ListArtifactsOK with default headers values
func NewListArtifactsOK() *ListArtifactsOK {

	return &ListArtifactsOK{}
}

// WithLink adds the link to the list artifacts o k response
func (o *ListArtifactsOK) WithLink(link string) *ListArtifactsOK {
	o.Link = link
	return o
}

// SetLink sets the link to the list artifacts o k response
func (o *ListArtifactsOK) SetLink(link string) {
	o.Link = link
}

// WithXTotalCount adds the xTotalCount to the list artifacts o k response
func (o *ListArtifactsOK) WithXTotalCount(xTotalCount int64) *ListArtifactsOK {
	o.XTotalCount = xTotalCount
	return o
}

// SetXTotalCount sets the xTotalCount to the list artifacts o k response
func (o *ListArtifactsOK) SetXTotalCount(xTotalCount int64) {
	o.XTotalCount = xTotalCount
}

// WithPayload adds the payload to the list artifacts o k response
func (o *ListArtifactsOK) WithPayload(payload []*models.Artifact) *ListArtifactsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list artifacts o k response
func (o *ListArtifactsOK) SetPayload(payload []*models.Artifact) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListArtifactsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Link

	link := o.Link
	if link != "" {
		rw.Header().Set("Link", link)
	}

	// response header X-Total-Count

	xTotalCount := swag.FormatInt64(o.XTotalCount)
	if xTotalCount != "" {
		rw.Header().Set("X-Total-Count", xTotalCount)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Artifact, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListArtifactsBadRequestCode is the HTTP code returned for type ListArtifactsBadRequest
const ListArtifactsBadRequestCode int = 400

/*ListArtifactsBadRequest Bad request

swagger:response listArtifactsBadRequest
*/
type ListArtifactsBadRequest struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListArtifactsBadRequest creates ListArtifactsBadRequest with default headers values
func NewListArtifactsBadRequest() *ListArtifactsBadRequest {

	return &ListArtifactsBadRequest{}
}

// WithXRequestID adds the xRequestId to the list artifacts bad request response
func (o *ListArtifactsBadRequest) WithXRequestID(xRequestID string) *ListArtifactsBadRequest {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list artifacts bad request response
func (o *ListArtifactsBadRequest) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list artifacts bad request response
func (o *ListArtifactsBadRequest) WithPayload(payload *models.Errors) *ListArtifactsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list artifacts bad request response
func (o *ListArtifactsBadRequest) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListArtifactsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListArtifactsUnauthorizedCode is the HTTP code returned for type ListArtifactsUnauthorized
const ListArtifactsUnauthorizedCode int = 401

/*ListArtifactsUnauthorized Unauthorized

swagger:response listArtifactsUnauthorized
*/
type ListArtifactsUnauthorized struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListArtifactsUnauthorized creates ListArtifactsUnauthorized with default headers values
func NewListArtifactsUnauthorized() *ListArtifactsUnauthorized {

	return &ListArtifactsUnauthorized{}
}

// WithXRequestID adds the xRequestId to the list artifacts unauthorized response
func (o *ListArtifactsUnauthorized) WithXRequestID(xRequestID string) *ListArtifactsUnauthorized {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list artifacts unauthorized response
func (o *ListArtifactsUnauthorized) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list artifacts unauthorized response
func (o *ListArtifactsUnauthorized) WithPayload(payload *models.Errors) *ListArtifactsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list artifacts unauthorized response
func (o *ListArtifactsUnauthorized) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListArtifactsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListArtifactsForbiddenCode is the HTTP code returned for type ListArtifactsForbidden
const ListArtifactsForbiddenCode int = 403

/*ListArtifactsForbidden Forbidden

swagger:response listArtifactsForbidden
*/
type ListArtifactsForbidden struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListArtifactsForbidden creates ListArtifactsForbidden with default headers values
func NewListArtifactsForbidden() *ListArtifactsForbidden {

	return &ListArtifactsForbidden{}
}

// WithXRequestID adds the xRequestId to the list artifacts forbidden response
func (o *ListArtifactsForbidden) WithXRequestID(xRequestID string) *ListArtifactsForbidden {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list artifacts forbidden response
func (o *ListArtifactsForbidden) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list artifacts forbidden response
func (o *ListArtifactsForbidden) WithPayload(payload *models.Errors) *ListArtifactsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list artifacts forbidden response
func (o *ListArtifactsForbidden) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListArtifactsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListArtifactsNotFoundCode is the HTTP code returned for type ListArtifactsNotFound
const ListArtifactsNotFoundCode int = 404

/*ListArtifactsNotFound Not found

swagger:response listArtifactsNotFound
*/
type ListArtifactsNotFound struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListArtifactsNotFound creates ListArtifactsNotFound with default headers values
func NewListArtifactsNotFound() *ListArtifactsNotFound {

	return &ListArtifactsNotFound{}
}

// WithXRequestID adds the xRequestId to the list artifacts not found response
func (o *ListArtifactsNotFound) WithXRequestID(xRequestID string) *ListArtifactsNotFound {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list artifacts not found response
func (o *ListArtifactsNotFound) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list artifacts not found response
func (o *ListArtifactsNotFound) WithPayload(payload *models.Errors) *ListArtifactsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list artifacts not found response
func (o *ListArtifactsNotFound) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListArtifactsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// ListArtifactsInternalServerErrorCode is the HTTP code returned for type ListArtifactsInternalServerError
const ListArtifactsInternalServerErrorCode int = 500

/*ListArtifactsInternalServerError Internal server error

swagger:response listArtifactsInternalServerError
*/
type ListArtifactsInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewListArtifactsInternalServerError creates ListArtifactsInternalServerError with default headers values
func NewListArtifactsInternalServerError() *ListArtifactsInternalServerError {

	return &ListArtifactsInternalServerError{}
}

// WithXRequestID adds the xRequestId to the list artifacts internal server error response
func (o *ListArtifactsInternalServerError) WithXRequestID(xRequestID string) *ListArtifactsInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the list artifacts internal server error response
func (o *ListArtifactsInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the list artifacts internal server error response
func (o *ListArtifactsInternalServerError) WithPayload(payload *models.Errors) *ListArtifactsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list artifacts internal server error response
func (o *ListArtifactsInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListArtifactsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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