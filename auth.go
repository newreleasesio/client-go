// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net"
	"net/http"
	"strings"
)

// AuthService provides information about API authentication.
type AuthService service

// AuthKey represents API authentication secret key, with its descriptive name
// and authorized networks.
type AuthKey struct {
	Name               string      `json:"name"`
	Secret             string      `json:"secret"`
	AuthorizedNetworks []net.IPNet `json:"authorized_networks"`
}

// List returns all authentication keys.
func (s *AuthService) List(ctx context.Context) (keys []AuthKey, err error) {
	var r authKeysResponse
	err = s.client.request(ctx, http.MethodGet, "v1/auth/keys", nil, &r)
	return r.AuthKeys(), err
}

// GetAuthKeys returns a list of all auth keys for an account by authenticating
// with account's email address and a password. This function can be used to get
// the authentication key without providing it explicitly to the client, but
// first asking for already known (to the user) credentials.
func GetAuthKeys(ctx context.Context, email, password string, o *ClientOptions) (keys []AuthKey, err error) {
	return newBasicAuthClient(email, password, o).Auth.List(ctx)
}

type authKeysResponse struct {
	Keys []authKey `json:"keys"`
}

func (r *authKeysResponse) AuthKeys() (ak []AuthKey) {
	ak = make([]AuthKey, len(r.Keys))
	for i, k := range r.Keys {
		ak[i] = AuthKey{
			Name:               k.Name,
			Secret:             k.Secret,
			AuthorizedNetworks: k.authorizedNetworks(),
		}
	}
	return ak
}

type authKey struct {
	Name               string  `json:"name"`
	Secret             string  `json:"secret"`
	AuthorizedNetworks []ipNet `json:"authorized_networks"`
}

func (k *authKey) authorizedNetworks() (an []net.IPNet) {
	an = make([]net.IPNet, len(k.AuthorizedNetworks))
	for i, n := range k.AuthorizedNetworks {
		an[i] = net.IPNet(n)
	}
	return an
}

type ipNet net.IPNet

func (n *ipNet) UnmarshalJSON(data []byte) (err error) {
	_, nn, err := net.ParseCIDR(strings.Trim(string(data), `"`))
	if err != nil {
		return err
	}
	*n = ipNet(*nn)
	return nil
}
