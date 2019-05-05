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

func TestMicrosoftTeamsWebhooksService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/microsoft-teams-webhooks", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, microsoftTeamsWebhooksServiceList)
	}))

	got, err := client.MicrosoftTeamsWebhooks.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, microsoftTeamsWebhooksServiceListWant)
}

var (
	microsoftTeamsWebhooksServiceList = `
	{
		"webhooks": [
			{
				"id": "w1a56q4es0dvtl84tdst4ezma8",
				"name": "My Team"
			}
		]
	}
	`
	microsoftTeamsWebhooksServiceListWant = []newreleases.Webhook{
		{
			ID:   "w1a56q4es0dvtl84tdst4ezma8",
			Name: "My Team",
		},
	}
)
