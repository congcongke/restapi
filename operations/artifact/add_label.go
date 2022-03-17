// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddLabelHandlerFunc turns a function with the right signature into a add label handler
type AddLabelHandlerFunc func(AddLabelParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn AddLabelHandlerFunc) Handle(params AddLabelParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// AddLabelHandler interface for that can handle valid add label params
type AddLabelHandler interface {
	Handle(AddLabelParams, interface{}) middleware.Responder
}

// NewAddLabel creates a new http.Handler for the add label operation
func NewAddLabel(ctx *middleware.Context, handler AddLabelHandler) *AddLabel {
	return &AddLabel{Context: ctx, Handler: handler}
}

/*AddLabel swagger:route POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels artifact addLabel

Add label to artifact

Add label to the specified artiact.

*/
type AddLabel struct {
	Context *middleware.Context
	Handler AddLabelHandler
}

func (o *AddLabel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddLabelParams()

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
