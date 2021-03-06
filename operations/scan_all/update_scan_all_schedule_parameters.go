// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// NewUpdateScanAllScheduleParams creates a new UpdateScanAllScheduleParams object
// no default values defined in spec.
func NewUpdateScanAllScheduleParams() UpdateScanAllScheduleParams {

	return UpdateScanAllScheduleParams{}
}

// UpdateScanAllScheduleParams contains all the bound params for the update scan all schedule operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateScanAllSchedule
type UpdateScanAllScheduleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Updates the schedule of scan all job, which scans all of images in Harbor.
	  Required: true
	  In: body
	*/
	Schedule *models.Schedule
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateScanAllScheduleParams() beforehand.
func (o *UpdateScanAllScheduleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Schedule
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("schedule", "body"))
			} else {
				res = append(res, errors.NewParseError("schedule", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Schedule = &body
			}
		}
	} else {
		res = append(res, errors.Required("schedule", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
