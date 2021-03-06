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
	"errors"
	"fmt"
	"log"
	"net"

	"cloud.google.com/go/datastore"
	gouuid "github.com/satori/go.uuid"

	"encoding/json"

	"github.com/weaviate/weaviate/connectors"
)

// Datastore has some basic variables.
type Datastore struct {
	client *datastore.Client
	kind   string
}

// Connect to datastore
func (f *Datastore) Connect() error {
	// Set ctx, your Google Cloud Platform project ID and kind.
	ctx := context.Background()
	f.kind = "weaviate"
	projectID := "weaviate-dev-001"

	// Create new client
	client, err := datastore.NewClient(ctx, projectID)

	// If error, return it. Otherwise set client.
	if err != nil {
		return err
	}

	f.client = client
	return nil
}

// Creates a root key and tables if not already available
func (f *Datastore) Init() error {

	ctx := context.Background()

	f.kind = "weaviate_users"

	// create query to check for root key
	query := datastore.NewQuery(f.kind).Filter("Parent =", "*").Limit(1)

	dbKeyObjects := []dbconnector.DatabaseUsersObject{}

	_, err := f.client.GetAll(ctx, query, &dbKeyObjects)

	if err != nil {
		panic("ERROR INITIALIZING SERVER")
	}

	// No key was found, create one
	if len(dbKeyObjects) == 0 {

		dbObject := dbconnector.DatabaseUsersObject{}

		// Create key token
		dbObject.KeyToken = fmt.Sprintf("%v", gouuid.NewV4())

		// Uuid + name
		uuid := fmt.Sprintf("%v", gouuid.NewV4())

		// Creates a Key instance.
		taskKey := datastore.NameKey(f.kind, uuid, nil)

		// Auto set the parent ID to root *
		dbObject.Parent = "*"

		// Set Uuid
		dbObject.Uuid = uuid

		// Set chmod variables
		dbObjectObject := dbconnector.DatabaseUsersObjectsObject{}
		dbObjectObject.Read = true
		dbObjectObject.Write = true
		dbObjectObject.Delete = true

		// Get ips as v6
		var ips []string
		ifaces, _ := net.Interfaces()
		for _, i := range ifaces {
			addrs, _ := i.Addrs()
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}

				ipv6 := ip.To16()
				ips = append(ips, ipv6.String())
			}
		}

		dbObjectObject.IpOrigin = ips

		// Marshall and add to object
		dbObjectObjectJson, _ := json.Marshal(dbObjectObject)
		dbObject.Object = string(dbObjectObjectJson)

		// Saves the new entity.
		if _, err := f.client.Put(ctx, taskKey, &dbObject); err != nil {
			log.Fatalf("Failed to save task: %v", err)
		}

		// Print the key
		log.Println("INFO: No root key was found, a new root key is created. More info: https://github.com/weaviate/weaviate/blob/develop/README.md#authentication")
		log.Println("INFO: Auto set allowed IPs to: ", ips)
		log.Println("ROOTKEY=" + dbObject.KeyToken)
	}

	return nil
}

// Add item to DB
func (f *Datastore) Add(dbObject dbconnector.DatabaseObject) (string, error) {
	// Set ctx and kind.
	ctx := context.Background()
	f.kind = "weaviate"

	// Generate an UUID
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
	// Set ctx and kind.
	ctx := context.Background()
	f.kind = "weaviate"

	// Make get Query
	query := datastore.NewQuery(f.kind).Filter("Uuid =", uuid).Order("-CreateTimeMs").Limit(1)

	// Fill object
	object := []dbconnector.DatabaseObject{}
	keys, err := f.client.GetAll(ctx, query, &object)

	// Return error
	if err != nil {
		log.Fatalf("Failed to load task: %v", err)
		return dbconnector.DatabaseObject{}, err
	}

	// Return error 'not found'
	if len(keys) == 0 {
		notFoundErr := errors.New("no object with such UUID found")
		return dbconnector.DatabaseObject{}, notFoundErr
	}

	// Return found object
	return object[0], nil
}

// List lists the items from Datastore by refType and limit
func (f *Datastore) List(refType string, limit int) ([]dbconnector.DatabaseObject, error) {
	// Set ctx and kind.
	ctx := context.Background()
	f.kind = "weaviate"

	// Make list query
	query := datastore.NewQuery(f.kind).Filter("RefType =", refType).Filter("Deleted =", false).Order("-CreateTimeMs").Limit(limit)

	// Fill object with results
	dbObjects := []dbconnector.DatabaseObject{}
	_, err := f.client.GetAll(ctx, query, &dbObjects)

	// Return error and empty object
	if err != nil {
		log.Fatalf("Failed to load task: %v", err)

		return []dbconnector.DatabaseObject{}, err
	}

	// Return list with objects
	return dbObjects, nil
}

// Validate if a user has access, returns permissions object
func (f *Datastore) ValidateKey(token string) ([]dbconnector.DatabaseUsersObject, error) {

	ctx := context.Background()

	f.kind = "weaviate_users"

	query := datastore.NewQuery(f.kind).Filter("KeyToken =", token).Limit(1)

	dbUsersObjects := []dbconnector.DatabaseUsersObject{}

	_, err := f.client.GetAll(ctx, query, &dbUsersObjects)

	if err != nil {
		return dbUsersObjects, err
	}

	// keys are found, return true
	return dbUsersObjects, nil
}

// AddUser to DB
func (f *Datastore) AddKey(parentUuid string, dbObject dbconnector.DatabaseUsersObject) (dbconnector.DatabaseUsersObject, error) {
	ctx := context.Background()

	f.kind = "weaviate_users"

	nameUUID := fmt.Sprintf("%v", gouuid.NewV4())

	// Create key token
	dbObject.KeyToken = fmt.Sprintf("%v", gouuid.NewV4())

	// Creates a Key instance.
	taskKey := datastore.NameKey(f.kind, nameUUID, nil)

	// Auto set the parent ID
	dbObject.Parent = parentUuid

	// Saves the new entity.
	if _, err := f.client.Put(ctx, taskKey, &dbObject); err != nil {
		log.Fatalf("Failed to save task: %v", err)
		return dbObject, err
	}

	// Return the ID that is used to create.
	return dbObject, nil
}
