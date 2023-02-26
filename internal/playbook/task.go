package playbook

import (
	"gopkg.in/yaml.v3"
)

type Task struct {
	Name               string               `yaml:"name"`
	RemoteUser         string               `yaml:"remote_user"`
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
