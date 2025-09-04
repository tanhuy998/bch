package middleware

import (
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

type (
	ContainerDependentMiddleware func(container *hero.Container) context.Handler
)
