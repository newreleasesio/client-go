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
	"strings"
)

// ProjectsService provides information and methods to manage tracked projects.
type ProjectsService service

// Project holds information about a tracked project and its configured options.
type Project struct {
	ID                     string            `json:"id"`
	Name                   string            `json:"name"`
	Provider               string            `json:"provider"`
	URL                    string            `json:"url"`
	EmailNotification      EmailNotification `json:"email_notification,omitempty"`
	SlackIDs               []string          `json:"slack_channels,omitempty"`
	TelegramChatIDs        []string          `json:"telegram_chats,omitempty"`
	DiscordIDs             []string          `json:"discord_channels,omitempty"`
	HangoutsChatWebhookIDs []string          `json:"hangouts_chat_webhooks,omitempty"`
	MSTeamsWebhookIDs      []string          `json:"microsoft_teams_webhooks,omitempty"`
	MattermostWebhookIDs   []string          `json:"mattermost_webhooks,omitempty"`
	WebhookIDs             []string          `json:"webhooks,omitempty"`
	Exclusions             []Exclusion       `json:"exclude_version_regexp,omitempty"`
	ExcludePrereleases     bool              `json:"exclude_prereleases,omitempty"`
	ExcludeUpdated         bool              `json:"exclude_updated,omitempty"`
}

// EmailNotification enumerates available options for email notifications.
type EmailNotification string

// Available email notification options.
var (
	EmailNotificationNone    EmailNotification = "none"
	EmailNotificationInstant EmailNotification = "instant"
	EmailNotificationHourly  EmailNotification = "hourly"
	EmailNotificationDaily   EmailNotification = "daily"
	EmailNotificationWeekly  EmailNotification = "weekly"
	EmailNotificationDefault EmailNotification = "default"
)

// Exclusion holds information about a regular expression used
// to filter release versions.
type Exclusion struct {
	Value   string `json:"value"`
	Inverse bool   `json:"inverse"`
}

type projectsResponse struct {
	Projects   []Project `json:"projects"`
	TotalPages int       `json:"total_pages"`
}

// ProjectListOptions holds information about a project list page.
type ProjectListOptions struct {
	Page     int
	Order    ProjectListOrder
	Reverse  bool
	Provider string
}

// ProjectListOrder enumerates available project list orders.
type ProjectListOrder string

// Available project list orders.
var (
	ProjectListOrderUpdated ProjectListOrder = "updated"
	ProjectListOrderAdded   ProjectListOrder = "added"
	ProjectListOrderName    ProjectListOrder = "name"
)

// List returns a paginated list of tracked projects and the number of the last
// available page.
func (s *ProjectsService) List(ctx context.Context, o ProjectListOptions) (projects []Project, lastPage int, err error) {
	path := "v1/projects"
	if o.Provider != "" {
		path += "/" + o.Provider
	}
	q := make(url.Values)
	if o.Page < 1 {
		return nil, 0, errInvalidPageNumber
	}
	if o.Page > 1 {
		q.Set("page", strconv.Itoa(o.Page))
	}
	if o.Order != "" {
		q.Set("order", string(o.Order))
	}
	if o.Reverse {
		q.Set("reverse", "")
	}
	if query := q.Encode(); query != "" {
		query = strings.ReplaceAll(query, "reverse=", "reverse")
		path += "?" + query
	}

	var r projectsResponse
	err = s.client.request(ctx, http.MethodGet, path, nil, &r)
	return r.Projects, r.TotalPages, err
}

// Search performs a search with provided query on names of all tracked
// projects. Provider argument is optional and all projects are searched if it
// is a blank string.
func (s *ProjectsService) Search(ctx context.Context, query, provider string) (projects []Project, err error) {
	q := make(url.Values)
	q.Set("q", query)
	if provider != "" {
		q.Set("provider", provider)
	}

	type projectsSearchResponse struct {
		Projects []Project `json:"projects"`
	}

	var r projectsSearchResponse
	err = s.client.request(ctx, http.MethodGet, "v1/projects/search?"+q.Encode(), nil, &r)
	return r.Projects, err
}

// GetByID returns a specific project referenced by its ID.
func (s *ProjectsService) GetByID(ctx context.Context, id string) (project *Project, err error) {
	return s.get(ctx, id)
}

// GetByName returns a specific project referenced by its provider and name.
func (s *ProjectsService) GetByName(ctx context.Context, provider, name string) (project *Project, err error) {
	return s.get(ctx, provider+"/"+name)
}

func (s *ProjectsService) get(ctx context.Context, projectRef string) (project *Project, err error) {
	err = s.client.request(ctx, http.MethodGet, "v1/projects/"+projectRef, nil, &project)
	return project, err
}

// ProjectOptions holds information for setting options for a specific project.
// If any of the fields have nil value, the option is not set by Add method or
// changed by UpdateByID or UpdateByName methods. When using update methods,
// removing all elements must be done by setting an initialized slice, not a nil
// slice. For boolean pointer methods, there is a convenient function Bool that
// returns boolean pointer by passing a regular bool value.
type ProjectOptions struct {
	EmailNotification      *EmailNotification `json:"email_notification"`
	SlackIDs               []string           `json:"slack_channels"`
	TelegramChatIDs        []string           `json:"telegram_chats"`
	DiscordIDs             []string           `json:"discord_channels"`
	HangoutsChatWebhookIDs []string           `json:"hangouts_chat_webhooks"`
	MSTeamsWebhookIDs      []string           `json:"microsoft_teams_webhooks"`
	MattermostWebhookIDs   []string           `json:"mattermost_webhooks"`
	WebhookIDs             []string           `json:"webhooks"`
	Exclusions             []Exclusion        `json:"exclude_version_regexp"`
	ExcludePrereleases     *bool              `json:"exclude_prereleases"`
	ExcludeUpdated         *bool              `json:"exclude_updated"`
}

// Add adds a new project to be tracked.
func (s *ProjectsService) Add(ctx context.Context, provider, name string, o *ProjectOptions) (project *Project, err error) {

	type projectAddRequest struct {
		Provider string `json:"provider"`
		Name     string `json:"name"`
		*ProjectOptions
	}

	err = s.client.request(ctx, http.MethodPost, "v1/projects", projectAddRequest{
		Provider:       provider,
		Name:           name,
		ProjectOptions: o,
	}, &project)
	return project, err
}

// UpdateByID changes project options referenced by its ID.
func (s *ProjectsService) UpdateByID(ctx context.Context, id string, o *ProjectOptions) (project *Project, err error) {
	return s.update(ctx, id, o)
}

// UpdateByName changes project options referenced by its provider and name.
func (s *ProjectsService) UpdateByName(ctx context.Context, provider, name string, o *ProjectOptions) (project *Project, err error) {
	return s.update(ctx, provider+"/"+name, o)
}

func (s *ProjectsService) update(ctx context.Context, projectRef string, o *ProjectOptions) (project *Project, err error) {
	err = s.client.request(ctx, http.MethodPost, "v1/projects/"+projectRef, o, &project)
	return project, err
}

// DeleteByID removes a project referenced by its ID.
func (s *ProjectsService) DeleteByID(ctx context.Context, id string) (err error) {
	return s.delete(ctx, id)
}

// DeleteByName removes a project referenced by its provider and name.
func (s *ProjectsService) DeleteByName(ctx context.Context, provider, name string) (err error) {
	return s.delete(ctx, provider+"/"+name)
}

func (s *ProjectsService) delete(ctx context.Context, projectRref string) (err error) {
	return s.client.request(ctx, http.MethodDelete, "v1/projects/"+projectRref, nil, nil)
}
