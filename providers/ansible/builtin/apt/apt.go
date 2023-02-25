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

package apt

import (
	"fmt"
	"strings"

	"github.com/edsonmichaque/go-ansible"
	"gopkg.in/yaml.v3"
)

func Build() *ansible.Provider {
	apt := &apt{}

	return &ansible.Provider{
		Init:   apt.Init,
		Run:    apt.Run,
		Finish: apt.Finish,
	}
}

type apt struct {
	Name              string   `yaml:"name"`
	State             string   `yaml:"state"`
	UpdateCache       bool     `yaml:"update_cache"`
	Upgrade           bool     `yaml:"upgrade"`
	PKG               []string `yaml:"pkg"`
	DefaultRelease    string   `yaml:"default_release"`
	AllowDowngrade    bool     `yaml:"allow_downgrade"`
	FailOnAutoremove  bool     `yaml:"fail_on_autoremove"`
	InstallRecommends bool     `yaml:"install_recommends"`
}

func (a *apt) Init() {
	fmt.Println("apt init")
}

func (a *apt) Finish() {
	fmt.Println("apt finish")
}

type Decoder interface {
	Decode(v interface{}) error
}

func (a *apt) Run(m *ansible.Provider, dec yaml.Node) (string, error) {
	if err := dec.Decode(a); err != nil {
		return "", err
	}

	m.Print("boom")

	var command string
	if a.State == "present" {
		command = "install -y"
	}

	if a.State == "absent" {
		command = "remove"
	}

	if a.Upgrade {
		command = "upgrade"
	}

	var pkg []string
	if a.Name != "" {
		pkg = append(pkg, a.Name)
	}

	if len(a.PKG) != 0 {
		pkg = append(pkg, a.PKG...)
	}

	return fmt.Sprintf("apt %s %s", command, strings.Join(pkg, " ")), nil
}
