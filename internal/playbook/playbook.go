package playbook

type Playbook []Play

type Play struct {
	Name        string `yaml:"name"`
	Hosts       string `yaml:"hosts"`
	RemoteUser  string `yaml:"remote_user"`
	GatherFacts bool   `yaml:"gather_facts"`
	Connection  bool   `yaml:"connection"`
	Tasks       []Task `yaml:"tasks"`
}
