package group

import (
	"context"
	"errors"

	"github.com/edsonmichaque/go-ansible/internal/provider"
	"gopkg.in/yaml.v3"
)

type linuxProvider struct{}

func (a *linuxProvider) Run(ctx context.Context, m *provider.Provider, dec yaml.Node) (string, error) {
	return "", errors.New("not implemented")
}
