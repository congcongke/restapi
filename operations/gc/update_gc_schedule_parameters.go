// Code generated by go-swagger; DO NOT EDIT.

package gc

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

// NewUpdateGCScheduleParams creates a new UpdateGCScheduleParams object
// no default values defined in spec.
func NewUpdateGCScheduleParams() UpdateGCScheduleParams {

	return UpdateGCScheduleParams{}
}

// UpdateGCScheduleParams contains all the bound params for the update GC schedule operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateGCSchedule
type UpdateGCScheduleParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Updates of gc's schedule.
	  Required: true
	  In: body
	*/
	Schedule *models.Schedule
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateGCScheduleParams() beforehand.
func (o *UpdateGCScheduleParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
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