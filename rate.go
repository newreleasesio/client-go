// Copyright (c) 2019, NewReleases Go client AUTHORS.
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package newreleases

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "X-RateLimit-Remaining"
	headerRateReset     = "X-RateLimit-Reset"
	headerRateRetry     = "Retry-After"
)

// Rate contains the request rate limit information.
type Rate struct {
	Limit     int       // The maximum number of requests that the user is permitted to make per hour.
	Remaining int       // The number of requests remaining in the current rate limit window.
	Reset     time.Time // Seconds until current rate limit window will reset to the maximal value.
	Retry     time.Time // Seconds until new requests are permitted when limit is reached.
}

func (r Rate) String() (s string) {
	return fmt.Sprintf("limit: %v, remaining %v, reset at %s", r.Limit, r.Remaining, r.Reset)
}

func (c *Client) setRate(r *http.Response) {
	rate := parseRate(r)

	c.rateMu.Lock()
	c.rate = rate
	c.rateMu.Unlock()
}

// Rate returns the current request rate limit information.
func (c *Client) Rate() (r Rate) {
	c.rateMu.RLock()
	r = c.rate
	c.rateMu.RUnlock()

	return r
}

// parseRate creates a new Rate with information from the response.
func parseRate(r *http.Response) (rate Rate) {
	if limit := r.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}
	if reset := r.Header.Get(headerRateReset); reset != "" {
		if v, _ := strconv.ParseInt(reset, 10, 64); v != 0 {
			rate.Reset = time.Now().Add(time.Duration(v) * time.Second)
		}
	}
	if retry := r.Header.Get(headerRateRetry); retry != "" {
		if v, _ := strconv.ParseInt(retry, 10, 64); v != 0 {
			rate.Retry = time.Now().Add(time.Duration(v) * time.Second)
		}
	}
	return rate
}
