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
 package operations




import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	"github.com/go-openapi/runtime/yamlpc"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/weaviate/weaviate/restapi/operations/commands"
	"github.com/weaviate/weaviate/restapi/operations/events"
	"github.com/weaviate/weaviate/restapi/operations/groups"
	"github.com/weaviate/weaviate/restapi/operations/keys"
	"github.com/weaviate/weaviate/restapi/operations/locations"
	"github.com/weaviate/weaviate/restapi/operations/thing_templates"
	"github.com/weaviate/weaviate/restapi/operations/things"
)

// NewWeaviateAPI creates a new Weaviate instance
func NewWeaviateAPI(spec *loads.Document) *WeaviateAPI {
	return &WeaviateAPI{
		handlers:              make(map[string]map[string]http.Handler),
		formats:               strfmt.Default,
		defaultConsumes:       "application/json",
		defaultProduces:       "application/json",
		ServerShutdown:        func() {},
		spec:                  spec,
		ServeError:            errors.ServeError,
		JSONConsumer:          runtime.JSONConsumer(),
		BinConsumer:           runtime.ByteStreamConsumer(),
		UrlformConsumer:       runtime.DiscardConsumer,
		YamlConsumer:          yamlpc.YAMLConsumer(),
		XMLConsumer:           runtime.XMLConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		TxtConsumer:           runtime.TextConsumer(),
		JSONProducer:          runtime.JSONProducer(),
		BinProducer:           runtime.ByteStreamProducer(),
		UrlformProducer:       runtime.DiscardProducer,
		YamlProducer:          yamlpc.YAMLProducer(),
		XMLProducer:           runtime.XMLProducer(),
		MultipartformProducer: runtime.DiscardProducer,
		TxtProducer:           runtime.TextProducer(),
		CommandsWeaviateCommandsDeleteHandler: commands.WeaviateCommandsDeleteHandlerFunc(func(params commands.WeaviateCommandsDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CommandsWeaviateCommandsDelete has not yet been implemented")
		}),
		CommandsWeaviateCommandsGetHandler: commands.WeaviateCommandsGetHandlerFunc(func(params commands.WeaviateCommandsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CommandsWeaviateCommandsGet has not yet been implemented")
		}),
		CommandsWeaviateCommandsInsertHandler: commands.WeaviateCommandsInsertHandlerFunc(func(params commands.WeaviateCommandsInsertParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CommandsWeaviateCommandsInsert has not yet been implemented")
		}),
		CommandsWeaviateCommandsListHandler: commands.WeaviateCommandsListHandlerFunc(func(params commands.WeaviateCommandsListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CommandsWeaviateCommandsList has not yet been implemented")
		}),
		CommandsWeaviateCommandsPatchHandler: commands.WeaviateCommandsPatchHandlerFunc(func(params commands.WeaviateCommandsPatchParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CommandsWeaviateCommandsPatch has not yet been implemented")
		}),
		CommandsWeaviateCommandsUpdateHandler: commands.WeaviateCommandsUpdateHandlerFunc(func(params commands.WeaviateCommandsUpdateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CommandsWeaviateCommandsUpdate has not yet been implemented")
		}),
		CommandsWeaviateCommandsValidateHandler: commands.WeaviateCommandsValidateHandlerFunc(func(params commands.WeaviateCommandsValidateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CommandsWeaviateCommandsValidate has not yet been implemented")
		}),
		GroupsWeaviateGroupsDeleteHandler: groups.WeaviateGroupsDeleteHandlerFunc(func(params groups.WeaviateGroupsDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GroupsWeaviateGroupsDelete has not yet been implemented")
		}),
		EventsWeaviateGroupsEventsCreateHandler: events.WeaviateGroupsEventsCreateHandlerFunc(func(params events.WeaviateGroupsEventsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation EventsWeaviateGroupsEventsCreate has not yet been implemented")
		}),
		EventsWeaviateGroupsEventsGetHandler: events.WeaviateGroupsEventsGetHandlerFunc(func(params events.WeaviateGroupsEventsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation EventsWeaviateGroupsEventsGet has not yet been implemented")
		}),
		EventsWeaviateGroupsEventsListHandler: events.WeaviateGroupsEventsListHandlerFunc(func(params events.WeaviateGroupsEventsListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation EventsWeaviateGroupsEventsList has not yet been implemented")
		}),
		GroupsWeaviateGroupsGetHandler: groups.WeaviateGroupsGetHandlerFunc(func(params groups.WeaviateGroupsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GroupsWeaviateGroupsGet has not yet been implemented")
		}),
		GroupsWeaviateGroupsInsertHandler: groups.WeaviateGroupsInsertHandlerFunc(func(params groups.WeaviateGroupsInsertParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GroupsWeaviateGroupsInsert has not yet been implemented")
		}),
		GroupsWeaviateGroupsListHandler: groups.WeaviateGroupsListHandlerFunc(func(params groups.WeaviateGroupsListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GroupsWeaviateGroupsList has not yet been implemented")
		}),
		GroupsWeaviateGroupsPatchHandler: groups.WeaviateGroupsPatchHandlerFunc(func(params groups.WeaviateGroupsPatchParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GroupsWeaviateGroupsPatch has not yet been implemented")
		}),
		GroupsWeaviateGroupsUpdateHandler: groups.WeaviateGroupsUpdateHandlerFunc(func(params groups.WeaviateGroupsUpdateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GroupsWeaviateGroupsUpdate has not yet been implemented")
		}),
		KeysWeaviateKeyCreateHandler: keys.WeaviateKeyCreateHandlerFunc(func(params keys.WeaviateKeyCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation KeysWeaviateKeyCreate has not yet been implemented")
		}),
		KeysWeaviateKeysChildrenGetHandler: keys.WeaviateKeysChildrenGetHandlerFunc(func(params keys.WeaviateKeysChildrenGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation KeysWeaviateKeysChildrenGet has not yet been implemented")
		}),
		KeysWeaviateKeysDeleteHandler: keys.WeaviateKeysDeleteHandlerFunc(func(params keys.WeaviateKeysDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation KeysWeaviateKeysDelete has not yet been implemented")
		}),
		KeysWeaviateKeysGetHandler: keys.WeaviateKeysGetHandlerFunc(func(params keys.WeaviateKeysGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation KeysWeaviateKeysGet has not yet been implemented")
		}),
		LocationsWeaviateLocationsDeleteHandler: locations.WeaviateLocationsDeleteHandlerFunc(func(params locations.WeaviateLocationsDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation LocationsWeaviateLocationsDelete has not yet been implemented")
		}),
		LocationsWeaviateLocationsGetHandler: locations.WeaviateLocationsGetHandlerFunc(func(params locations.WeaviateLocationsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation LocationsWeaviateLocationsGet has not yet been implemented")
		}),
		LocationsWeaviateLocationsInsertHandler: locations.WeaviateLocationsInsertHandlerFunc(func(params locations.WeaviateLocationsInsertParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation LocationsWeaviateLocationsInsert has not yet been implemented")
		}),
		LocationsWeaviateLocationsListHandler: locations.WeaviateLocationsListHandlerFunc(func(params locations.WeaviateLocationsListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation LocationsWeaviateLocationsList has not yet been implemented")
		}),
		LocationsWeaviateLocationsPatchHandler: locations.WeaviateLocationsPatchHandlerFunc(func(params locations.WeaviateLocationsPatchParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation LocationsWeaviateLocationsPatch has not yet been implemented")
		}),
		LocationsWeaviateLocationsUpdateHandler: locations.WeaviateLocationsUpdateHandlerFunc(func(params locations.WeaviateLocationsUpdateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation LocationsWeaviateLocationsUpdate has not yet been implemented")
		}),
		ThingTemplatesWeaviateThingTemplatesCreateHandler: thing_templates.WeaviateThingTemplatesCreateHandlerFunc(func(params thing_templates.WeaviateThingTemplatesCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingTemplatesWeaviateThingTemplatesCreate has not yet been implemented")
		}),
		ThingTemplatesWeaviateThingTemplatesDeleteHandler: thing_templates.WeaviateThingTemplatesDeleteHandlerFunc(func(params thing_templates.WeaviateThingTemplatesDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingTemplatesWeaviateThingTemplatesDelete has not yet been implemented")
		}),
		ThingTemplatesWeaviateThingTemplatesGetHandler: thing_templates.WeaviateThingTemplatesGetHandlerFunc(func(params thing_templates.WeaviateThingTemplatesGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingTemplatesWeaviateThingTemplatesGet has not yet been implemented")
		}),
		ThingTemplatesWeaviateThingTemplatesListHandler: thing_templates.WeaviateThingTemplatesListHandlerFunc(func(params thing_templates.WeaviateThingTemplatesListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingTemplatesWeaviateThingTemplatesList has not yet been implemented")
		}),
		ThingTemplatesWeaviateThingTemplatesPatchHandler: thing_templates.WeaviateThingTemplatesPatchHandlerFunc(func(params thing_templates.WeaviateThingTemplatesPatchParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingTemplatesWeaviateThingTemplatesPatch has not yet been implemented")
		}),
		ThingTemplatesWeaviateThingTemplatesUpdateHandler: thing_templates.WeaviateThingTemplatesUpdateHandlerFunc(func(params thing_templates.WeaviateThingTemplatesUpdateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingTemplatesWeaviateThingTemplatesUpdate has not yet been implemented")
		}),
		ThingTemplatesWeaviateThingTemplatesValidateHandler: thing_templates.WeaviateThingTemplatesValidateHandlerFunc(func(params thing_templates.WeaviateThingTemplatesValidateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingTemplatesWeaviateThingTemplatesValidate has not yet been implemented")
		}),
		ThingsWeaviateThingsDeleteHandler: things.WeaviateThingsDeleteHandlerFunc(func(params things.WeaviateThingsDeleteParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingsWeaviateThingsDelete has not yet been implemented")
		}),
		EventsWeaviateThingsEventsCreateHandler: events.WeaviateThingsEventsCreateHandlerFunc(func(params events.WeaviateThingsEventsCreateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation EventsWeaviateThingsEventsCreate has not yet been implemented")
		}),
		EventsWeaviateThingsEventsGetHandler: events.WeaviateThingsEventsGetHandlerFunc(func(params events.WeaviateThingsEventsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation EventsWeaviateThingsEventsGet has not yet been implemented")
		}),
		EventsWeaviateThingsEventsListHandler: events.WeaviateThingsEventsListHandlerFunc(func(params events.WeaviateThingsEventsListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation EventsWeaviateThingsEventsList has not yet been implemented")
		}),
		ThingsWeaviateThingsGetHandler: things.WeaviateThingsGetHandlerFunc(func(params things.WeaviateThingsGetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingsWeaviateThingsGet has not yet been implemented")
		}),
		ThingsWeaviateThingsInsertHandler: things.WeaviateThingsInsertHandlerFunc(func(params things.WeaviateThingsInsertParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingsWeaviateThingsInsert has not yet been implemented")
		}),
		ThingsWeaviateThingsListHandler: things.WeaviateThingsListHandlerFunc(func(params things.WeaviateThingsListParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingsWeaviateThingsList has not yet been implemented")
		}),
		ThingsWeaviateThingsPatchHandler: things.WeaviateThingsPatchHandlerFunc(func(params things.WeaviateThingsPatchParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingsWeaviateThingsPatch has not yet been implemented")
		}),
		ThingsWeaviateThingsUpdateHandler: things.WeaviateThingsUpdateHandlerFunc(func(params things.WeaviateThingsUpdateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ThingsWeaviateThingsUpdate has not yet been implemented")
		}),

		// Applies when the "X-API-KEY" header is set
		APIKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (apiKey) X-API-KEY from header param [X-API-KEY] has not yet been implemented")
		},
	}
}

