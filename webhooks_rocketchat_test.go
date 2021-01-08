// Copyright (c) 2021, NewReleases Go client AUTHORS.
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

func TestRocketchatWebhooksService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/rocketchat-webhooks", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, rocketchatWebhooksServiceList)
	}))

	got, err := client.RocketchatWebhooks.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, rocketchatWebhooksServiceListWant)
}

var (
	rocketchatWebhooksServiceList = `
	{
		"webhooks": [
			{
				"id": "lst4ezma884w1a56q4es0dvttd",
				"name": "My Channel"
			}
		]
	}
	`
	rocketchatWebhooksServiceListWant = []newreleases.Webhook{
		{
			ID:   "lst4ezma884w1a56q4es0dvttd",
			Name: "My Channel",
		},
	}
)
