// Copyright (c) 2022, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"newreleases.io/newreleases"
)

func TestTagsService_Get(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/tags/db733f1254b9", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, `{
			"id": "db733f1254b9",
			"name": "Awesome"
		}`)
	}))

	got, err := client.Tags.Get(context.Background(), "db733f1254b9")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, &newreleases.Tag{
		ID:   "db733f1254b9",
		Name: "Awesome",
	})
}

func TestTagsService_List(t *testing.T) {

	var (
		tagsServiceList = `
		{
			"tags": [
				{
					"id": "db733f1254b9",
					"name": "Awesome"
				}
			]
		}
		`
		tagsServiceListWant = []newreleases.Tag{
			{
				ID:   "db733f1254b9",
				Name: "Awesome",
			},
		}
	)

	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/tags", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, tagsServiceList)
	}))

	got, err := client.Tags.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, tagsServiceListWant)
}

func TestTagsService_Add(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	tagOptions := &newreleases.TagOptionsRequest{
		Name: "awesome",
	}

	tagWant := &newreleases.Tag{
		ID:   "db733f1254b9",
		Name: "awesome",
	}

	mux.HandleFunc("/v1/tags", requireMethod("POST", func(w http.ResponseWriter, r *http.Request) {
		var o *newreleases.TagOptionsRequest
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(tagOptions, o) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		d, err := json.Marshal(tagWant)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", jsonContentType)
		_, _ = w.Write(d)
	}))

	got, err := client.Tags.Add(context.Background(), "awesome")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, tagWant)
}

func TestTagsService_Update(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	tagOptions := &newreleases.TagOptionsRequest{
		Name: "new name",
	}

	tagWant := &newreleases.Tag{
		ID:   "db733f1254b9",
		Name: "new name",
	}

	mux.HandleFunc("/v1/tags/db733f1254b9", requireMethod("POST", func(w http.ResponseWriter, r *http.Request) {
		var o *newreleases.TagOptionsRequest
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(tagOptions, o) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		d, err := json.Marshal(tagWant)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", jsonContentType)
		_, _ = w.Write(d)
	}))

	got, err := client.Tags.Update(context.Background(), "db733f1254b9", "new name")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, tagWant)
}

func TestTagsService_Delete(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/tags/db733f1254b9", requireMethod("DELETE", func(w http.ResponseWriter, r *http.Request) {}))

	err := client.Tags.Delete(context.Background(), "db733f1254b9")
	if err != nil {
		t.Fatal(err)
	}
}
