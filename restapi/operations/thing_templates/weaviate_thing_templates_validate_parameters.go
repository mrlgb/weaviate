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
 package thing_templates




import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/weaviate/weaviate/models"
)

// NewWeaviateThingTemplatesValidateParams creates a new WeaviateThingTemplatesValidateParams object
// with the default values initialized.
func NewWeaviateThingTemplatesValidateParams() WeaviateThingTemplatesValidateParams {
	var ()
	return WeaviateThingTemplatesValidateParams{}
}

// WeaviateThingTemplatesValidateParams contains all the bound params for the weaviate thing templates validate operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.thingTemplates.validate
type WeaviateThingTemplatesValidateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: body
	*/
	Body *models.ThingTemplate
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *WeaviateThingTemplatesValidateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ThingTemplate
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
