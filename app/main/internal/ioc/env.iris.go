package ioc

import (
	"app/internal/env"
	irisIoc "app/internal/lib/iris/ioc"
	"app/port/envConsumerServicePort"

	"github.com/kataras/iris/v12/hero"
)

func InitializeENV(container *hero.Container) {

	irisIoc.BindDependency[envConsumerServicePort.IENVConsumer, env.ENVConsumer](container, nil)
}
