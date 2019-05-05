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

func TestHangoutsChatWebhooksService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/hangouts-chat-webhooks", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, hangoutsChatWebhooksServiceList)
	}))

	got, err := client.HangoutsChatWebhooks.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, hangoutsChatWebhooksServiceListWant)
}

var (
	hangoutsChatWebhooksServiceList = `
	{
		"webhooks": [
			{
				"id": "ewcppcsk781h1bwlp0q3pe8gf7",
				"name": "My Hangouts"
			}
		]
	}
	`
	hangoutsChatWebhooksServiceListWant = []newreleases.Webhook{
		{
			ID:   "ewcppcsk781h1bwlp0q3pe8gf7",
			Name: "My Hangouts",
		},
	}
)
