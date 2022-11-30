# NewReleases API v1 Go client

[![Go](https://github.com/newreleasesio/client-go/workflows/Go/badge.svg)](https://github.com/newreleasesio/client-go/actions)
[![GoDoc](https://godoc.org/newreleases.io/newreleases?status.svg)](https://godoc.org/newreleases.io/newreleases)
[![NewReleases](https://newreleases.io/badge.svg)](https://newreleases.io/github/newreleasesio/client-go)

Package newreleases is a Go client library for accessing the [NewReleases](https://newreleases.io) v1 API.

You can view the client API docs here: [https://godoc.org/newreleases.io/newreleases](https://godoc.org/newreleases.io/newreleases)

You can view NewReleases API docs here: [https://newreleases.io/api/v1](https://newreleases.io/api/v1)

## Installation

Run `go get newreleases.io/newreleases` from command line.

## Usage

```go
import "newreleases.io/newreleases"
```

Create a new Client, then use the exposed services to access different parts of the API.

## Authentication

Currently, API keys is the only method of
authenticating with the API. You can manage your keys
at the NewReleases [API keys settings page](https://newreleases.io/settings/api-keys).

You can then use your token to create a new Client.

## Features

This client implements all NewReleases API features.

- List projects
- Search projects
- Get project
- Add project
- Update project
- Delete project
- List projects releases
- Get project release
- Get latest non-excluded project release
- Get project release note
- Get tracked providers
- Get added Slack Channels
- Get added Telegram Chats
- Get added Dissord Channels
- Get added Hangouts Chat webhooks
- Get added Microsoft Teams webhooks
- Get added Mattermost webhooks
- Get added Rocket.Chat webhooks
- Get added Matrix Rooms
- Get added custom Webhooks
- List tags
- Get tag
- Add tag
- Update tag
- Delete tag
- Get auth keys

## Examples

To add a new project:

```go
package main

import (
    "context"
    "log"

    "newreleases.io/newreleases"
)

var key = "myapikey"

func main() {
    client := newreleases.NewClient(key, nil)
    p, err := client.Projects.Add(
        context.Background(),
        "github",
        "golang/go",
        newreleases.ProjectOptions{
            EmailNotification: &newreleases.EmailNotificationHourly,
        }
    )
    if err != nil {
        log.Fatal(err)
    }
    log.Print(p.ID)
}
```

List projects with pagination:

```go
func AllProjects(ctx context.Context, client *newreleases.Client) (pp []newreleases.Project, err error) {
    o := &newreleases.ProjectListOptions{
        Page: 1,
    }
    for {
        projects, lastPage, err := client.Projects.List(ctx, o)
        if err != nil {
            return nil, err
        }

        pp = append(pp, projects...)

        if o.Page >= lastPage {
            break
        }
        o.Page++
    }

    return pp, nil
}
```

## Versioning

Each version of the client is tagged and the version is updated accordingly.

This package uses Go modules.

To see the list of past versions, run `git tag`.

## Contributing

We love pull requests! Please see the [contribution guidelines](CONTRIBUTING.md).

## License

This library is distributed under the BSD-style license found in the [LICENSE](LICENSE) file.
