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
 package models




import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CommandProgress Progress of the command. Will also be included in the events.
// swagger:model CommandProgress
type CommandProgress string

const (
	// CommandProgressAborted captures enum value "aborted"
	CommandProgressAborted CommandProgress = "aborted"
	// CommandProgressCancelled captures enum value "cancelled"
	CommandProgressCancelled CommandProgress = "cancelled"
	// CommandProgressDone captures enum value "done"
	CommandProgressDone CommandProgress = "done"
	// CommandProgressError captures enum value "error"
	CommandProgressError CommandProgress = "error"
	// CommandProgressExpired captures enum value "expired"
	CommandProgressExpired CommandProgress = "expired"
	// CommandProgressInProgress captures enum value "inProgress"
	CommandProgressInProgress CommandProgress = "inProgress"
	// CommandProgressQueued captures enum value "queued"
	CommandProgressQueued CommandProgress = "queued"
)

// for schema
var commandProgressEnum []interface{}

func init() {
	var res []CommandProgress
	if err := json.Unmarshal([]byte(`["aborted","cancelled","done","error","expired","inProgress","queued"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commandProgressEnum = append(commandProgressEnum, v)
	}
}

func (m CommandProgress) validateCommandProgressEnum(path, location string, value CommandProgress) error {
	if err := validate.Enum(path, location, value, commandProgressEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this command progress
func (m CommandProgress) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCommandProgressEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
