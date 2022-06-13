// Copyright (c) 2022, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"newreleases.io/newreleases"
)

func TestMatrixRoomsService_List(t *testing.T) {
	client, mux, _, teardown := newClient(t, "")
	defer teardown()

	mux.HandleFunc("/v1/matrix-rooms", requireMethod("GET", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonContentType)
		fmt.Fprintln(w, matrixRoomsServiceList)
	}))

	got, err := client.MatrixRooms.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(t, "", got, matrixRoomsServiceListWant)
}

var (
	matrixRoomsServiceList = `
	{
		"rooms": [
			{
				"id": "lst4ezma884w1a56q4es0dvttd",
				"name": "My Room",
				"homeserver_url": "https://matrix-client.matrix.org",
				"internal_room_id": "!CkgGilbIulFwycshYK:matrix.org"
			}
		]
	}
	`
	matrixRoomsServiceListWant = []newreleases.MatrixRoom{
		{
			ID:             "lst4ezma884w1a56q4es0dvttd",
			Name:           "My Room",
			HomeserverURL:  "https://matrix-client.matrix.org",
			InternalRoomID: "!CkgGilbIulFwycshYK:matrix.org",
		},
	}
)
