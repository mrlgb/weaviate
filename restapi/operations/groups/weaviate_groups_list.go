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
 package groups


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateGroupsListHandlerFunc turns a function with the right signature into a weaviate groups list handler
type WeaviateGroupsListHandlerFunc func(WeaviateGroupsListParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateGroupsListHandlerFunc) Handle(params WeaviateGroupsListParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateGroupsListHandler interface for that can handle valid weaviate groups list params
type WeaviateGroupsListHandler interface {
	Handle(WeaviateGroupsListParams, interface{}) middleware.Responder
}

// NewWeaviateGroupsList creates a new http.Handler for the weaviate groups list operation
func NewWeaviateGroupsList(ctx *middleware.Context, handler WeaviateGroupsListHandler) *WeaviateGroupsList {
	return &WeaviateGroupsList{Context: ctx, Handler: handler}
}

/*WeaviateGroupsList swagger:route GET /groups groups weaviateGroupsList

Get a list of groups related to this key.

Lists all groups.

*/
type WeaviateGroupsList struct {
	Context *middleware.Context
	Handler WeaviateGroupsListHandler
}

func (o *WeaviateGroupsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateGroupsListParams()

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
