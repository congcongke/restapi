// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetLatestScanAllMetricsParams creates a new GetLatestScanAllMetricsParams object
// no default values defined in spec.
func NewGetLatestScanAllMetricsParams() GetLatestScanAllMetricsParams {

	return GetLatestScanAllMetricsParams{}
}

// GetLatestScanAllMetricsParams contains all the bound params for the get latest scan all metrics operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLatestScanAllMetrics
type GetLatestScanAllMetricsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLatestScanAllMetricsParams() beforehand.
func (o *GetLatestScanAllMetricsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
