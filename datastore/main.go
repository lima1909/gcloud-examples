package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

const (
	projectID = "goheros-207118"
	kind      = "person"
)

// 1. install emulator: https://cloud.google.com/datastore/docs/tools/datastore-emulator
// 2. example: https://github.com/googleapis/google-cloud-go/blob/master/datastore/datastore.go
// 3. start emulator: gcloud beta emulators datastore start
// 4. export Env: export DATASTORE_EMULATOR_HOST=localhost:8081
func main() {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}

	type Person struct {
		Name  string
		Alter int
	}
	p := Person{"blub", 3}

	key := datastore.IncompleteKey(kind, nil)
	key, err = client.Put(ctx, key, &p)
	if err != nil {
		panic("Failed to save quote: " + err.Error())
	}

	query := datastore.NewQuery(kind)
	it := client.Run(ctx, query)
	for {
		var per Person
		k, err := it.Next(&per)
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic("Error fetching next data: " + err.Error())
		}
		fmt.Printf("Key: %v -- %#v\n", *k, per)
	}
}
