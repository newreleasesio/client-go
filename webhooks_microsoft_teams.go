// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// MicrosoftTeamsWebhooksService provides information about Microsoft Teams
// Webhooks notifications.
type MicrosoftTeamsWebhooksService service

// List returns all Microsoft Teams webhooks.
func (s *MicrosoftTeamsWebhooksService) List(ctx context.Context) (webhooks []Webhook, err error) {
	var r webhooksResponse
	err = s.client.request(ctx, http.MethodGet, "v1/microsoft-teams-webhooks", nil, &r)
	return r.Webhooks, err
}
