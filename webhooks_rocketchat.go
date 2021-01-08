// Copyright (c) 2021, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// RocketchatWebhooksService provides information about Rocket.Chat
// Webhooks notifications.
type RocketchatWebhooksService service

// List returns all Rocket.Chat webhooks.
func (s *RocketchatWebhooksService) List(ctx context.Context) (webhooks []Webhook, err error) {
	var r webhooksResponse
	err = s.client.request(ctx, http.MethodGet, "v1/rocketchat-webhooks", nil, &r)
	return r.Webhooks, err
}
