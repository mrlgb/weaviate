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
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package model_manifests




import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveModelManifestsValidateDeviceStateHandlerFunc turns a function with the right signature into a weave model manifests validate device state handler
type WeaveModelManifestsValidateDeviceStateHandlerFunc func(WeaveModelManifestsValidateDeviceStateParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveModelManifestsValidateDeviceStateHandlerFunc) Handle(params WeaveModelManifestsValidateDeviceStateParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveModelManifestsValidateDeviceStateHandler interface for that can handle valid weave model manifests validate device state params
type WeaveModelManifestsValidateDeviceStateHandler interface {
	Handle(WeaveModelManifestsValidateDeviceStateParams, interface{}) middleware.Responder
}

// NewWeaveModelManifestsValidateDeviceState creates a new http.Handler for the weave model manifests validate device state operation
func NewWeaveModelManifestsValidateDeviceState(ctx *middleware.Context, handler WeaveModelManifestsValidateDeviceStateHandler) *WeaveModelManifestsValidateDeviceState {
	return &WeaveModelManifestsValidateDeviceState{Context: ctx, Handler: handler}
}

/*WeaveModelManifestsValidateDeviceState swagger:route POST /modelManifests/validateDeviceState modelManifests weaveModelManifestsValidateDeviceState

Validates given device state object and returns errors.

*/
type WeaveModelManifestsValidateDeviceState struct {
	Context *middleware.Context
	Handler WeaveModelManifestsValidateDeviceStateHandler
}

func (o *WeaveModelManifestsValidateDeviceState) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveModelManifestsValidateDeviceStateParams()

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