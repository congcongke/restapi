// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetLatestScanAllMetricsHandlerFunc turns a function with the right signature into a get latest scan all metrics handler
type GetLatestScanAllMetricsHandlerFunc func(GetLatestScanAllMetricsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLatestScanAllMetricsHandlerFunc) Handle(params GetLatestScanAllMetricsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetLatestScanAllMetricsHandler interface for that can handle valid get latest scan all metrics params
type GetLatestScanAllMetricsHandler interface {
	Handle(GetLatestScanAllMetricsParams, interface{}) middleware.Responder
}

// NewGetLatestScanAllMetrics creates a new http.Handler for the get latest scan all metrics operation
func NewGetLatestScanAllMetrics(ctx *middleware.Context, handler GetLatestScanAllMetricsHandler) *GetLatestScanAllMetrics {
	return &GetLatestScanAllMetrics{Context: ctx, Handler: handler}
}

/*GetLatestScanAllMetrics swagger:route GET /scans/all/metrics scanAll getLatestScanAllMetrics

Get the metrics of the latest scan all process

Get the metrics of the latest scan all process

*/
type GetLatestScanAllMetrics struct {
	Context *middleware.Context
	Handler GetLatestScanAllMetricsHandler
}

func (o *GetLatestScanAllMetrics) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLatestScanAllMetricsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}