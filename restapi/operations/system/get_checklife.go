// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetChecklifeHandlerFunc turns a function with the right signature into a get checklife handler
type GetChecklifeHandlerFunc func(GetChecklifeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetChecklifeHandlerFunc) Handle(params GetChecklifeParams) middleware.Responder {
	return fn(params)
}

// GetChecklifeHandler interface for that can handle valid get checklife params
type GetChecklifeHandler interface {
	Handle(GetChecklifeParams) middleware.Responder
}

// NewGetChecklife creates a new http.Handler for the get checklife operation
func NewGetChecklife(ctx *middleware.Context, handler GetChecklifeHandler) *GetChecklife {
	return &GetChecklife{Context: ctx, Handler: handler}
}

/* GetChecklife swagger:route GET /checklife System getChecklife

get response

*/
type GetChecklife struct {
	Context *middleware.Context
	Handler GetChecklifeHandler
}

func (o *GetChecklife) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetChecklifeParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}