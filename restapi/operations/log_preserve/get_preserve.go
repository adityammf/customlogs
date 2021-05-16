// Code generated by go-swagger; DO NOT EDIT.

package log_preserve

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPreserveHandlerFunc turns a function with the right signature into a get preserve handler
type GetPreserveHandlerFunc func(GetPreserveParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPreserveHandlerFunc) Handle(params GetPreserveParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetPreserveHandler interface for that can handle valid get preserve params
type GetPreserveHandler interface {
	Handle(GetPreserveParams, interface{}) middleware.Responder
}

// NewGetPreserve creates a new http.Handler for the get preserve operation
func NewGetPreserve(ctx *middleware.Context, handler GetPreserveHandler) *GetPreserve {
	return &GetPreserve{Context: ctx, Handler: handler}
}

/* GetPreserve swagger:route GET /preserve Log Preserve getPreserve

get response

*/
type GetPreserve struct {
	Context *middleware.Context
	Handler GetPreserveHandler
}

func (o *GetPreserve) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPreserveParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
