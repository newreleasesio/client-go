// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// WebhooksService provides information about Webhooks notifications.
type WebhooksService service

// Webhook holds the information about webhook ID and its descriptive name.
type Webhook struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// List returns all webhooks.
func (s *WebhooksService) List(ctx context.Context) (webhooks []Webhook, err error) {
	var r webhooksResponse
	err = s.client.request(ctx, http.MethodGet, "v1/webhooks", nil, &r)
	return r.Webhooks, err
}

type webhooksResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}
