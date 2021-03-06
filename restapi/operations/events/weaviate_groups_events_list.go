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

// WeaviateGroupsEventsListHandlerFunc turns a function with the right signature into a weaviate groups events list handler
type WeaviateGroupsEventsListHandlerFunc func(WeaviateGroupsEventsListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateGroupsEventsListHandlerFunc) Handle(params WeaviateGroupsEventsListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateGroupsEventsListHandler interface for that can handle valid weaviate groups events list params
type WeaviateGroupsEventsListHandler interface {
	Handle(WeaviateGroupsEventsListParams, interface{}) middleware.Responder
}

// NewWeaviateGroupsEventsList creates a new http.Handler for the weaviate groups events list operation
func NewWeaviateGroupsEventsList(ctx *middleware.Context, handler WeaviateGroupsEventsListHandler) *WeaviateGroupsEventsList {
	return &WeaviateGroupsEventsList{Context: ctx, Handler: handler}
}

/*WeaviateGroupsEventsList swagger:route GET /groups/{groupId}/events events weaviateGroupsEventsList

Get a list of events based on a groups's uuid (also available as MQTT channel) related to this key.

Lists events.

*/
type WeaviateGroupsEventsList struct {
	Context *middleware.Context
	Handler WeaviateGroupsEventsListHandler
}

func (o *WeaviateGroupsEventsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateGroupsEventsListParams()

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
