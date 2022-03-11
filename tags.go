// Copyright (c) 2022, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"context"
	"net/http"
)

// TagsService provides information about project Tags.
type TagsService service

// Tag holds the information about tag ID and its descriptive name.
type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Get returns the tag by its ID.
func (s *TagsService) Get(ctx context.Context, id string) (tag *Tag, err error) {
	err = s.client.request(ctx, http.MethodGet, "v1/tags/"+id, nil, &tag)
	return tag, err
}

// List returns all tags.
func (s *TagsService) List(ctx context.Context) (tags []Tag, err error) {

	type tagsResponse struct {
		Tags []Tag `json:"tags"`
	}

	var r tagsResponse
	err = s.client.request(ctx, http.MethodGet, "v1/tags", nil, &r)
	return r.Tags, err
}

type tagOptionsRequest struct {
	Name string `json:"name"`
}

// Add adds a new tag.
func (s *TagsService) Add(ctx context.Context, name string) (tag *Tag, err error) {
	err = s.client.request(ctx, http.MethodPost, "v1/tags", tagOptionsRequest{
		Name: name,
	}, &tag)
	return tag, err
}

// Update changes the name of the tag referenced by the ID.
func (s *TagsService) Update(ctx context.Context, id, name string) (tag *Tag, err error) {
	err = s.client.request(ctx, http.MethodPost, "v1/tags/"+id, tagOptionsRequest{
		Name: name,
	}, &tag)
	return tag, err
}

// Delete removes the tag by its ID.
func (s *TagsService) Delete(ctx context.Context, id string) error {
	return s.client.request(ctx, http.MethodDelete, "v1/tags/"+id, nil, nil)
}
