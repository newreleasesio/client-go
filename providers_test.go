// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestProvidersService(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/providers", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		if _, ok := r.URL.Query()["added"]; ok {
			fmt.Fprintln(w, providersServiceListAdded)
			return
		}
		fmt.Fprintln(w, providersServiceList)
	}))

	t.Run("List", func(t *testing.T) {
		got, err := client.Providers.List(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		assertEqual(t, "", got, providersServiceListWant)
	})

	t.Run("ListAdded", func(t *testing.T) {
		got, err := client.Providers.ListAdded(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		assertEqual(t, "", got, providersServiceListAddedWant)
	})
}

var (
	providersServiceList = `
	{
		"providers": [
			"bitbucket",
			"cargo",
			"dockerhub",
			"gems",
			"github",
			"gitlab",
			"maven",
			"npm",
			"packagist",
			"pypi",
			"yarn"
		]
	}
	`
	providersServiceListWant = []string{
		"bitbucket",
		"cargo",
		"dockerhub",
		"gems",
		"github",
		"gitlab",
		"maven",
		"npm",
		"packagist",
		"pypi",
		"yarn",
	}
	providersServiceListAdded = `
	{
		"providers": [
			"github",
			"npm",
			"pypi"
		]
	}
	`
	providersServiceListAddedWant = []string{
		"github",
		"npm",
		"pypi",
	}
)
