// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateRetentionHandlerFunc turns a function with the right signature into a create retention handler
type CreateRetentionHandlerFunc func(CreateRetentionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateRetentionHandlerFunc) Handle(params CreateRetentionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateRetentionHandler interface for that can handle valid create retention params
type CreateRetentionHandler interface {
	Handle(CreateRetentionParams, interface{}) middleware.Responder
}

// NewCreateRetention creates a new http.Handler for the create retention operation
func NewCreateRetention(ctx *middleware.Context, handler CreateRetentionHandler) *CreateRetention {
	return &CreateRetention{Context: ctx, Handler: handler}
}

/*CreateRetention swagger:route POST /retentions Retention createRetention

Create Retention Policy

Create Retention Policy, you can reference metadatas API for the policy model. You can check project metadatas to find whether a retention policy is already binded. This method should only be called when no retention policy binded to project yet.

*/
type CreateRetention struct {
	Context *middleware.Context
	Handler CreateRetentionHandler
}

func (o *CreateRetention) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateRetentionParams()

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