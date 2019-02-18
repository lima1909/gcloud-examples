package main

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const (
	// NAMESPACE ...
	NAMESPACE = "data"
	kind      = "person"
)

// Person ...
type Person struct {
	Name  string
	Alter int
}

// dev_appserver.py --enable_console --port=8082 app.yaml
// oder
// dev_appserver.py --enable_console --clear_datastore  --port=8082 app.yaml
//
// Aufruf: http://localhost:8082/
// :+1:
func main() {

	http.HandleFunc("/", list)
	http.HandleFunc("/add", add)
	appengine.Main()

}

func list(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ctx = setNamespace(ctx)

	var ps []Person
	_, err := datastore.NewQuery(kind).Filter("Alter =", 3).GetAll(ctx, &ps)
	if err != nil {
		panic(err)
	}

	log.Infof(ctx, "--> %v", ps)
	fmt.Fprint(w, ps)
}

func add(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ctx = setNamespace(ctx)

	k := datastore.NewIncompleteKey(ctx, kind, nil)
	p := Person{"blub", 3}
	_, err := datastore.Put(ctx, k, &p)
	if err != nil {
		log.Errorf(ctx, "Err by datastore.Put: %v", err)
	}
}

func setNamespace(c context.Context) context.Context {
	c, err := appengine.Namespace(c, NAMESPACE)
	if err != nil {
		log.Errorf(c, fmt.Sprintf("Err by set Namespace: %v", err))
	}
	return c
}
