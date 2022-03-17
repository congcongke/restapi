// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListPoliciesHandlerFunc turns a function with the right signature into a list policies handler
type ListPoliciesHandlerFunc func(ListPoliciesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListPoliciesHandlerFunc) Handle(params ListPoliciesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListPoliciesHandler interface for that can handle valid list policies params
type ListPoliciesHandler interface {
	Handle(ListPoliciesParams, interface{}) middleware.Responder
}

// NewListPolicies creates a new http.Handler for the list policies operation
func NewListPolicies(ctx *middleware.Context, handler ListPoliciesHandler) *ListPolicies {
	return &ListPolicies{Context: ctx, Handler: handler}
}

/*ListPolicies swagger:route GET /projects/{project_name}/preheat/policies preheat listPolicies

List preheat policies

List preheat policies

*/
type ListPolicies struct {
	Context *middleware.Context
	Handler ListPoliciesHandler
}

func (o *ListPolicies) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListPoliciesParams()

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
