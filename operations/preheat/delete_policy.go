// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeletePolicyHandlerFunc turns a function with the right signature into a delete policy handler
type DeletePolicyHandlerFunc func(DeletePolicyParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePolicyHandlerFunc) Handle(params DeletePolicyParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeletePolicyHandler interface for that can handle valid delete policy params
type DeletePolicyHandler interface {
	Handle(DeletePolicyParams, interface{}) middleware.Responder
}

// NewDeletePolicy creates a new http.Handler for the delete policy operation
func NewDeletePolicy(ctx *middleware.Context, handler DeletePolicyHandler) *DeletePolicy {
	return &DeletePolicy{Context: ctx, Handler: handler}
}

/*DeletePolicy swagger:route DELETE /projects/{project_name}/preheat/policies/{preheat_policy_name} preheat deletePolicy

Delete a preheat policy

Delete a preheat policy

*/
type DeletePolicy struct {
	Context *middleware.Context
	Handler DeletePolicyHandler
}

func (o *DeletePolicy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeletePolicyParams()

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