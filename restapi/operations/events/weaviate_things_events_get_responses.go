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

// WeaviateThingsEventsGetOKCode is the HTTP code returned for type WeaviateThingsEventsGetOK
const WeaviateThingsEventsGetOKCode int = 200

/*WeaviateThingsEventsGetOK Successful response.

swagger:response weaviateThingsEventsGetOK
*/
type WeaviateThingsEventsGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.EventGetResponse `json:"body,omitempty"`
}

// NewWeaviateThingsEventsGetOK creates WeaviateThingsEventsGetOK with default headers values
func NewWeaviateThingsEventsGetOK() *WeaviateThingsEventsGetOK {
	return &WeaviateThingsEventsGetOK{}
}

// WithPayload adds the payload to the weaviate things events get o k response
func (o *WeaviateThingsEventsGetOK) WithPayload(payload *models.EventGetResponse) *WeaviateThingsEventsGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate things events get o k response
func (o *WeaviateThingsEventsGetOK) SetPayload(payload *models.EventGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateThingsEventsGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateThingsEventsGetUnauthorizedCode is the HTTP code returned for type WeaviateThingsEventsGetUnauthorized
const WeaviateThingsEventsGetUnauthorizedCode int = 401

/*WeaviateThingsEventsGetUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateThingsEventsGetUnauthorized
*/
type WeaviateThingsEventsGetUnauthorized struct {
}

// NewWeaviateThingsEventsGetUnauthorized creates WeaviateThingsEventsGetUnauthorized with default headers values
func NewWeaviateThingsEventsGetUnauthorized() *WeaviateThingsEventsGetUnauthorized {
	return &WeaviateThingsEventsGetUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateThingsEventsGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// WeaviateThingsEventsGetForbiddenCode is the HTTP code returned for type WeaviateThingsEventsGetForbidden
const WeaviateThingsEventsGetForbiddenCode int = 403

/*WeaviateThingsEventsGetForbidden The used API-key has insufficient permissions.

swagger:response weaviateThingsEventsGetForbidden
*/
type WeaviateThingsEventsGetForbidden struct {
}

// NewWeaviateThingsEventsGetForbidden creates WeaviateThingsEventsGetForbidden with default headers values
func NewWeaviateThingsEventsGetForbidden() *WeaviateThingsEventsGetForbidden {
	return &WeaviateThingsEventsGetForbidden{}
}

// WriteResponse to the client
func (o *WeaviateThingsEventsGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// WeaviateThingsEventsGetNotFoundCode is the HTTP code returned for type WeaviateThingsEventsGetNotFound
const WeaviateThingsEventsGetNotFoundCode int = 404

/*WeaviateThingsEventsGetNotFound Successful query result but no resource was found.

swagger:response weaviateThingsEventsGetNotFound
*/
type WeaviateThingsEventsGetNotFound struct {
}

// NewWeaviateThingsEventsGetNotFound creates WeaviateThingsEventsGetNotFound with default headers values
func NewWeaviateThingsEventsGetNotFound() *WeaviateThingsEventsGetNotFound {
	return &WeaviateThingsEventsGetNotFound{}
}

// WriteResponse to the client
func (o *WeaviateThingsEventsGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// WeaviateThingsEventsGetNotImplementedCode is the HTTP code returned for type WeaviateThingsEventsGetNotImplemented
const WeaviateThingsEventsGetNotImplementedCode int = 501

/*WeaviateThingsEventsGetNotImplemented Not (yet) implemented.

swagger:response weaviateThingsEventsGetNotImplemented
*/
type WeaviateThingsEventsGetNotImplemented struct {
}

// NewWeaviateThingsEventsGetNotImplemented creates WeaviateThingsEventsGetNotImplemented with default headers values
func NewWeaviateThingsEventsGetNotImplemented() *WeaviateThingsEventsGetNotImplemented {
	return &WeaviateThingsEventsGetNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateThingsEventsGetNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
