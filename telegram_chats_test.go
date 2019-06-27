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

func TestTelegramChatsService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/telegram-chats", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, telegramChatsServiceList)
	}))

	got, err := client.TelegramChats.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, telegramChatsServiceListWant)
}

var (
	telegramChatsServiceList = `
	{
		"chats": [
			{
				"id": "00q3pe8gf7322d8n52bgrl551",
				"type": "group",
				"name": "My Telegram Chat"
			}
		]
	}
	`
	telegramChatsServiceListWant = []newreleases.TelegramChat{
		{
			ID:   "00q3pe8gf7322d8n52bgrl551",
			Type: "group",
			Name: "My Telegram Chat",
		},
	}
)
