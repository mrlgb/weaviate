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
 package commands


// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaviateCommandsValidateHandlerFunc turns a function with the right signature into a weaviate commands validate handler
type WeaviateCommandsValidateHandlerFunc func(WeaviateCommandsValidateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateCommandsValidateHandlerFunc) Handle(params WeaviateCommandsValidateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaviateCommandsValidateHandler interface for that can handle valid weaviate commands validate params
type WeaviateCommandsValidateHandler interface {
	Handle(WeaviateCommandsValidateParams, interface{}) middleware.Responder
}

// NewWeaviateCommandsValidate creates a new http.Handler for the weaviate commands validate operation
func NewWeaviateCommandsValidate(ctx *middleware.Context, handler WeaviateCommandsValidateHandler) *WeaviateCommandsValidate {
	return &WeaviateCommandsValidate{Context: ctx, Handler: handler}
}

/*WeaviateCommandsValidate swagger:route POST /commands/validate commands weaviateCommandsValidate

Validate a command object.

Validate a command.

*/
type WeaviateCommandsValidate struct {
	Context *middleware.Context
	Handler WeaviateCommandsValidateHandler
}

func (o *WeaviateCommandsValidate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaviateCommandsValidateParams()

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
