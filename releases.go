// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// ReleasesService provides information about releases for every project that is
// tracked.
type ReleasesService service

// Release holds information about a specific released version.
type Release struct {
	Version      string    `json:"version"`
	Date         time.Time `json:"date"`
	IsPrerelease bool      `json:"is_prerelease,omitempty"`
	IsUpdated    bool      `json:"is_updated,omitempty"`
	IsExcluded   bool      `json:"is_excluded,omitempty"`
	HasNote      bool      `json:"has_note,omitempty"`
}

// ListByProjectID returns a paginated list of project releases and the number
// of the last page. The project is referenced by its ID.
func (s *ReleasesService) ListByProjectID(ctx context.Context, projectID string, page int) (releases []Release, lastPage int, err error) {
	return s.list(ctx, projectID, page)
}

// ListByProjectName returns a paginated list of project releases and the number
// of the last page. The project is referenced by its provider and name.
func (s *ReleasesService) ListByProjectName(ctx context.Context, provider, projectName string, page int) (releases []Release, lastPage int, err error) {
	return s.list(ctx, provider+"/"+projectName, page)
}

func (s *ReleasesService) list(ctx context.Context, projectRef string, page int) (releases []Release, lastPage int, err error) {

	type releasesResponse struct {
		Releases   []Release `json:"releases"`
		TotalPages int       `json:"total_pages"`
	}

	var r releasesResponse
	path := "v1/projects/" + projectRef + "/releases"
	if page <= 0 {
		return nil, 0, errInvalidPageNumber
	}
	if page > 1 {
		path += "?page=" + strconv.Itoa(page)
	}
	err = s.client.request(ctx, http.MethodGet, path, nil, &r)
	return r.Releases, r.TotalPages, err
}

// GetByProjectID returns a specific version release for a project referenced by
// its ID.
func (s *ReleasesService) GetByProjectID(ctx context.Context, projectID, version string) (release *Release, err error) {
	return s.get(ctx, projectID, version)
}

// GetByProjectName returns a specific version release for a project referenced
// by its provider and name.
func (s *ReleasesService) GetByProjectName(ctx context.Context, provider, projectName, version string) (release *Release, err error) {
	return s.get(ctx, provider+"/"+projectName, version)
}

func (s *ReleasesService) get(ctx context.Context, projectRef, version string) (release *Release, err error) {
	err = s.client.request(ctx, http.MethodGet, "v1/projects/"+projectRef+"/releases/"+url.PathEscape(version), nil, &release)
	return release, err
}

// ReleaseNote holds information about an additional note for a specific
// version.
type ReleaseNote struct {
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
	URL     string `json:"url,omitempty"`
}

// GetNoteByProjectID returns a specific release note for a project referenced
// by its ID.
func (s *ReleasesService) GetNoteByProjectID(ctx context.Context, projectID string, version string) (release *ReleaseNote, err error) {
	return s.getNote(ctx, projectID, version)
}

// GetNoteByProjectName returns a specific release note for a project referenced
// by its provider and name.
func (s *ReleasesService) GetNoteByProjectName(ctx context.Context, provider, projectName string, version string) (release *ReleaseNote, err error) {
	return s.getNote(ctx, provider+"/"+projectName, version)
}

func (s *ReleasesService) getNote(ctx context.Context, projectRef, version string) (note *ReleaseNote, err error) {
	err = s.client.request(ctx, http.MethodGet, "v1/projects/"+projectRef+"/releases/"+url.PathEscape(version)+"/note", nil, &note)
	return note, err
}
