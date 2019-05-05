// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// ProvidersService provides information about supported project providers and
// providers for projects that are added for tracking.
type ProvidersService service

// List returns all supported project providers.
func (s *ProvidersService) List(ctx context.Context) (providers []string, err error) {
	var r providersResponse
	err = s.client.request(ctx, http.MethodGet, "v1/providers", nil, &r)
	return r.Providers, err
}

// ListAdded returns poviders for projects that are added for tracking.
func (s *ProvidersService) ListAdded(ctx context.Context) (providers []string, err error) {
	var r providersResponse
	err = s.client.request(ctx, http.MethodGet, "v1/providers?added", nil, &r)
	return r.Providers, err
}

type providersResponse struct {
	Providers []string `json:"providers"`
}
