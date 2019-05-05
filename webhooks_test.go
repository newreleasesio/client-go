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

	"newreleases.io/newreleases"
)

func TestWebhooksService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/webhooks", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, webhooksServiceList)
	}))

	got, err := client.Webhooks.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, webhooksServiceListWant)
}

var (
	webhooksServiceList = `
	{
		"webhooks": [
			{
				"id": "aw680q7n6vv2s75snp2lpr006m",
				"name": "My Webhook"
			}
		]
	}
	`
	webhooksServiceListWant = []newreleases.Webhook{
		{
			ID:   "aw680q7n6vv2s75snp2lpr006m",
			Name: "My Webhook",
		},
	}
)
