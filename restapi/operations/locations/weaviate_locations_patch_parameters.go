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
 package locations




import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/models"
)

// NewWeaviateLocationsPatchParams creates a new WeaviateLocationsPatchParams object
// with the default values initialized.
func NewWeaviateLocationsPatchParams() WeaviateLocationsPatchParams {
	var ()
	return WeaviateLocationsPatchParams{}
}

// WeaviateLocationsPatchParams contains all the bound params for the weaviate locations patch operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.locations.patch
type WeaviateLocationsPatchParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*JSONPatch document as defined by RFC 6902.
	  Required: true
	  In: body
	*/
	Body []*models.PatchDocument
	/*Unique ID of the location.
	  Required: true
	  In: path
	*/
	LocationID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *WeaviateLocationsPatchParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.PatchDocument
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			for _, io := range o.Body {
				if err := io.Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.Body = body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	rLocationID, rhkLocationID, _ := route.Params.GetOK("locationId")
	if err := o.bindLocationID(rLocationID, rhkLocationID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *WeaviateLocationsPatchParams) bindLocationID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.LocationID = raw

	return nil
}
