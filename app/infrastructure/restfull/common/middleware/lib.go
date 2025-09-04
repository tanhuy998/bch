package middleware

import (
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

func TransformMiddleware(container *hero.Container, m interface{}) context.Handler {

	switch actual := m.(type) {
	case ContainerDependentMiddleware:
		return actual(container)
	case context.Handler:
		return actual
	default:
		panic("invalid middleware type, bad type given")
	}
}
