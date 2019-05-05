// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// HangoutsChatWebhooksService provides information about Google Hangouts Chat
// Webhooks notifications.
type HangoutsChatWebhooksService service

// List returns all Google Hangouts Chat webhooks.
func (s *HangoutsChatWebhooksService) List(ctx context.Context) (webhooks []Webhook, err error) {
	var r webhooksResponse
	err = s.client.request(ctx, http.MethodGet, "v1/hangouts-chat-webhooks", nil, &r)
	return r.Webhooks, err
}
