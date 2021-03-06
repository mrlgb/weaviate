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




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateGroupsEventsCreateAcceptedCode is the HTTP code returned for type WeaviateGroupsEventsCreateAccepted
const WeaviateGroupsEventsCreateAcceptedCode int = 202

/*WeaviateGroupsEventsCreateAccepted Successfully received.

swagger:response weaviateGroupsEventsCreateAccepted
*/
type WeaviateGroupsEventsCreateAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.EventGetResponse `json:"body,omitempty"`
}

// NewWeaviateGroupsEventsCreateAccepted creates WeaviateGroupsEventsCreateAccepted with default headers values
func NewWeaviateGroupsEventsCreateAccepted() *WeaviateGroupsEventsCreateAccepted {
	return &WeaviateGroupsEventsCreateAccepted{}
}

// WithPayload adds the payload to the weaviate groups events create accepted response
func (o *WeaviateGroupsEventsCreateAccepted) WithPayload(payload *models.EventGetResponse) *WeaviateGroupsEventsCreateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate groups events create accepted response
func (o *WeaviateGroupsEventsCreateAccepted) SetPayload(payload *models.EventGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateGroupsEventsCreateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateGroupsEventsCreateUnauthorizedCode is the HTTP code returned for type WeaviateGroupsEventsCreateUnauthorized
const WeaviateGroupsEventsCreateUnauthorizedCode int = 401

/*WeaviateGroupsEventsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateGroupsEventsCreateUnauthorized
*/
type WeaviateGroupsEventsCreateUnauthorized struct {
}

// NewWeaviateGroupsEventsCreateUnauthorized creates WeaviateGroupsEventsCreateUnauthorized with default headers values
func NewWeaviateGroupsEventsCreateUnauthorized() *WeaviateGroupsEventsCreateUnauthorized {
	return &WeaviateGroupsEventsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateGroupsEventsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateGroupsEventsCreateForbiddenCode is the HTTP code returned for type WeaviateGroupsEventsCreateForbidden
const WeaviateGroupsEventsCreateForbiddenCode int = 403

/*WeaviateGroupsEventsCreateForbidden The used API-key has insufficient permissions.

swagger:response weaviateGroupsEventsCreateForbidden
*/
type WeaviateGroupsEventsCreateForbidden struct {
}

// NewWeaviateGroupsEventsCreateForbidden creates WeaviateGroupsEventsCreateForbidden with default headers values
func NewWeaviateGroupsEventsCreateForbidden() *WeaviateGroupsEventsCreateForbidden {
	return &WeaviateGroupsEventsCreateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateGroupsEventsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateGroupsEventsCreateUnprocessableEntityCode is the HTTP code returned for type WeaviateGroupsEventsCreateUnprocessableEntity
const WeaviateGroupsEventsCreateUnprocessableEntityCode int = 422

/*WeaviateGroupsEventsCreateUnprocessableEntity Can not execute this command, because the commandParameters{} are set incorrectly.

swagger:response weaviateGroupsEventsCreateUnprocessableEntity
*/
type WeaviateGroupsEventsCreateUnprocessableEntity struct {
}

// NewWeaviateGroupsEventsCreateUnprocessableEntity creates WeaviateGroupsEventsCreateUnprocessableEntity with default headers values
func NewWeaviateGroupsEventsCreateUnprocessableEntity() *WeaviateGroupsEventsCreateUnprocessableEntity {
	return &WeaviateGroupsEventsCreateUnprocessableEntity{}
}

// WriteResponse to the client
func (o *WeaviateGroupsEventsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
}

// WeaviateGroupsEventsCreateNotImplementedCode is the HTTP code returned for type WeaviateGroupsEventsCreateNotImplemented
const WeaviateGroupsEventsCreateNotImplementedCode int = 501

/*WeaviateGroupsEventsCreateNotImplemented Not (yet) implemented.

swagger:response weaviateGroupsEventsCreateNotImplemented
*/
type WeaviateGroupsEventsCreateNotImplemented struct {
}

// NewWeaviateGroupsEventsCreateNotImplemented creates WeaviateGroupsEventsCreateNotImplemented with default headers values
func NewWeaviateGroupsEventsCreateNotImplemented() *WeaviateGroupsEventsCreateNotImplemented {
	return &WeaviateGroupsEventsCreateNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateGroupsEventsCreateNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
