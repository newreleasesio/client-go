// Copyright (c) 2022, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// MatrixRoomsService provides information about Matrix
// notifications.
type MatrixRoomsService service

// MatrixRoom holds information about a Slack Channel that is connected to the
// account.
type MatrixRoom struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	HomeserverURL  string `json:"homeserver_url"`
	InternalRoomID string `json:"internal_room_id"`
}

// List returns all Matrix rooms.
func (s *MatrixRoomsService) List(ctx context.Context) (rooms []MatrixRoom, err error) {
	var r matrixRoomsResponse
	err = s.client.request(ctx, http.MethodGet, "v1/matrix-rooms", nil, &r)
	return r.Rooms, err
}

type matrixRoomsResponse struct {
	Rooms []MatrixRoom `json:"rooms"`
}
