// Code generated by go-swagger; DO NOT EDIT.

package authors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllAuthorsHandlerFunc turns a function with the right signature into a get all authors handler
type GetAllAuthorsHandlerFunc func(GetAllAuthorsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllAuthorsHandlerFunc) Handle(params GetAllAuthorsParams) middleware.Responder {
	return fn(params)
}

// GetAllAuthorsHandler interface for that can handle valid get all authors params
type GetAllAuthorsHandler interface {
	Handle(GetAllAuthorsParams) middleware.Responder
}

// NewGetAllAuthors creates a new http.Handler for the get all authors operation
func NewGetAllAuthors(ctx *middleware.Context, handler GetAllAuthorsHandler) *GetAllAuthors {
	return &GetAllAuthors{Context: ctx, Handler: handler}
}

/*
	GetAllAuthors swagger:route GET /authors Authors getAllAuthors

get all authors
*/
type GetAllAuthors struct {
	Context *middleware.Context
	Handler GetAllAuthorsHandler
}

func (o *GetAllAuthors) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAllAuthorsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
