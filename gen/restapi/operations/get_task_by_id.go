// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTaskByIDHandlerFunc turns a function with the right signature into a get task by ID handler
type GetTaskByIDHandlerFunc func(GetTaskByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTaskByIDHandlerFunc) Handle(params GetTaskByIDParams) middleware.Responder {
	return fn(params)
}

// GetTaskByIDHandler interface for that can handle valid get task by ID params
type GetTaskByIDHandler interface {
	Handle(GetTaskByIDParams) middleware.Responder
}

// NewGetTaskByID creates a new http.Handler for the get task by ID operation
func NewGetTaskByID(ctx *middleware.Context, handler GetTaskByIDHandler) *GetTaskByID {
	return &GetTaskByID{Context: ctx, Handler: handler}
}

/* GetTaskByID swagger:route GET /task/{ID} getTaskById

GetTaskByID get task by ID API

*/
type GetTaskByID struct {
	Context *middleware.Context
	Handler GetTaskByIDHandler
}

func (o *GetTaskByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTaskByIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
