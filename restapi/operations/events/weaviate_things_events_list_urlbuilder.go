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
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// WeaviateThingsEventsListURL generates an URL for the weaviate things events list operation
type WeaviateThingsEventsListURL struct {
	ThingID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *WeaviateThingsEventsListURL) WithBasePath(bp string) *WeaviateThingsEventsListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *WeaviateThingsEventsListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *WeaviateThingsEventsListURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/things/{thingId}/events"

	thingID := o.ThingID
	if thingID != "" {
		_path = strings.Replace(_path, "{thingId}", thingID, -1)
	} else {
		return nil, errors.New("ThingID is required on WeaviateThingsEventsListURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/weaviate/v1"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *WeaviateThingsEventsListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *WeaviateThingsEventsListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *WeaviateThingsEventsListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on WeaviateThingsEventsListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on WeaviateThingsEventsListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *WeaviateThingsEventsListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
