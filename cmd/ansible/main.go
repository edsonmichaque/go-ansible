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

package main

import (
	"fmt"

	"github.com/edsonmichaque/go-ansible/registry"
	"gopkg.in/yaml.v3"
)

func main() {
	data := `
tasks:
- name: a
  ansible.builtin.dnf:
    name: httpd
    state: present
- name: b
  ansible.builtin.apt:
    state: present
    pkg:
    - apache2
    - nginx
- name: b
  ansible.builtin.apk:
    state: present
    pkg:
    - apache2
    - nginx
- name: e
  ansible.builtin.user:
    name: edson
`

	var s Play
	if err := yaml.Unmarshal([]byte(data), &s); err != nil {
		panic(err)
	}

	for _, task := range s.Tasks {
		fmt.Printf("Running %s\n", task.Name)

		if len(task.Entries) != 1 {
			panic("exactly one action is required")
		}

		for k, v := range task.Entries {
			buildProvider, ok := registry.R.Providers[k]
			if !ok {
				panic("no provider found")
			}

			provider := buildProvider()

			if provider.Init != nil {
				provider.Init()
			}

			if provider.Run == nil {
				panic("no runner")
			}

			cmd, err := provider.Run(provider, v)
			if err != nil {
				panic(err)
			}

			fmt.Println(cmd)

			if provider.Finish != nil {
				provider.Finish()
			}
		}
	}

}

type Playbook []Play

type Play struct {
	Name       string `yaml:"name"`
	Hosts      string `yaml:"hosts"`
	RemoteUser string `yaml:"remote_user"`
	Tasks      []Task `yaml:"tasks"`
}

type Task struct {
	Name               string               `yaml:"name"`
	Entries            map[string]yaml.Node `yaml:",inline"`
	Loop               interface{}          `yaml:"loop"`
	Vars               interface{}          `yaml:"vars"`
	When               interface{}          `yaml:"when"`
	Register           interface{}          `yaml:"register"`
	Until              interface{}          `yaml:"until"`
	Retries            interface{}          `yaml:"retries"`
	DelegateTo         interface{}          `yaml:"delegate_to"`
	DelegateFacts      interface{}          `yaml:"delegate_facts"`
	LocalAction        interface{}          `yaml:"local_action"`
	Action             interface{}          `yaml:"action"`
	LoopControl        interface{}          `yaml:"loop_control"`
	RunOnce            interface{}          `yaml:"run_once"`
	Lineinfile         interface{}          `yaml:"lineinfile"`
	ImportTasks        interface{}          `yaml:"import_tasks"`
	IncludeTasks       interface{}          `yaml:"include_tasks"`
	IgnoreUnreacheable interface{}          `yaml:"ignore_unreacheable"`
	FailedWhen         interface{}          `yaml:"failed_when"`
}
