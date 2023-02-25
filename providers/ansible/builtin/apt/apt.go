package apt

import (
	"fmt"
	"github.com/edsonmichaque/go-ansible"
	"gopkg.in/yaml.v3"
	"strings"
)

func BuildAPT() *ansible.Provider {
	apt := &providerApt{}

	return &ansible.Provider{
		Init:   apt.Init,
		Run:    apt.Run,
		Finish: apt.Finish,
	}
}

type providerApt struct {
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

func (a *providerApt) Init() {
	fmt.Println("apt init")
}

func (a *providerApt) Finish() {
	fmt.Println("apt finish")
}

type Decoder interface {
	Decode(v interface{}) error
}

func (a *providerApt) Run(m *ansible.Provider, dec yaml.Node) (string, error) {
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
