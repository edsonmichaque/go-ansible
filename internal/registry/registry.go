// Copyright (c) 2023 Edson Michaque
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
// SPDX-License-Identifier: MIT

package registry

import (
	"errors"
	"sync"

	"github.com/edsonmichaque/go-ansible/internal/provider"
	"github.com/edsonmichaque/go-ansible/internal/provider/ansible/builtin/apt"
	"github.com/edsonmichaque/go-ansible/internal/provider/ansible/builtin/dnf"
)

var r *Registry
var once sync.Once

func init() {
	once.Do(func() {
		r = &Registry{
			Providers: map[string]ProviderFunc{
				"ansible.builtin.apt": apt.Build,
				"ansible.builtin.dnf": dnf.Build,
			},
		}
	})
}

func Find(name string) (ProviderFunc, error) {
	return r.Find(name)
}

type ProviderFunc func() *provider.Provider

type Registry struct {
	Providers map[string]ProviderFunc
}

func (r Registry) Find(name string) (ProviderFunc, error) {
	f, ok := r.Providers[name]
	if !ok {
		return nil, errors.New("not found")
	}

	return f, nil
}
