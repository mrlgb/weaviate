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
 package events


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateThingsEventsGetHandlerFunc turns a function with the right signature into a weaviate things events get handler
type WeaviateThingsEventsGetHandlerFunc func(WeaviateThingsEventsGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateThingsEventsGetHandlerFunc) Handle(params WeaviateThingsEventsGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateThingsEventsGetHandler interface for that can handle valid weaviate things events get params
type WeaviateThingsEventsGetHandler interface {
	Handle(WeaviateThingsEventsGetParams, interface{}) middleware.Responder
}

// NewWeaviateThingsEventsGet creates a new http.Handler for the weaviate things events get operation
func NewWeaviateThingsEventsGet(ctx *middleware.Context, handler WeaviateThingsEventsGetHandler) *WeaviateThingsEventsGet {
	return &WeaviateThingsEventsGet{Context: ctx, Handler: handler}
}

/*WeaviateThingsEventsGet swagger:route GET /things/{thingId}/events/{eventId} events weaviateThingsEventsGet

Get a specific event based on its uuid and a thing uuid related to this key.

Lists events.

*/
type WeaviateThingsEventsGet struct {
	Context *middleware.Context
	Handler WeaviateThingsEventsGetHandler
}

func (o *WeaviateThingsEventsGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateThingsEventsGetParams()

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
