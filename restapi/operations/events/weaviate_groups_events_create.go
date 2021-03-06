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

// WeaviateGroupsEventsCreateHandlerFunc turns a function with the right signature into a weaviate groups events create handler
type WeaviateGroupsEventsCreateHandlerFunc func(WeaviateGroupsEventsCreateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateGroupsEventsCreateHandlerFunc) Handle(params WeaviateGroupsEventsCreateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateGroupsEventsCreateHandler interface for that can handle valid weaviate groups events create params
type WeaviateGroupsEventsCreateHandler interface {
	Handle(WeaviateGroupsEventsCreateParams, interface{}) middleware.Responder
}

// NewWeaviateGroupsEventsCreate creates a new http.Handler for the weaviate groups events create operation
func NewWeaviateGroupsEventsCreate(ctx *middleware.Context, handler WeaviateGroupsEventsCreateHandler) *WeaviateGroupsEventsCreate {
	return &WeaviateGroupsEventsCreate{Context: ctx, Handler: handler}
}

/*WeaviateGroupsEventsCreate swagger:route POST /groups/{groupId}/events events weaviateGroupsEventsCreate

Create events for a group (also available as MQTT channel) related to this key.

Create event in group.

*/
type WeaviateGroupsEventsCreate struct {
	Context *middleware.Context
	Handler WeaviateGroupsEventsCreateHandler
}

func (o *WeaviateGroupsEventsCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateGroupsEventsCreateParams()

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
