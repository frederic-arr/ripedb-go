// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package ripedb

import (
	"fmt"
	"log/slog"
)

// Create a new instance of RipeClient using the provided options.
//
// Parameters:
//   - opts: A pointer to RipeClientOptions containing the configuration settings for the client.
//
// Returns:
//   - *RipeClient: A pointer to the initialized RipeClient instance.
//   - error: An error if the client initialization fails.
//
// Example Usage:
//
//	opts := &RipeClientOptions{ /* configure options */ }
//	client, err := NewRipeClient(opts)
//	if err != nil {
//	    log.Fatalf("Failed to create RipeClient: %v", err)
//	}
func NewRipeClient(opts *RipeClientOptions) (*RipeClient, error) {
	if opts.User != nil && opts.Password == nil {
		return nil, fmt.Errorf("a username was provided without a password")
	}

	if opts.Password != nil && *opts.Password == "" {
		return nil, fmt.Errorf("an empty password was provided")
	}

	if (opts.Certificate != nil || opts.Key != nil) && (opts.Certificate == nil || opts.Key == nil) {
		return nil, fmt.Errorf("incomplete x.509 client authentication parameters")
	}

	if opts.Password != nil && (opts.Certificate != nil || opts.Key != nil) {
		return nil, fmt.Errorf("cannot use X.509 client authentication simulatenously with basic username/password authentication")
	}

	if opts.Password != nil {
		slog.Debug("Using basic authentication")
		return newPasswordClient(opts)
	} else if opts.Certificate != nil {
		slog.Debug("Using X.509 client certificate authentication")
		return newX509Client(opts)
	} else {
		slog.Debug("Using anonymous authentication")
		return newAnonymousClient(opts)
	}
}
