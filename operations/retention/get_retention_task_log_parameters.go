// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRetentionTaskLogParams creates a new GetRetentionTaskLogParams object
// no default values defined in spec.
func NewGetRetentionTaskLogParams() GetRetentionTaskLogParams {

	return GetRetentionTaskLogParams{}
}

// GetRetentionTaskLogParams contains all the bound params for the get retention task log operation
// typically these are obtained from a http.Request
//
// swagger:parameters getRetentionTaskLog
type GetRetentionTaskLogParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Retention execution ID.
	  Required: true
	  In: path
	*/
	Eid int64
	/*Retention ID.
	  Required: true
	  In: path
	*/
	ID int64
	/*Retention execution ID.
	  Required: true
	  In: path
	*/
	Tid int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRetentionTaskLogParams() beforehand.
func (o *GetRetentionTaskLogParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rEid, rhkEid, _ := route.Params.GetOK("eid")
	if err := o.bindEid(rEid, rhkEid, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	rTid, rhkTid, _ := route.Params.GetOK("tid")
	if err := o.bindTid(rTid, rhkTid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEid binds and validates parameter Eid from path.
func (o *GetRetentionTaskLogParams) bindEid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("eid", "path", "int64", raw)
	}
	o.Eid = value

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetRetentionTaskLogParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("id", "path", "int64", raw)
	}
	o.ID = value

	return nil
}

// bindTid binds and validates parameter Tid from path.
func (o *GetRetentionTaskLogParams) bindTid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("tid", "path", "int64", raw)
	}
	o.Tid = value

	return nil
}
