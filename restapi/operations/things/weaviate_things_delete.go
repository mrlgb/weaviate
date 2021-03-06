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
 package things


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateThingsDeleteHandlerFunc turns a function with the right signature into a weaviate things delete handler
type WeaviateThingsDeleteHandlerFunc func(WeaviateThingsDeleteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateThingsDeleteHandlerFunc) Handle(params WeaviateThingsDeleteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateThingsDeleteHandler interface for that can handle valid weaviate things delete params
type WeaviateThingsDeleteHandler interface {
	Handle(WeaviateThingsDeleteParams, interface{}) middleware.Responder
}

// NewWeaviateThingsDelete creates a new http.Handler for the weaviate things delete operation
func NewWeaviateThingsDelete(ctx *middleware.Context, handler WeaviateThingsDeleteHandler) *WeaviateThingsDelete {
	return &WeaviateThingsDelete{Context: ctx, Handler: handler}
}

/*WeaviateThingsDelete swagger:route DELETE /things/{thingId} things weaviateThingsDelete

Delete a thing based on its uuid related to this key.

Deletes a thing from the system.

*/
type WeaviateThingsDelete struct {
	Context *middleware.Context
	Handler WeaviateThingsDeleteHandler
}

func (o *WeaviateThingsDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateThingsDeleteParams()

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
