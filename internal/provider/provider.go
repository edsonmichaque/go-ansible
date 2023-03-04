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
	"errors"
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

func (r PreRunnerFunc) PreRun(ctx context.Context, provider *Provider, node yaml.Node) (string, error) {
	return r(ctx, provider, node)
}

type PreRunner interface {
	PreRun(context.Context, *Provider, yaml.Node) error
}

type PostRunnerFunc func(context.Context, *Provider, yaml.Node) (string, error)

func (r PostRunnerFunc) PostRun(ctx context.Context, provider *Provider, node yaml.Node) (string, error) {
	return r(ctx, provider, node)
}

type PostRunner interface {
	PostRun(context.Context, *Provider, yaml.Node) error
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
	Name        string
	Version     string
	Description string
	Author      string

	Initializer Initializer
	PreRunner   PreRunner
	Runner      Runner
	PostRunner  PostRunner
}

func (r Provider) PreRun(ctx context.Context, provider *Provider, node yaml.Node) error {
	return r.PreRunner.PreRun(ctx, provider, node)
}

func (r Provider) Run(ctx context.Context, provider *Provider, node yaml.Node) (string, error) {
	if r.Runner == nil {
		return "", errors.New("no provider find")
	}

	if r.PreRunner != nil {
		if err := r.PreRun(ctx, provider, node); err != nil {
			return "", err
		}
	}

	script, err := r.Runner.Run(ctx, provider, node)
	if err != nil {
		return "", err
	}

	if r.PostRunner != nil {
		if err := r.PostRun(ctx, provider, node); err != nil {
			return script, err
		}
	}

	return script, nil
}

func (r Provider) PostRun(ctx context.Context, provider *Provider, node yaml.Node) error {
	return r.PostRunner.PostRun(ctx, provider, node)
}

func (p Provider) Init(ctx context.Context) {
	p.Initializer.Init(ctx)
}

func (m Provider) Print(data ...interface{}) {
	fmt.Println(data...)
}

func (m Provider) Printf(format string, data ...interface{}) {
	fmt.Printf(format, data...)
}
