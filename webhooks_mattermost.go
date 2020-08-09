// Copyright (c) 2020, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// MattermostWebhooksService provides information about Mattermost
// Webhooks notifications.
type MattermostWebhooksService service

// List returns all Mattermost webhooks.
func (s *MattermostWebhooksService) List(ctx context.Context) (webhooks []Webhook, err error) {
	var r webhooksResponse
	err = s.client.request(ctx, http.MethodGet, "v1/mattermost-webhooks", nil, &r)
	return r.Webhooks, err
}
