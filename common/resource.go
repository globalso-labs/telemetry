/*
 * telemetry
 * resource.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Wed, 25 Sep 2024 21:28:24 -0500 by nick.
 *
 * DISCLAIMER: This software is provided "as is" without warranty of any kind, either expressed or implied. The entire
 * risk as to the quality and performance of the software is with you. In no event will the author be liable for any
 * damages, including any general, special, incidental, or consequential damages arising out of the use or inability
 * to use the software (that includes, but not limited to, loss of data, data being rendered inaccurate, or losses
 * sustained by you or third parties, or a failure of the software to operate with any other programs), even if the
 * author has been advised of the possibility of such damages.
 * If a license file is provided with this software, all use of this software is governed by the terms and conditions
 * set forth in that license file. If no license file is provided, no rights are granted to use, modify, distribute,
 * or otherwise exploit this software.
 */

package common

import (
	"github.com/denisbrodbeck/machineid"
	"go.globalso.dev/x/telemetry/internal"
)

// Option is a function type that modifies a Resource.
type Option func(*Resource)

// Resource represents a telemetry resource with an ID, name, namespace, and version.
type Resource struct {
	id        string
	name      string
	namespace string
	version   string
}

// GetID returns the ID of the Resource.
//
// Returns:
// - string: The ID of the Resource.
func (r *Resource) GetID() string {
	return r.id
}

// GetName returns the name of the Resource.
//
// Returns:
// - string: The name of the Resource.
func (r *Resource) GetName() string {
	return r.name
}

// GetNamespace returns the namespace of the Resource.
//
// Returns:
// - string: The namespace of the Resource.
func (r *Resource) GetNamespace() string {
	return r.namespace
}

// GetVersion returns the version of the Resource.
//
// Returns:
// - string: The version of the Resource.
func (r *Resource) GetVersion() string {
	return r.version
}

// WithName sets the name of the Resource.
//
// Parameters:
// - name string: The name to set.
//
// Returns:
// - Option: An Option function that sets the name of the Resource.
func WithName(name string) Option {
	return func(r *Resource) {
		r.name = name
	}
}

// WithNamespace sets the namespace of the Resource.
//
// Parameters:
// - namespace string: The namespace to set.
//
// Returns:
// - Option: An Option function that sets the namespace of the Resource.
func WithNamespace(namespace string) Option {
	return func(r *Resource) {
		r.namespace = namespace
	}
}

// WithVersion sets the version of the Resource.
//
// Parameters:
// - version string: The version to set.
//
// Returns:
// - Option: An Option function that sets the version of the Resource.
func WithVersion(version string) Option {
	return func(r *Resource) {
		r.version = version
	}
}

// NewResource creates a new Resource with the provided options.
//
// Parameters:
// - opts ...Option: A variadic list of Option functions to configure the Resource.
//
// Returns:
// - *Resource: A pointer to the newly created Resource.
func NewResource(opts ...Option) *Resource {
	var m, _ = machineid.ID()

	r := &Resource{
		id:        m,
		name:      "unknown",
		namespace: "default",
		version:   internal.Version,
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}
