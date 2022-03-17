// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetProjectSummaryHandlerFunc turns a function with the right signature into a get project summary handler
type GetProjectSummaryHandlerFunc func(GetProjectSummaryParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectSummaryHandlerFunc) Handle(params GetProjectSummaryParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetProjectSummaryHandler interface for that can handle valid get project summary params
type GetProjectSummaryHandler interface {
	Handle(GetProjectSummaryParams, interface{}) middleware.Responder
}

// NewGetProjectSummary creates a new http.Handler for the get project summary operation
func NewGetProjectSummary(ctx *middleware.Context, handler GetProjectSummaryHandler) *GetProjectSummary {
	return &GetProjectSummary{Context: ctx, Handler: handler}
}

/*GetProjectSummary swagger:route GET /projects/{project_name_or_id}/summary project getProjectSummary

Get summary of the project.

Get summary of the project.

*/
type GetProjectSummary struct {
	Context *middleware.Context
	Handler GetProjectSummaryHandler
}

func (o *GetProjectSummary) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetProjectSummaryParams()

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
