// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"

	"newreleases.io/newreleases"
)

func TestClient_key(t *testing.T) {
	client, mux, _, teardown := newClient(t, "myauthkey")
	defer teardown()

	mux.HandleFunc("/v1/providers", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Key") != "myauthkey" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, struct{}{})
	}))

	_, err := client.Providers.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_userAgent(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/providers", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") != newreleases.UserAgent {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, struct{}{})
	}))

	_, err := client.Providers.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

const jsonContentType = "application/json; charset=utf-8"

func newClient(t testing.TB, key string) (client *newreleases.Client, mux *http.ServeMux, baseURL *url.URL, teardown func()) {
	t.Helper()

	mux = http.NewServeMux()
	server := httptest.NewServer(mux)

	baseURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	client = newreleases.NewClient(key, &newreleases.ClientOptions{
		BaseURL:    baseURL,
		HTTPClient: server.Client(),
	})

	teardown = func() {
		server.Close()
	}

	return client, mux, baseURL, teardown
}

func newStaticHandler(body string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, body)
	}
}

func newPagedStaticHndler(pages ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := 1
		if p := r.URL.Query().Get("page"); p != "" {
			var err error
			page, err = strconv.Atoi(p)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, pages[page-1])
	}
}

func requireMethod(method string, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		f(w, r)
	}
}

func assertEqual(t *testing.T, name string, got, want interface{}) {
	t.Helper()

	if name != "" {
		name += ": "
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%sgot %+v, want %+v", name, got, want)
	}
}