/*WeaviateAPI Internet of Things platform that lets you register, view and manage cloud ready things. */
type WeaviateAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	// JSONConsumer registers a consumer for a "application/json-patch+json" mime type
	JSONConsumer runtime.Consumer
	// BinConsumer registers a consumer for a "application/octet-stream" mime type
	BinConsumer runtime.Consumer
	// UrlformConsumer registers a consumer for a "application/x-www-form-urlencoded" mime type
	UrlformConsumer runtime.Consumer
	// YamlConsumer registers a consumer for a "application/x-yaml" mime type
	YamlConsumer runtime.Consumer
	// XMLConsumer registers a consumer for a "application/xml" mime type
	XMLConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer
	// TxtConsumer registers a consumer for a "text/plain" mime type
	TxtConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer
	// BinProducer registers a producer for a "application/octet-stream" mime type
	BinProducer runtime.Producer
	// UrlformProducer registers a producer for a "application/x-www-form-urlencoded" mime type
	UrlformProducer runtime.Producer
	// YamlProducer registers a producer for a "application/x-yaml" mime type
	YamlProducer runtime.Producer
	// XMLProducer registers a producer for a "application/xml" mime type
	XMLProducer runtime.Producer
	// MultipartformProducer registers a producer for a "multipart/form-data" mime type
	MultipartformProducer runtime.Producer
	// TxtProducer registers a producer for a "text/plain" mime type
	TxtProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key X-API-KEY provided in the header
	APIKeyAuth func(string) (interface{}, error)

	// CommandsWeaviateCommandsDeleteHandler sets the operation handler for the weaviate commands delete operation
	CommandsWeaviateCommandsDeleteHandler commands.WeaviateCommandsDeleteHandler
	// CommandsWeaviateCommandsGetHandler sets the operation handler for the weaviate commands get operation
	CommandsWeaviateCommandsGetHandler commands.WeaviateCommandsGetHandler
	// CommandsWeaviateCommandsInsertHandler sets the operation handler for the weaviate commands insert operation
	CommandsWeaviateCommandsInsertHandler commands.WeaviateCommandsInsertHandler
	// CommandsWeaviateCommandsListHandler sets the operation handler for the weaviate commands list operation
	CommandsWeaviateCommandsListHandler commands.WeaviateCommandsListHandler
	// CommandsWeaviateCommandsPatchHandler sets the operation handler for the weaviate commands patch operation
	CommandsWeaviateCommandsPatchHandler commands.WeaviateCommandsPatchHandler
	// CommandsWeaviateCommandsUpdateHandler sets the operation handler for the weaviate commands update operation
	CommandsWeaviateCommandsUpdateHandler commands.WeaviateCommandsUpdateHandler
	// CommandsWeaviateCommandsValidateHandler sets the operation handler for the weaviate commands validate operation
	CommandsWeaviateCommandsValidateHandler commands.WeaviateCommandsValidateHandler
	// GroupsWeaviateGroupsDeleteHandler sets the operation handler for the weaviate groups delete operation
	GroupsWeaviateGroupsDeleteHandler groups.WeaviateGroupsDeleteHandler
	// EventsWeaviateGroupsEventsCreateHandler sets the operation handler for the weaviate groups events create operation
	EventsWeaviateGroupsEventsCreateHandler events.WeaviateGroupsEventsCreateHandler
	// EventsWeaviateGroupsEventsGetHandler sets the operation handler for the weaviate groups events get operation
	EventsWeaviateGroupsEventsGetHandler events.WeaviateGroupsEventsGetHandler
	// EventsWeaviateGroupsEventsListHandler sets the operation handler for the weaviate groups events list operation
	EventsWeaviateGroupsEventsListHandler events.WeaviateGroupsEventsListHandler
	// GroupsWeaviateGroupsGetHandler sets the operation handler for the weaviate groups get operation
	GroupsWeaviateGroupsGetHandler groups.WeaviateGroupsGetHandler
	// GroupsWeaviateGroupsInsertHandler sets the operation handler for the weaviate groups insert operation
	GroupsWeaviateGroupsInsertHandler groups.WeaviateGroupsInsertHandler
	// GroupsWeaviateGroupsListHandler sets the operation handler for the weaviate groups list operation
	GroupsWeaviateGroupsListHandler groups.WeaviateGroupsListHandler
	// GroupsWeaviateGroupsPatchHandler sets the operation handler for the weaviate groups patch operation
	GroupsWeaviateGroupsPatchHandler groups.WeaviateGroupsPatchHandler
	// GroupsWeaviateGroupsUpdateHandler sets the operation handler for the weaviate groups update operation
	GroupsWeaviateGroupsUpdateHandler groups.WeaviateGroupsUpdateHandler
	// KeysWeaviateKeyCreateHandler sets the operation handler for the weaviate key create operation
	KeysWeaviateKeyCreateHandler keys.WeaviateKeyCreateHandler
	// KeysWeaviateKeysChildrenGetHandler sets the operation handler for the weaviate keys children get operation
	KeysWeaviateKeysChildrenGetHandler keys.WeaviateKeysChildrenGetHandler
	// KeysWeaviateKeysDeleteHandler sets the operation handler for the weaviate keys delete operation
	KeysWeaviateKeysDeleteHandler keys.WeaviateKeysDeleteHandler
	// KeysWeaviateKeysGetHandler sets the operation handler for the weaviate keys get operation
	KeysWeaviateKeysGetHandler keys.WeaviateKeysGetHandler
	// LocationsWeaviateLocationsDeleteHandler sets the operation handler for the weaviate locations delete operation
	LocationsWeaviateLocationsDeleteHandler locations.WeaviateLocationsDeleteHandler
	// LocationsWeaviateLocationsGetHandler sets the operation handler for the weaviate locations get operation
	LocationsWeaviateLocationsGetHandler locations.WeaviateLocationsGetHandler
	// LocationsWeaviateLocationsInsertHandler sets the operation handler for the weaviate locations insert operation
	LocationsWeaviateLocationsInsertHandler locations.WeaviateLocationsInsertHandler
	// LocationsWeaviateLocationsListHandler sets the operation handler for the weaviate locations list operation
	LocationsWeaviateLocationsListHandler locations.WeaviateLocationsListHandler
	// LocationsWeaviateLocationsPatchHandler sets the operation handler for the weaviate locations patch operation
	LocationsWeaviateLocationsPatchHandler locations.WeaviateLocationsPatchHandler
	// LocationsWeaviateLocationsUpdateHandler sets the operation handler for the weaviate locations update operation
	LocationsWeaviateLocationsUpdateHandler locations.WeaviateLocationsUpdateHandler
	// ThingTemplatesWeaviateThingTemplatesCreateHandler sets the operation handler for the weaviate thing templates create operation
	ThingTemplatesWeaviateThingTemplatesCreateHandler thing_templates.WeaviateThingTemplatesCreateHandler
	// ThingTemplatesWeaviateThingTemplatesDeleteHandler sets the operation handler for the weaviate thing templates delete operation
	ThingTemplatesWeaviateThingTemplatesDeleteHandler thing_templates.WeaviateThingTemplatesDeleteHandler
	// ThingTemplatesWeaviateThingTemplatesGetHandler sets the operation handler for the weaviate thing templates get operation
	ThingTemplatesWeaviateThingTemplatesGetHandler thing_templates.WeaviateThingTemplatesGetHandler
	// ThingTemplatesWeaviateThingTemplatesListHandler sets the operation handler for the weaviate thing templates list operation
	ThingTemplatesWeaviateThingTemplatesListHandler thing_templates.WeaviateThingTemplatesListHandler
	// ThingTemplatesWeaviateThingTemplatesPatchHandler sets the operation handler for the weaviate thing templates patch operation
	ThingTemplatesWeaviateThingTemplatesPatchHandler thing_templates.WeaviateThingTemplatesPatchHandler
	// ThingTemplatesWeaviateThingTemplatesUpdateHandler sets the operation handler for the weaviate thing templates update operation
	ThingTemplatesWeaviateThingTemplatesUpdateHandler thing_templates.WeaviateThingTemplatesUpdateHandler
	// ThingTemplatesWeaviateThingTemplatesValidateHandler sets the operation handler for the weaviate thing templates validate operation
	ThingTemplatesWeaviateThingTemplatesValidateHandler thing_templates.WeaviateThingTemplatesValidateHandler
	// ThingsWeaviateThingsDeleteHandler sets the operation handler for the weaviate things delete operation
	ThingsWeaviateThingsDeleteHandler things.WeaviateThingsDeleteHandler
	// EventsWeaviateThingsEventsCreateHandler sets the operation handler for the weaviate things events create operation
	EventsWeaviateThingsEventsCreateHandler events.WeaviateThingsEventsCreateHandler
	// EventsWeaviateThingsEventsGetHandler sets the operation handler for the weaviate things events get operation
	EventsWeaviateThingsEventsGetHandler events.WeaviateThingsEventsGetHandler
	// EventsWeaviateThingsEventsListHandler sets the operation handler for the weaviate things events list operation
	EventsWeaviateThingsEventsListHandler events.WeaviateThingsEventsListHandler
	// ThingsWeaviateThingsGetHandler sets the operation handler for the weaviate things get operation
	ThingsWeaviateThingsGetHandler things.WeaviateThingsGetHandler
	// ThingsWeaviateThingsInsertHandler sets the operation handler for the weaviate things insert operation
	ThingsWeaviateThingsInsertHandler things.WeaviateThingsInsertHandler
	// ThingsWeaviateThingsListHandler sets the operation handler for the weaviate things list operation
	ThingsWeaviateThingsListHandler things.WeaviateThingsListHandler
	// ThingsWeaviateThingsPatchHandler sets the operation handler for the weaviate things patch operation
	ThingsWeaviateThingsPatchHandler things.WeaviateThingsPatchHandler
	// ThingsWeaviateThingsUpdateHandler sets the operation handler for the weaviate things update operation
	ThingsWeaviateThingsUpdateHandler things.WeaviateThingsUpdateHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *WeaviateAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *WeaviateAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *WeaviateAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *WeaviateAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *WeaviateAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *WeaviateAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *WeaviateAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the WeaviateAPI
