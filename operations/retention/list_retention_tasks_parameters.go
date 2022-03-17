// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListRetentionTasksParams creates a new ListRetentionTasksParams object
// no default values defined in spec.
func NewListRetentionTasksParams() ListRetentionTasksParams {

	return ListRetentionTasksParams{}
}

// ListRetentionTasksParams contains all the bound params for the list retention tasks operation
// typically these are obtained from a http.Request
//
// swagger:parameters listRetentionTasks
type ListRetentionTasksParams struct {

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
	/*The page number.
	  In: query
	*/
	Page *int64
	/*The size of per page.
	  In: query
	*/
	PageSize *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListRetentionTasksParams() beforehand.
func (o *ListRetentionTasksParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rEid, rhkEid, _ := route.Params.GetOK("eid")
	if err := o.bindEid(rEid, rhkEid, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPage, qhkPage, _ := qs.GetOK("page")
	if err := o.bindPage(qPage, qhkPage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("page_size")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEid binds and validates parameter Eid from path.
func (o *ListRetentionTasksParams) bindEid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *ListRetentionTasksParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindPage binds and validates parameter Page from query.
func (o *ListRetentionTasksParams) bindPage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page", "query", "int64", raw)
	}
	o.Page = &value

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *ListRetentionTasksParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("page_size", "query", "int64", raw)
	}
	o.PageSize = &value

	return nil
}
