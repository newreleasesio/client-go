// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// DiscordChannelsService provides information about Discord notifications.
type DiscordChannelsService service

// DiscordChannel holds information about a Discord Channel that is connected to the
// account.
type DiscordChannel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// List returns all connected Discord Channels.
func (s *DiscordChannelsService) List(ctx context.Context) (channels []DiscordChannel, err error) {

	type discordChannelsResponse struct {
		Channels []DiscordChannel `json:"channels"`
	}

	var r discordChannelsResponse
	err = s.client.request(ctx, http.MethodGet, "v1/discord-channels", nil, &r)
	return r.Channels, err
}
