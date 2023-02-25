package registry

import (
	"github.com/edsonmichaque/go-ansible"
	"github.com/edsonmichaque/go-ansible/providers/ansible/builtin/apt"

)

type ModuleFunc func() *ansible.Provider

type Registry struct {
	Modules map[string]ModuleFunc
}

var R = &Registry{
	Modules: map[string]ModuleFunc{
		"ansible.builtin.apt": apt.BuildAPT,
	},
}

