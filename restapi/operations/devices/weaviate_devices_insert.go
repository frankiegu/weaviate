/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package devices


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateDevicesInsertHandlerFunc turns a function with the right signature into a weaviate devices insert handler
type WeaviateDevicesInsertHandlerFunc func(WeaviateDevicesInsertParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateDevicesInsertHandlerFunc) Handle(params WeaviateDevicesInsertParams) middleware.Responder {
	return fn(params)
}

// WeaviateDevicesInsertHandler interface for that can handle valid weaviate devices insert params
type WeaviateDevicesInsertHandler interface {
	Handle(WeaviateDevicesInsertParams) middleware.Responder
}

// NewWeaviateDevicesInsert creates a new http.Handler for the weaviate devices insert operation
func NewWeaviateDevicesInsert(ctx *middleware.Context, handler WeaviateDevicesInsertHandler) *WeaviateDevicesInsert {
	return &WeaviateDevicesInsert{Context: ctx, Handler: handler}
}

/*WeaviateDevicesInsert swagger:route POST /devices devices weaviateDevicesInsert

Registers a new device. This method may be used only by aggregator devices or adapters.

*/
type WeaviateDevicesInsert struct {
	Context *middleware.Context
	Handler WeaviateDevicesInsertHandler
}

func (o *WeaviateDevicesInsert) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateDevicesInsertParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
