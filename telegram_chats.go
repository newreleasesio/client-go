// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// TelegramChatsService provides information about Telegram notifications.
type TelegramChatsService service

// TelegramChat holds information about a Telegram Chat which receives
// notifications.
type TelegramChat struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// List returns all connected Slack Channels.
func (s *TelegramChatsService) List(ctx context.Context) (channels []TelegramChat, err error) {

	type TelegramChatsResponse struct {
		Chats []TelegramChat `json:"chats"`
	}

	var r TelegramChatsResponse
	err = s.client.request(ctx, http.MethodGet, "v1/telegram-chats", nil, &r)
	return r.Chats, err
}
