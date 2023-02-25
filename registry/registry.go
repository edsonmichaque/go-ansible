package registry

import (
	"github.com/edsonmichaque/go-ansible"
	"github.com/edsonmichaque/go-ansible/providers/ansible/builtin/apt"

)

type ProviderFunc func() *ansible.Provider

type Registry struct {
	Providers map[string]ProviderFunc
}

var R = &Registry{
	Providers: map[string]ProviderFunc{
		"ansible.builtin.apt": apt.Build,
	},
}

