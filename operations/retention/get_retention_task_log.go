// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRetentionTaskLogHandlerFunc turns a function with the right signature into a get retention task log handler
type GetRetentionTaskLogHandlerFunc func(GetRetentionTaskLogParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRetentionTaskLogHandlerFunc) Handle(params GetRetentionTaskLogParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetRetentionTaskLogHandler interface for that can handle valid get retention task log params
type GetRetentionTaskLogHandler interface {
	Handle(GetRetentionTaskLogParams, interface{}) middleware.Responder
}

// NewGetRetentionTaskLog creates a new http.Handler for the get retention task log operation
func NewGetRetentionTaskLog(ctx *middleware.Context, handler GetRetentionTaskLogHandler) *GetRetentionTaskLog {
	return &GetRetentionTaskLog{Context: ctx, Handler: handler}
}

/*GetRetentionTaskLog swagger:route GET /retentions/{id}/executions/{eid}/tasks/{tid} Retention getRetentionTaskLog

Get Retention job task log

Get Retention job task log, tags ratain or deletion detail will be shown in a table.

*/
type GetRetentionTaskLog struct {
	Context *middleware.Context
	Handler GetRetentionTaskLogHandler
}

func (o *GetRetentionTaskLog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRetentionTaskLogParams()

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
