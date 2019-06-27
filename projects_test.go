// Copyright (c) 2019, NewReleases Go client AUTHORS.
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
	"strconv"
	"testing"

	"newreleases.io/newreleases"
)

func TestProjectsService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects", requireMethod("GET", newPagedStaticHndler(projectsServiceList...)))

	for i, page := range projectsServiceListWant {
		name := "page " + strconv.Itoa(i+1)

		got, lastPage, err := client.Projects.List(context.Background(), newreleases.ProjectListOptions{
			Page: i + 1,
		})
		if err != nil {
			t.Fatal(name, err)
		}

		assertEqual(t, name, got, page)
		assertEqual(t, name, lastPage, 2)
	}
}

func TestProjectsService_List_order(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("order") != "name" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, struct{}{})
	}))

	got, _, err := client.Projects.List(context.Background(), newreleases.ProjectListOptions{
		Page:  1,
		Order: newreleases.ProjectListOrderName,
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, []newreleases.Project(nil))
}

func TestProjectsService_List_reverse(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		if _, reverse := r.URL.Query()["reverse"]; !reverse {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, struct{}{})
	}))

	got, _, err := client.Projects.List(context.Background(), newreleases.ProjectListOptions{
		Page:    1,
		Reverse: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, []newreleases.Project(nil))
}

func TestProjectsService_List_provider(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/github", requireMethod("GET", newPagedStaticHndler(projectsServiceList...)))

	for i, page := range projectsServiceListWant {
		name := "page " + strconv.Itoa(i+1)

		got, lastPage, err := client.Projects.List(context.Background(), newreleases.ProjectListOptions{
			Page:     i + 1,
			Provider: "github",
		})
		if err != nil {
			t.Fatal(name, err)
		}

		assertEqual(t, name, got, page)
		assertEqual(t, name, lastPage, 2)
	}
}

func TestProjectsService_List_providerOrder(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("order") != "name" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, struct{}{})
	}))

	got, _, err := client.Projects.List(context.Background(), newreleases.ProjectListOptions{
		Page:  1,
		Order: newreleases.ProjectListOrderName,
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, []newreleases.Project(nil))
}

func TestProjectsService_List_providerReverse(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/github", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		if _, reverse := r.URL.Query()["reverse"]; !reverse {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, struct{}{})
	}))

	got, _, err := client.Projects.List(context.Background(), newreleases.ProjectListOptions{
		Page:     1,
		Reverse:  true,
		Provider: "github",
	})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, []newreleases.Project(nil))
}

var (
	projectsServiceList = []string{`
	{
		"projects": [
			{
				"id": "pf4w494lbjsd3ydp5hnf4gsptw",
				"name": "golang/go",
				"provider": "github",
				"url": "https://github.com/golang/go/releases",
				"email_notification": "hourly",
				"exclude_updated": true
			}
		],
		"total_pages": 2
	}
	`, `
	{
		"projects": [
			{
				"id": "dp5hnf4gsptwpf4w494lbjsd3y",
				"name": "django/django",
				"provider": "github",
				"url": "https://github.com/django/django/releases",
				"email_notification": "daily"
			}
		],
		"total_pages": 2
	}
	`,
	}
	projectsServiceListWant = [][]newreleases.Project{
		{
			{
				ID:                "pf4w494lbjsd3ydp5hnf4gsptw",
				Name:              "golang/go",
				Provider:          "github",
				URL:               "https://github.com/golang/go/releases",
				EmailNotification: "hourly",
				ExcludeUpdated:    true,
			},
		},
		{
			{
				ID:                "dp5hnf4gsptwpf4w494lbjsd3y",
				Name:              "django/django",
				Provider:          "github",
				URL:               "https://github.com/django/django/releases",
				EmailNotification: "daily",
			},
		},
	}
)

func TestProjectsService_Search(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/search", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("q") != "go" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.URL.Query().Get("provider") != "github" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, projectsServiceSearch)
	}))

	got, err := client.Projects.Search(context.Background(), "go", "github")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, projectsServiceSearchWant)
}

var (
	projectsServiceSearch = `
	{
		"projects": [
			{
				"id": "pf4w494lbjsd3ydp5hnf4gsptw",
				"name": "golang/go",
				"provider": "github",
				"url": "https://github.com/golang/go/releases",
				"email_notification": "hourly",
				"exclude_updated": true
			}
		]
	}
	`
	projectsServiceSearchWant = []newreleases.Project{
		{
			ID:                "pf4w494lbjsd3ydp5hnf4gsptw",
			Name:              "golang/go",
			Provider:          "github",
			URL:               "https://github.com/golang/go/releases",
			EmailNotification: newreleases.EmailNotificationHourly,
			ExcludeUpdated:    true,
		},
	}
)