func (o *WeaviateAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.BinConsumer == nil {
		unregistered = append(unregistered, "BinConsumer")
	}

	if o.UrlformConsumer == nil {
		unregistered = append(unregistered, "UrlformConsumer")
	}

	if o.YamlConsumer == nil {
		unregistered = append(unregistered, "YamlConsumer")
	}

	if o.XMLConsumer == nil {
		unregistered = append(unregistered, "XMLConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.TxtConsumer == nil {
		unregistered = append(unregistered, "TxtConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}

	if o.UrlformProducer == nil {
		unregistered = append(unregistered, "UrlformProducer")
	}

	if o.YamlProducer == nil {
		unregistered = append(unregistered, "YamlProducer")
	}

	if o.XMLProducer == nil {
		unregistered = append(unregistered, "XMLProducer")
	}

	if o.MultipartformProducer == nil {
		unregistered = append(unregistered, "MultipartformProducer")
	}

	if o.TxtProducer == nil {
		unregistered = append(unregistered, "TxtProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "XAPIKEYAuth")
	}

	if o.CommandsWeaviateCommandsDeleteHandler == nil {
		unregistered = append(unregistered, "commands.WeaviateCommandsDeleteHandler")
	}

	if o.CommandsWeaviateCommandsGetHandler == nil {
		unregistered = append(unregistered, "commands.WeaviateCommandsGetHandler")
	}

	if o.CommandsWeaviateCommandsInsertHandler == nil {
		unregistered = append(unregistered, "commands.WeaviateCommandsInsertHandler")
	}

	if o.CommandsWeaviateCommandsListHandler == nil {
		unregistered = append(unregistered, "commands.WeaviateCommandsListHandler")
	}

	if o.CommandsWeaviateCommandsPatchHandler == nil {
		unregistered = append(unregistered, "commands.WeaviateCommandsPatchHandler")
	}

	if o.CommandsWeaviateCommandsUpdateHandler == nil {
		unregistered = append(unregistered, "commands.WeaviateCommandsUpdateHandler")
	}

	if o.CommandsWeaviateCommandsValidateHandler == nil {
		unregistered = append(unregistered, "commands.WeaviateCommandsValidateHandler")
	}

	if o.GroupsWeaviateGroupsDeleteHandler == nil {
		unregistered = append(unregistered, "groups.WeaviateGroupsDeleteHandler")
	}

	if o.EventsWeaviateGroupsEventsCreateHandler == nil {
		unregistered = append(unregistered, "events.WeaviateGroupsEventsCreateHandler")
	}

	if o.EventsWeaviateGroupsEventsGetHandler == nil {
		unregistered = append(unregistered, "events.WeaviateGroupsEventsGetHandler")
	}

	if o.EventsWeaviateGroupsEventsListHandler == nil {
		unregistered = append(unregistered, "events.WeaviateGroupsEventsListHandler")
	}

	if o.GroupsWeaviateGroupsGetHandler == nil {
		unregistered = append(unregistered, "groups.WeaviateGroupsGetHandler")
	}

	if o.GroupsWeaviateGroupsInsertHandler == nil {
		unregistered = append(unregistered, "groups.WeaviateGroupsInsertHandler")
	}

	if o.GroupsWeaviateGroupsListHandler == nil {
		unregistered = append(unregistered, "groups.WeaviateGroupsListHandler")
	}

	if o.GroupsWeaviateGroupsPatchHandler == nil {
		unregistered = append(unregistered, "groups.WeaviateGroupsPatchHandler")
	}

	if o.GroupsWeaviateGroupsUpdateHandler == nil {
		unregistered = append(unregistered, "groups.WeaviateGroupsUpdateHandler")
	}

	if o.KeysWeaviateKeyCreateHandler == nil {
		unregistered = append(unregistered, "keys.WeaviateKeyCreateHandler")
	}

	if o.KeysWeaviateKeysChildrenGetHandler == nil {
		unregistered = append(unregistered, "keys.WeaviateKeysChildrenGetHandler")
	}

	if o.KeysWeaviateKeysDeleteHandler == nil {
		unregistered = append(unregistered, "keys.WeaviateKeysDeleteHandler")
	}

	if o.KeysWeaviateKeysGetHandler == nil {
		unregistered = append(unregistered, "keys.WeaviateKeysGetHandler")
	}

	if o.LocationsWeaviateLocationsDeleteHandler == nil {
		unregistered = append(unregistered, "locations.WeaviateLocationsDeleteHandler")
	}

	if o.LocationsWeaviateLocationsGetHandler == nil {
		unregistered = append(unregistered, "locations.WeaviateLocationsGetHandler")
	}

	if o.LocationsWeaviateLocationsInsertHandler == nil {
		unregistered = append(unregistered, "locations.WeaviateLocationsInsertHandler")
	}

	if o.LocationsWeaviateLocationsListHandler == nil {
		unregistered = append(unregistered, "locations.WeaviateLocationsListHandler")
	}

	if o.LocationsWeaviateLocationsPatchHandler == nil {
		unregistered = append(unregistered, "locations.WeaviateLocationsPatchHandler")
	}

	if o.LocationsWeaviateLocationsUpdateHandler == nil {
		unregistered = append(unregistered, "locations.WeaviateLocationsUpdateHandler")
	}

	if o.ThingTemplatesWeaviateThingTemplatesCreateHandler == nil {
		unregistered = append(unregistered, "thing_templates.WeaviateThingTemplatesCreateHandler")
	}

	if o.ThingTemplatesWeaviateThingTemplatesDeleteHandler == nil {
		unregistered = append(unregistered, "thing_templates.WeaviateThingTemplatesDeleteHandler")
	}

	if o.ThingTemplatesWeaviateThingTemplatesGetHandler == nil {
		unregistered = append(unregistered, "thing_templates.WeaviateThingTemplatesGetHandler")
	}

	if o.ThingTemplatesWeaviateThingTemplatesListHandler == nil {
		unregistered = append(unregistered, "thing_templates.WeaviateThingTemplatesListHandler")
	}

	if o.ThingTemplatesWeaviateThingTemplatesPatchHandler == nil {
		unregistered = append(unregistered, "thing_templates.WeaviateThingTemplatesPatchHandler")
	}

	if o.ThingTemplatesWeaviateThingTemplatesUpdateHandler == nil {
		unregistered = append(unregistered, "thing_templates.WeaviateThingTemplatesUpdateHandler")
	}

	if o.ThingTemplatesWeaviateThingTemplatesValidateHandler == nil {
		unregistered = append(unregistered, "thing_templates.WeaviateThingTemplatesValidateHandler")
	}

	if o.ThingsWeaviateThingsDeleteHandler == nil {
		unregistered = append(unregistered, "things.WeaviateThingsDeleteHandler")
	}

	if o.EventsWeaviateThingsEventsCreateHandler == nil {
		unregistered = append(unregistered, "events.WeaviateThingsEventsCreateHandler")
	}

	if o.EventsWeaviateThingsEventsGetHandler == nil {
		unregistered = append(unregistered, "events.WeaviateThingsEventsGetHandler")
	}

	if o.EventsWeaviateThingsEventsListHandler == nil {
		unregistered = append(unregistered, "events.WeaviateThingsEventsListHandler")
	}

	if o.ThingsWeaviateThingsGetHandler == nil {
		unregistered = append(unregistered, "things.WeaviateThingsGetHandler")
	}

	if o.ThingsWeaviateThingsInsertHandler == nil {
		unregistered = append(unregistered, "things.WeaviateThingsInsertHandler")
	}

	if o.ThingsWeaviateThingsListHandler == nil {
		unregistered = append(unregistered, "things.WeaviateThingsListHandler")
	}

	if o.ThingsWeaviateThingsPatchHandler == nil {
		unregistered = append(unregistered, "things.WeaviateThingsPatchHandler")
	}

	if o.ThingsWeaviateThingsUpdateHandler == nil {
		unregistered = append(unregistered, "things.WeaviateThingsUpdateHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *WeaviateAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *WeaviateAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "apiKey":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, o.APIKeyAuth)

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *WeaviateAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "application/json-patch+json":
			result["application/json-patch+json"] = o.JSONConsumer

		case "application/octet-stream":
			result["application/octet-stream"] = o.BinConsumer

		case "application/x-www-form-urlencoded":
			result["application/x-www-form-urlencoded"] = o.UrlformConsumer

		case "application/x-yaml":
			result["application/x-yaml"] = o.YamlConsumer

		case "application/xml":
			result["application/xml"] = o.XMLConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		case "text/plain":
			result["text/plain"] = o.TxtConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *WeaviateAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer

		case "application/x-www-form-urlencoded":
			result["application/x-www-form-urlencoded"] = o.UrlformProducer

		case "application/x-yaml":
			result["application/x-yaml"] = o.YamlProducer

		case "application/xml":
			result["application/xml"] = o.XMLProducer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformProducer

		case "text/plain":
			result["text/plain"] = o.TxtProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *WeaviateAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the weaviate API
func (o *WeaviateAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *WeaviateAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/commands/{commandId}"] = commands.NewWeaviateCommandsDelete(o.context, o.CommandsWeaviateCommandsDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/commands/{commandId}"] = commands.NewWeaviateCommandsGet(o.context, o.CommandsWeaviateCommandsGetHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/commands"] = commands.NewWeaviateCommandsInsert(o.context, o.CommandsWeaviateCommandsInsertHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/commands"] = commands.NewWeaviateCommandsList(o.context, o.CommandsWeaviateCommandsListHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/commands/{commandId}"] = commands.NewWeaviateCommandsPatch(o.context, o.CommandsWeaviateCommandsPatchHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/commands/{commandId}"] = commands.NewWeaviateCommandsUpdate(o.context, o.CommandsWeaviateCommandsUpdateHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/commands/validate"] = commands.NewWeaviateCommandsValidate(o.context, o.CommandsWeaviateCommandsValidateHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/groups/{groupId}"] = groups.NewWeaviateGroupsDelete(o.context, o.GroupsWeaviateGroupsDeleteHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/groups/{groupId}/events"] = events.NewWeaviateGroupsEventsCreate(o.context, o.EventsWeaviateGroupsEventsCreateHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/groups/{groupId}/events/{eventId}"] = events.NewWeaviateGroupsEventsGet(o.context, o.EventsWeaviateGroupsEventsGetHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/groups/{groupId}/events"] = events.NewWeaviateGroupsEventsList(o.context, o.EventsWeaviateGroupsEventsListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/groups/{groupId}"] = groups.NewWeaviateGroupsGet(o.context, o.GroupsWeaviateGroupsGetHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/groups"] = groups.NewWeaviateGroupsInsert(o.context, o.GroupsWeaviateGroupsInsertHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/groups"] = groups.NewWeaviateGroupsList(o.context, o.GroupsWeaviateGroupsListHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/groups/{groupId}"] = groups.NewWeaviateGroupsPatch(o.context, o.GroupsWeaviateGroupsPatchHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/groups/{groupId}"] = groups.NewWeaviateGroupsUpdate(o.context, o.GroupsWeaviateGroupsUpdateHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/keys"] = keys.NewWeaviateKeyCreate(o.context, o.KeysWeaviateKeyCreateHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/keys/{keyId}/children"] = keys.NewWeaviateKeysChildrenGet(o.context, o.KeysWeaviateKeysChildrenGetHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/keys/{keyId}"] = keys.NewWeaviateKeysDelete(o.context, o.KeysWeaviateKeysDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/keys/{keyId}"] = keys.NewWeaviateKeysGet(o.context, o.KeysWeaviateKeysGetHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/locations/{locationId}"] = locations.NewWeaviateLocationsDelete(o.context, o.LocationsWeaviateLocationsDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/locations/{locationId}"] = locations.NewWeaviateLocationsGet(o.context, o.LocationsWeaviateLocationsGetHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/locations"] = locations.NewWeaviateLocationsInsert(o.context, o.LocationsWeaviateLocationsInsertHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/locations"] = locations.NewWeaviateLocationsList(o.context, o.LocationsWeaviateLocationsListHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/locations/{locationId}"] = locations.NewWeaviateLocationsPatch(o.context, o.LocationsWeaviateLocationsPatchHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/locations/{locationId}"] = locations.NewWeaviateLocationsUpdate(o.context, o.LocationsWeaviateLocationsUpdateHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/thingTemplates"] = thing_templates.NewWeaviateThingTemplatesCreate(o.context, o.ThingTemplatesWeaviateThingTemplatesCreateHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/thingTemplates/{thingTemplateId}"] = thing_templates.NewWeaviateThingTemplatesDelete(o.context, o.ThingTemplatesWeaviateThingTemplatesDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/thingTemplates/{thingTemplateId}"] = thing_templates.NewWeaviateThingTemplatesGet(o.context, o.ThingTemplatesWeaviateThingTemplatesGetHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/thingTemplates"] = thing_templates.NewWeaviateThingTemplatesList(o.context, o.ThingTemplatesWeaviateThingTemplatesListHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/thingTemplates/{thingTemplateId}"] = thing_templates.NewWeaviateThingTemplatesPatch(o.context, o.ThingTemplatesWeaviateThingTemplatesPatchHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/thingTemplates/{thingTemplateId}"] = thing_templates.NewWeaviateThingTemplatesUpdate(o.context, o.ThingTemplatesWeaviateThingTemplatesUpdateHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/thingTemplates/validate"] = thing_templates.NewWeaviateThingTemplatesValidate(o.context, o.ThingTemplatesWeaviateThingTemplatesValidateHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/things/{thingId}"] = things.NewWeaviateThingsDelete(o.context, o.ThingsWeaviateThingsDeleteHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/things/{thingId}/events"] = events.NewWeaviateThingsEventsCreate(o.context, o.EventsWeaviateThingsEventsCreateHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/things/{thingId}/events/{eventId}"] = events.NewWeaviateThingsEventsGet(o.context, o.EventsWeaviateThingsEventsGetHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/things/{thingId}/events"] = events.NewWeaviateThingsEventsList(o.context, o.EventsWeaviateThingsEventsListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/things/{thingId}"] = things.NewWeaviateThingsGet(o.context, o.ThingsWeaviateThingsGetHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/things"] = things.NewWeaviateThingsInsert(o.context, o.ThingsWeaviateThingsInsertHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/things"] = things.NewWeaviateThingsList(o.context, o.ThingsWeaviateThingsListHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/things/{thingId}"] = things.NewWeaviateThingsPatch(o.context, o.ThingsWeaviateThingsPatchHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/things/{thingId}"] = things.NewWeaviateThingsUpdate(o.context, o.ThingsWeaviateThingsUpdateHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *WeaviateAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *WeaviateAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
