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

package provider

import (
	"context"
	"fmt"

	"gopkg.in/yaml.v3"
)

type RunnerFunc func(context.Context, *Provider, yaml.Node) (string, error)

func (r RunnerFunc) Run(ctx context.Context, provider *Provider, node yaml.Node) (string, error) {
	return r(ctx, provider, node)
}

type Runner interface {
	Run(context.Context, *Provider, yaml.Node) (string, error)
}

type PreRunnerFunc func(context.Context, *Provider, yaml.Node) (string, error)

func (r PreRunnerFunc) Run(ctx context.Context, provider *Provider, node yaml.Node) (string, error) {
	return r(ctx, provider, node)
}

type PreRunner interface {
	Run(context.Context, *Provider, yaml.Node) (string, error)
}

type PostRunnerFunc func(context.Context, *Provider, yaml.Node) (string, error)

func (r PostRunnerFunc) Run(ctx context.Context, provider *Provider, node yaml.Node) (string, error) {
	return r(ctx, provider, node)
}

type PostRunner interface {
	Run(context.Context, *Provider, yaml.Node) (string, error)
}

type InitializerFunc func(ctx context.Context)

func (i InitializerFunc) Init(ctx context.Context) {
	i(ctx)
}

type Initializer interface {
	Init(context.Context)
}

type Provider struct {
	ID          string
	Initializer Initializer
	PreRun      PreRunner
	Runner      Runner
	PostRunner  PostRunner
}

func (m Provider) Print(data ...interface{}) {
	fmt.Println(data...)
}

func (m Provider) Printf(format string, data ...interface{}) {
	fmt.Printf(format, data...)
}
