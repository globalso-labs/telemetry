/*
 * telemetry
 * resource.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Sat, 14 Sep 2024 21:35:27 -0500 by nick.
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

package internal

import (
	"github.com/denisbrodbeck/machineid"
)

type Option func(*Resource)

type Resource struct {
	id        string
	name      string
	namespace string
	version   string
}

func (r *Resource) GetID() string {
	if r.id == "" {
		m, _ := machineid.ID()
		r.id = m
	}

	return r.id
}

func (r *Resource) GetName() string {
	if r.name == "" {
		r.name = "unknown"
	}

	return r.name
}

func (r *Resource) GetNamespace() string {
	if r.namespace == "" {
		r.namespace = "default"
	}

	return r.namespace
}

func (r *Resource) GetVersion() string {
	return r.version
}

func WithName(name string) Option {
	return func(r *Resource) {
		r.name = name
	}
}

func WithNamespace(namespace string) Option {
	return func(r *Resource) {
		r.namespace = namespace
	}
}

func WithVersion(version string) Option {
	return func(r *Resource) {
		r.version = version
	}
}

func NewResource(opts ...Option) *Resource {
	r := new(Resource)
	for _, opt := range opts {
		opt(r)
	}

	return r
}
