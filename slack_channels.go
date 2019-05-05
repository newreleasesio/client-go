// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// SlackChannelsService provides information about Slack notifications.
type SlackChannelsService service

// SlackChannel holds information about a Slack Channel that is connected to the
// account.
type SlackChannel struct {
	ID       string `json:"id"`
	Channel  string `json:"channel"`
	TeamName string `json:"team_name"`
}

// List returns all connected Slack Channels.
func (s *SlackChannelsService) List(ctx context.Context) (channels []SlackChannel, err error) {

	type slackChannelsResponse struct {
		Channels []SlackChannel `json:"channels"`
	}

	var r slackChannelsResponse
	err = s.client.request(ctx, http.MethodGet, "v1/slack-channels", nil, &r)
	return r.Channels, err
}
