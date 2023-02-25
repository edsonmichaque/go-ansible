package ansible

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type RunFunc func(*Provider, yaml.Node) (string, error)

type Provider struct {
	PreInit  func()
	Init     func()
	PostInit func()
	PreRun   func()
	Run      RunFunc
	PostRun  func()
	Finish   func()
}

func (m Provider) Print(data ...interface{}) {
	fmt.Println(data...)
}
