// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListReplicationExecutionsHandlerFunc turns a function with the right signature into a list replication executions handler
type ListReplicationExecutionsHandlerFunc func(ListReplicationExecutionsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ListReplicationExecutionsHandlerFunc) Handle(params ListReplicationExecutionsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ListReplicationExecutionsHandler interface for that can handle valid list replication executions params
type ListReplicationExecutionsHandler interface {
	Handle(ListReplicationExecutionsParams, interface{}) middleware.Responder
}

// NewListReplicationExecutions creates a new http.Handler for the list replication executions operation
func NewListReplicationExecutions(ctx *middleware.Context, handler ListReplicationExecutionsHandler) *ListReplicationExecutions {
	return &ListReplicationExecutions{Context: ctx, Handler: handler}
}

/*ListReplicationExecutions swagger:route GET /replication/executions replication listReplicationExecutions

List replication executions

List replication executions

*/
type ListReplicationExecutions struct {
	Context *middleware.Context
	Handler ListReplicationExecutionsHandler
}

func (o *ListReplicationExecutions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListReplicationExecutionsParams()

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