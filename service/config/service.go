package config

import (
	"github.com/gernest/legend/core"
	"golang.org/x/net/context"
)

type Service struct{}

func (s *Service) Init(ctx context.Context, id string) error {
	return nil
}

func (s *Service) Serve(talk core.Talk) {}

func (s *Service) Info() core.ServiceInfo {
	return nil
}
