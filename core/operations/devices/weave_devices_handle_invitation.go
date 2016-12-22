package devices


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveDevicesHandleInvitationHandlerFunc turns a function with the right signature into a weave devices handle invitation handler
type WeaveDevicesHandleInvitationHandlerFunc func(WeaveDevicesHandleInvitationParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveDevicesHandleInvitationHandlerFunc) Handle(params WeaveDevicesHandleInvitationParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveDevicesHandleInvitationHandler interface for that can handle valid weave devices handle invitation params
type WeaveDevicesHandleInvitationHandler interface {
	Handle(WeaveDevicesHandleInvitationParams, interface{}) middleware.Responder
}

// NewWeaveDevicesHandleInvitation creates a new http.Handler for the weave devices handle invitation operation
func NewWeaveDevicesHandleInvitation(ctx *middleware.Context, handler WeaveDevicesHandleInvitationHandler) *WeaveDevicesHandleInvitation {
	return &WeaveDevicesHandleInvitation{Context: ctx, Handler: handler}
}

/*WeaveDevicesHandleInvitation swagger:route POST /devices/{deviceId}/handleInvitation devices weaveDevicesHandleInvitation

Confirms or rejects a pending device.

*/
type WeaveDevicesHandleInvitation struct {
	Context *middleware.Context
	Handler WeaveDevicesHandleInvitationHandler
}

func (o *WeaveDevicesHandleInvitation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveDevicesHandleInvitationParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
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