// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/goharbor/harbor/src/server/v2.0/models"
)

// NewRefreshSecParams creates a new RefreshSecParams object
// no default values defined in spec.
func NewRefreshSecParams() RefreshSecParams {

	return RefreshSecParams{}
}

// RefreshSecParams contains all the bound params for the refresh sec operation
// typically these are obtained from a http.Request
//
// swagger:parameters RefreshSec
type RefreshSecParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*An unique ID for the request
	  Min Length: 1
	  In: header
	*/
	XRequestID *string
	/*The JSON object of a robot account.
	  Required: true
	  In: body
	*/
	RobotSec *models.RobotSec
	/*Robot ID
	  Required: true
	  In: path
	*/
	RobotID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRefreshSecParams() beforehand.
func (o *RefreshSecParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXRequestID(r.Header[http.CanonicalHeaderKey("X-Request-Id")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.RobotSec
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("robotSec", "body"))
			} else {
				res = append(res, errors.NewParseError("robotSec", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.RobotSec = &body
			}
		}
	} else {
		res = append(res, errors.Required("robotSec", "body"))
	}
	rRobotID, rhkRobotID, _ := route.Params.GetOK("robot_id")
	if err := o.bindRobotID(rRobotID, rhkRobotID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRequestID binds and validates parameter XRequestID from header.
func (o *RefreshSecParams) bindXRequestID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XRequestID = &raw

	if err := o.validateXRequestID(formats); err != nil {
		return err
	}

	return nil
}

// validateXRequestID carries on validations for parameter XRequestID
func (o *RefreshSecParams) validateXRequestID(formats strfmt.Registry) error {

	if err := validate.MinLength("X-Request-Id", "header", (*o.XRequestID), 1); err != nil {
		return err
	}

	return nil
}

// bindRobotID binds and validates parameter RobotID from path.
func (o *RefreshSecParams) bindRobotID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("robot_id", "path", "int64", raw)
	}
	o.RobotID = value

	return nil
}