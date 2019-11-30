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

func TestDiscordChannelsService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/discord-channels", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, discordChannelsServiceList)
	}))

	got, err := client.DiscordChannels.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, discordChannelsServiceListWant)
}

var (
	discordChannelsServiceList = `
	{
		"channels": [
			{
				"id": "02bgrl5510q3pe8gf7322d8n5",
				"name": "NewReleases Yeah"
			}
		]
	}
	`
	discordChannelsServiceListWant = []newreleases.DiscordChannel{
		{
			ID:   "02bgrl5510q3pe8gf7322d8n5",
			Name: "NewReleases Yeah",
		},
	}
)
