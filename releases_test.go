// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"newreleases.io/newreleases"
)

func TestReleasesService_ListByProjectID(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/8wdvh4w9bhsvzclz4ynaqpcpvg/releases", requireMethod("GET", newPagedStaticHandler(releasesServiceList...)))

	for i, page := range releasesServiceListWant {
		name := "page " + strconv.Itoa(i+1)

		got, lastPage, err := client.Releases.ListByProjectID(context.Background(), "8wdvh4w9bhsvzclz4ynaqpcpvg", i+1)
		if err != nil {
			t.Fatal(name, err)
		}

		assertEqual(t, name, got, page)
		assertEqual(t, name, lastPage, 2)
	}
}

func TestReleasesService_ListByProjectName(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/github/nodejs/node/releases", requireMethod("GET", newPagedStaticHandler(releasesServiceList...)))

	for i, page := range releasesServiceListWant {
		name := "page " + strconv.Itoa(i+1)

		got, lastPage, err := client.Releases.ListByProjectName(context.Background(), "github", "nodejs/node", i+1)
		if err != nil {
			t.Fatal(name, err)
		}

		assertEqual(t, name, got, page)
		assertEqual(t, name, lastPage, 2)
	}
}

var (
	releasesServiceList = []string{`
	{
		"releases": [
			{
				"version": "v11.12.0",
				"date": "2019-03-15T21:16:54Z",
				"is_prerelease": true
			},
			{
				"version": "v11.11.0",
				"date": "2019-03-06T19:44:17Z",
				"is_updated": true
			},
			{
				"version": "v10.15.3",
				"date": "2019-03-05T17:37:13Z",
				"CVE": ["CVE-2020-11080"],
				"has_note": true
			}
		],
		"total_pages": 2
	}
	`, `
	{
		"releases": [
			{
				"version": "v6.17.0",
				"date": "2019-02-28T11:23:55Z",
				"is_excluded": true
			},
			{
				"version": "v8.15.1",
				"date": "2019-02-28T11:20:35Z"
			}
		],
		"total_pages": 2
	}
	`,
	}
	releasesServiceListWant = [][]newreleases.Release{
		{
			{
				Version:      "v11.12.0",
				Date:         parseTime("2019-03-15T21:16:54Z"),
				IsPrerelease: true,
			},
			{
				Version:   "v11.11.0",
				Date:      parseTime("2019-03-06T19:44:17Z"),
				IsUpdated: true,
			},
			{
				Version: "v10.15.3",
				Date:    parseTime("2019-03-05T17:37:13Z"),
				CVE:     []string{"CVE-2020-11080"},
				HasNote: true,
			},
		},
		{
			{
				Version:    "v6.17.0",
				Date:       parseTime("2019-02-28T11:23:55Z"),
				IsExcluded: true,
			},
			{
				Version: "v8.15.1",
				Date:    parseTime("2019-02-28T11:20:35Z"),
			},
		},
	}
)

func TestReleasesService_GetByProjectID(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/8wdvh4w9bhsvzclz4ynaqpcpvg/releases/v11.12.0", requireMethod("GET", newStaticHandler(releasesServiceGet)))

	got, err := client.Releases.GetByProjectID(context.Background(), "8wdvh4w9bhsvzclz4ynaqpcpvg", "v11.12.0")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, releasesServiceGetWant)
}

func TestReleasesService_GetByProjectName(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/github/nodejs/node/releases/v11.12.0", requireMethod("GET", newStaticHandler(releasesServiceGet)))

	got, err := client.Releases.GetByProjectName(context.Background(), "github", "nodejs/node", "v11.12.0")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, releasesServiceGetWant)
}

var (
	releasesServiceGet = `
	{
		"version": "v11.12.0",
		"date": "2019-03-15T21:16:54Z",
		"is_prerelease": true
	}
	`
	releasesServiceGetWant = &newreleases.Release{
		Version:      "v11.12.0",
		Date:         parseTime("2019-03-15T21:16:54Z"),
		IsPrerelease: true,
	}
)

func TestReleasesService_GetNoteByProjectID(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/8wdvh4w9bhsvzclz4ynaqpcpvg/releases/v11.12.0/note", requireMethod("GET", newStaticHandler(releasesServiceGetNote)))

	got, err := client.Releases.GetNoteByProjectID(context.Background(), "8wdvh4w9bhsvzclz4ynaqpcpvg", "v11.12.0")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, releasesServiceGetNoteWant)
}

func TestReleasesService_GetNoteByProjectName(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/projects/github/nodejs/node/releases/v11.12.0/note", requireMethod("GET", newStaticHandler(releasesServiceGetNote)))

	got, err := client.Releases.GetNoteByProjectName(context.Background(), "github", "nodejs/node", "v11.12.0")
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, releasesServiceGetNoteWant)
}

var (
	releasesServiceGetNote = `
	{
		"title": "2019-03-15, Version 11.12.0 (Current), @BridgeAR",
		"message": "<h3>Notable Changes</h3>\n\n",
		"url": "https://github.com/nodejs/node/releases/tag/v11.12.0"
	}
	`
	releasesServiceGetNoteWant = &newreleases.ReleaseNote{
		Title:   "2019-03-15, Version 11.12.0 (Current), @BridgeAR",
		Message: "<h3>Notable Changes</h3>\n\n",
		URL:     "https://github.com/nodejs/node/releases/tag/v11.12.0",
	}
)

func parseTime(s string) (t time.Time) {
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		panic(err)
	}
	return t
}
