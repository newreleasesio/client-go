// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases_test

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	"newreleases.io/newreleases"
)

func TestAuthService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/auth/keys", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, authServiceList)
	}))

	got, err := client.Auth.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, authServiceListWant)
}

func TestGetAuthKeys(t *testing.T) {
	_, mux, baseURL, teardown := newClient(t, "")
	defer teardown()

	wantUsername := "me@example.com"
	wantPassword := "password12345"

	mux.HandleFunc("/v1/auth/keys", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != wantUsername || password != wantPassword {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, authServiceList)
	}))

	got, err := newreleases.GetAuthKeys(
		context.Background(),
		wantUsername,
		wantPassword,
		&newreleases.ClientOptions{
			BaseURL: baseURL,
		})
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, authServiceListWant)
}

var (
	authServiceList = `
	{
		"keys": [
			{
				"name": "My Key",
				"secret": "ewcppcsk781h1bwxplq3pe8gf7322d8n52bg",
				"authorized_networks": [
					"::/0",
					"0.0.0.0/0"
				]
			},
			{
				"name": "Internal",
				"secret": "awcppcsk781h1bwxplq3pe8gf7322d8n52b1",
				"authorized_networks": [
					"2001:19f0:5:5f0:5400:ff:fe72:ce59/128",
					"173.199.123.52/32"
				]
			}
		]
	}
	`
	authServiceListWant = []newreleases.AuthKey{
		{
			Name:   "My Key",
			Secret: "ewcppcsk781h1bwxplq3pe8gf7322d8n52bg",
			AuthorizedNetworks: []net.IPNet{
				ipNet("::/0"),
				ipNet("0.0.0.0/0"),
			},
		},
		{
			Name:   "Internal",
			Secret: "awcppcsk781h1bwxplq3pe8gf7322d8n52b1",
			AuthorizedNetworks: []net.IPNet{
				ipNet("2001:19f0:5:5f0:5400:ff:fe72:ce59/128"),
				ipNet("173.199.123.52/32"),
			},
		},
	}
)

func ipNet(s string) net.IPNet {
	_, nn, err := net.ParseCIDR(s)
	if err != nil {
		panic(err)
	}
	return *nn
}