func TestProjectsService_GetByID(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/pf4w494lbjsd3ydp5hnf4gsptw", requireMethod("GET", newStaticHandler(projectsServiceGet)))

	got, err := client.Projects.GetByID(context.Background(), "pf4w494lbjsd3ydp5hnf4gsptw")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, projectsServiceGetWant)
}

func TestProjectsService_GetByName(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/github/golang/go", requireMethod("GET", newStaticHandler(projectsServiceGet)))

	got, err := client.Projects.GetByName(context.Background(), "github", "golang/go")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, projectsServiceGetWant)
}

var (
	projectsServiceGet = `
	{
		"id": "pf4w494lbjsd3ydp5hnf4gsptw",
		"name": "golang/go",
		"provider": "github",
		"url": "https://github.com/golang/go/releases",
		"email_notification": "hourly",
		"exclude_updated": true
	}
	`
	projectsServiceGetWant = &newreleases.Project{
		ID:                "pf4w494lbjsd3ydp5hnf4gsptw",
		Name:              "golang/go",
		Provider:          "github",
		URL:               "https://github.com/golang/go/releases",
		EmailNotification: newreleases.EmailNotificationHourly,
		ExcludeUpdated:    true,
	}
)

func TestProjectsService_Add(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects", requireMethod("POST", func(w http.ResponseWriter, r *http.Request) {
		var o *newreleases.ProjectOptions
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(projectOptions, o) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		d, err := json.Marshal(projectWant)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", jsonContentType)
		w.Write(d)
	}))

	got, err := client.Projects.Add(context.Background(), "github", "golang/go", projectOptions)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, projectWant)
}

func TestProjectsService_UpdateByName(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/github/golang/go", requireMethod("POST", func(w http.ResponseWriter, r *http.Request) {
		var o *newreleases.ProjectOptions
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(projectOptions, o) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		d, err := json.Marshal(projectWant)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", jsonContentType)
		w.Write(d)
	}))

	got, err := client.Projects.UpdateByName(context.Background(), "github", "golang/go", projectOptions)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, projectWant)
}

func TestProjectsService_UpdateByID(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/pf4w494lbjsd3ydp5hnf4gsptw", requireMethod("POST", func(w http.ResponseWriter, r *http.Request) {
		var o *newreleases.ProjectOptions
		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(projectOptions, o) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		d, err := json.Marshal(projectWant)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", jsonContentType)
		w.Write(d)
	}))

	got, err := client.Projects.UpdateByID(context.Background(), "pf4w494lbjsd3ydp5hnf4gsptw", projectOptions)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, projectWant)
}

var (
	projectWant = &newreleases.Project{
		Name:                   "golang/go",
		Provider:               "github",
		EmailNotification:      newreleases.EmailNotificationHourly,
		SlackIDs:               []string{"slack123"},
		TelegramChatIDs:        []string{"telegram123"},
		HangoutsChatWebhookIDs: []string{"hangouts123"},
		MSTeamsWebhookIDs:      []string{"teams123"},
		WebhookIDs:             []string{"webhook123", "webhook124"},
		Exclusions:             []newreleases.Exclusion{{Value: "^1.9"}},
		ExcludePrereleases:     true,
		ExcludeUpdated:         false,
	}
	projectOptions = &newreleases.ProjectOptions{
		EmailNotification:      &newreleases.EmailNotificationHourly,
		SlackIDs:               []string{"slack123"},
		TelegramChatIDs:        []string{"telegram123"},
		HangoutsChatWebhookIDs: []string{"hangouts123"},
		MSTeamsWebhookIDs:      []string{"teams123"},
		WebhookIDs:             []string{"webhook123", "webhook124"},
		Exclusions:             []newreleases.Exclusion{{Value: "^1.9"}},
		ExcludePrereleases:     newreleases.Bool(true),
		ExcludeUpdated:         newreleases.Bool(false),
	}
)

func TestProjectsService_DeleteByID(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/pf4w494lbjsd3ydp5hnf4gsptw", requireMethod("DELETE", func(w http.ResponseWriter, r *http.Request) {}))

	err := client.Projects.DeleteByID(context.Background(), "pf4w494lbjsd3ydp5hnf4gsptw")
	if err != nil {
		t.Fatal(err)
	}
}

func TestProjectsService_DeleteByName(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/github/golang/go", requireMethod("DELETE", func(w http.ResponseWriter, r *http.Request) {}))

	err := client.Projects.DeleteByName(context.Background(), "github", "golang/go")
	if err != nil {
		t.Fatal(err)
	}
}
