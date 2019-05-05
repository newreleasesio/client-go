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

func TestSlackChannelsService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/slack-channels", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, clackChannelsServiceList)
	}))

	got, err := client.SlackChannels.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, clackChannelsServiceListWant)
}

var (
	clackChannelsServiceList = `
	{
		"channels": [
			{
				"id": "00q3pe8gf7322d8n52bgrl551",
				"channel": "releases",
				"team_name": "My Slack Team"
			}
		]
	}
	`
	clackChannelsServiceListWant = []newreleases.SlackChannel{
		{
			ID:       "00q3pe8gf7322d8n52bgrl551",
			Channel:  "releases",
			TeamName: "My Slack Team",
		},
	}
)
