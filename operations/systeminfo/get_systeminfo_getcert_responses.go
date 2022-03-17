// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// GetSysteminfoGetcertOKCode is the HTTP code returned for type GetSysteminfoGetcertOK
const GetSysteminfoGetcertOKCode int = 200

/*GetSysteminfoGetcertOK Get default root certificate successfully.

swagger:response getSysteminfoGetcertOK
*/
type GetSysteminfoGetcertOK struct {
	/*To set the filename of the downloaded file.

	 */
	ContentDisposition string `json:"Content-Disposition"`

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetSysteminfoGetcertOK creates GetSysteminfoGetcertOK with default headers values
func NewGetSysteminfoGetcertOK() *GetSysteminfoGetcertOK {

	return &GetSysteminfoGetcertOK{}
}

// WithContentDisposition adds the contentDisposition to the get systeminfo getcert o k response
func (o *GetSysteminfoGetcertOK) WithContentDisposition(contentDisposition string) *GetSysteminfoGetcertOK {
	o.ContentDisposition = contentDisposition
	return o
}

// SetContentDisposition sets the contentDisposition to the get systeminfo getcert o k response
func (o *GetSysteminfoGetcertOK) SetContentDisposition(contentDisposition string) {
	o.ContentDisposition = contentDisposition
}

// WithPayload adds the payload to the get systeminfo getcert o k response
func (o *GetSysteminfoGetcertOK) WithPayload(payload io.ReadCloser) *GetSysteminfoGetcertOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get systeminfo getcert o k response
func (o *GetSysteminfoGetcertOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSysteminfoGetcertOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Disposition

	contentDisposition := o.ContentDisposition
	if contentDisposition != "" {
		rw.Header().Set("Content-Disposition", contentDisposition)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetSysteminfoGetcertNotFoundCode is the HTTP code returned for type GetSysteminfoGetcertNotFound
const GetSysteminfoGetcertNotFoundCode int = 404

/*GetSysteminfoGetcertNotFound Not found the default root certificate.

swagger:response getSysteminfoGetcertNotFound
*/
type GetSysteminfoGetcertNotFound struct {
}

// NewGetSysteminfoGetcertNotFound creates GetSysteminfoGetcertNotFound with default headers values
func NewGetSysteminfoGetcertNotFound() *GetSysteminfoGetcertNotFound {

	return &GetSysteminfoGetcertNotFound{}
}

// WriteResponse to the client
func (o *GetSysteminfoGetcertNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) // Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetSysteminfoGetcertInternalServerErrorCode is the HTTP code returned for type GetSysteminfoGetcertInternalServerError
const GetSysteminfoGetcertInternalServerErrorCode int = 500

/*GetSysteminfoGetcertInternalServerError Internal server error

swagger:response getSysteminfoGetcertInternalServerError
*/
type GetSysteminfoGetcertInternalServerError struct {
	/*The ID of the corresponding request for the response

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.Errors `json:"body,omitempty"`
}

// NewGetSysteminfoGetcertInternalServerError creates GetSysteminfoGetcertInternalServerError with default headers values
func NewGetSysteminfoGetcertInternalServerError() *GetSysteminfoGetcertInternalServerError {

	return &GetSysteminfoGetcertInternalServerError{}
}

// WithXRequestID adds the xRequestId to the get systeminfo getcert internal server error response
func (o *GetSysteminfoGetcertInternalServerError) WithXRequestID(xRequestID string) *GetSysteminfoGetcertInternalServerError {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the get systeminfo getcert internal server error response
func (o *GetSysteminfoGetcertInternalServerError) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the get systeminfo getcert internal server error response
func (o *GetSysteminfoGetcertInternalServerError) WithPayload(payload *models.Errors) *GetSysteminfoGetcertInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get systeminfo getcert internal server error response
func (o *GetSysteminfoGetcertInternalServerError) SetPayload(payload *models.Errors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSysteminfoGetcertInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
