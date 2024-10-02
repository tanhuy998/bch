package domain

import (
	"github.com/kataras/iris/v12/hero"
)

type (
	ServiceConfigurator struct {
	}
)

func (s *ServiceConfigurator) ConfigureServices(container *hero.Container) {
	panic("TODO: Implement")
}
