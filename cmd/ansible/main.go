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
	"context"
	"fmt"

	"github.com/edsonmichaque/go-ansible/internal/playbook"
	"github.com/edsonmichaque/go-ansible/internal/registry"
	"gopkg.in/yaml.v3"
)

func main() {
	data := `
tasks:
- name: a
  ansible.builtin.dnf:
    state: present
    pkg:
    - apache2
    - nginx
- name: b
  ansible.builtin.apt:
    state: present
    pkg:
    - apache2
    - nginx
`

	var s playbook.Play
	if err := yaml.Unmarshal([]byte(data), &s); err != nil {
		panic(err)
	}

	for _, task := range s.Tasks {
		fmt.Printf("Running %s\n", task.Name)

		if len(task.Entries) != 1 {
			panic("exactly one action is required")
		}

		for k, v := range task.Entries {
			buildProvider, err := registry.Find(k)
			if err != nil {
				panic("no provider found")
			}

			foundProvider := buildProvider()

			if foundProvider.Initializer != nil {
				foundProvider.Init(context.Background())
			}

			if foundProvider.Runner == nil {
				panic("no runner")
			}

			cmd, err := foundProvider.Run(context.Background(), foundProvider, v)
			if err != nil {
				panic(err)
			}

			fmt.Println(cmd)
		}
	}

}
