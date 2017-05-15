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

package datastore

import (
	"context"
	"log"

	gouuid "github.com/satori/go.uuid"
	"github.com/weaviate/weaviate/connectors"

	"fmt"

	"errors"

	"cloud.google.com/go/datastore"
)

// Datastore has some basic variables.
type Datastore struct {
	client *datastore.Client
	kind   string
}

// Connect to datastore
func (f *Datastore) Connect() error {
	// Set ctx and your Google Cloud Platform project ID.
	ctx := context.Background()

	f.kind = "weaviate"

	projectID := "weaviate-dev-001"

	client, err := datastore.NewClient(ctx, projectID)

	if err != nil {
		return err
	}

	f.client = client
	return nil
}

// Add item to DB
func (f *Datastore) Add(dbObject dbconnector.DatabaseObject) (string, error) {
	ctx := context.Background()

	f.kind = "weaviate"

	nameUUID := fmt.Sprintf("%v", gouuid.NewV4())

	// Creates a Key instance.
	taskKey := datastore.NameKey(f.kind, nameUUID, nil)

	// Saves the new entity.
	if _, err := f.client.Put(ctx, taskKey, &dbObject); err != nil {
		log.Fatalf("Failed to save task: %v", err)
		return "Error", err
	}

	// Return the ID that is used to create.
	return dbObject.Uuid, nil
}

// Get DatabaseObject from DB by uuid
func (f *Datastore) Get(uuid string) (dbconnector.DatabaseObject, error) {
	ctx := context.Background()

	f.kind = "weaviate"

	query := datastore.NewQuery(f.kind).Filter("Uuid =", uuid).Order("-CreateTimeMs").Limit(1)

	object := []dbconnector.DatabaseObject{}

	keys, err := f.client.GetAll(ctx, query, &object)

	if err != nil {
		log.Fatalf("Failed to load task: %v", err)

		return dbconnector.DatabaseObject{}, err
	}

	if len(keys) == 0 {
		notFoundErr := errors.New("no object with such UUID found")
		return dbconnector.DatabaseObject{}, notFoundErr
	}

	return object[0], nil
}

// List lists the items from Datastore by refType and limit
func (f *Datastore) List(refType string, limit int) ([]dbconnector.DatabaseObject, error) {
	ctx := context.Background()

	f.kind = "weaviate"

	query := datastore.NewQuery(f.kind).Filter("RefType =", refType).Filter("Deleted =", false).Order("-CreateTimeMs").Limit(limit)

	dbObjects := []dbconnector.DatabaseObject{}

	keys, err := f.client.GetAll(ctx, query, &dbObjects)

	if err != nil {
		log.Fatalf("Failed to load task: %v", err)

		return []dbconnector.DatabaseObject{}, err
	}

	if len(keys) == 0 {
		return []dbconnector.DatabaseObject{}, nil
	}

	return dbObjects, nil
}

// Validate if a user has access, returns permissions object
func (f *Datastore) ValidateUser(token string) []dbconnector.DatabaseUsersObject {

	ctx := context.Background()

	f.kind = "weaviate_users"

	query := datastore.NewQuery(f.kind).Filter("KeyToken =", token).Limit(1)

	dbUsersObjects := []dbconnector.DatabaseUsersObject{}

	_, err := f.client.GetAll(ctx, query, &dbUsersObjects)

	if err != nil {
		log.Fatalf("Failed to load task: %v", err)
	}

	// keys are found, return true
	return dbUsersObjects
}