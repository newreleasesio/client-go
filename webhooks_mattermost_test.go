// Copyright (c) 2020, NewReleases Go client AUTHORS.
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

func TestMattermostWebhooksService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/mattermost-webhooks", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, mattermostWebhooksServiceList)
	}))

	got, err := client.MattermostWebhooks.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, mattermostWebhooksServiceListWant)
}

var (
	mattermostWebhooksServiceList = `
	{
		"webhooks": [
			{
				"id": "lst4ezma884w1a56q4es0dvttd",
				"name": "My Channel"
			}
		]
	}
	`
	mattermostWebhooksServiceListWant = []newreleases.Webhook{
		{
			ID:   "lst4ezma884w1a56q4es0dvttd",
			Name: "My Channel",
		},
	}
)
