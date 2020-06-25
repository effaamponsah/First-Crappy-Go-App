// Code generated by go-swagger; DO NOT EDIT.

package foods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateFoodHandlerFunc turns a function with the right signature into a update food handler
type UpdateFoodHandlerFunc func(UpdateFoodParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateFoodHandlerFunc) Handle(params UpdateFoodParams) middleware.Responder {
	return fn(params)
}

// UpdateFoodHandler interface for that can handle valid update food params
type UpdateFoodHandler interface {
	Handle(UpdateFoodParams) middleware.Responder
}

// NewUpdateFood creates a new http.Handler for the update food operation
func NewUpdateFood(ctx *middleware.Context, handler UpdateFoodHandler) *UpdateFood {
	return &UpdateFood{Context: ctx, Handler: handler}
}

/*UpdateFood swagger:route PUT /food/{food_id} foods updateFood

UpdateFood update food API

*/
type UpdateFood struct {
	Context *middleware.Context
	Handler UpdateFoodHandler
}

func (o *UpdateFood) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateFoodParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
