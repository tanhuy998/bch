package activator

import (
	"github.com/kataras/iris/v12/mvc"
)

type (
	IActivator interface {
		mvc.BeforeActivation
		// Handle(httpMethod string, path string, methodName string, middilewares ...context.Handler) *router.Route
		// Router() router.Party
	}
)

type (
	ActivateController struct {
		activator mvc.BeforeActivation
	}
)

func (this *ActivateController) BeforeActivation(activator mvc.BeforeActivation) {

	this.activator = activator
}

// func (this *ActivateController) UseMiddleware(middlewares ...interface{}) {

// 	container := this.activator.Dependencies()

// 	for _, fn := range middlewares {

// 		this.activator.Router().Use(
// 			middleware.TransformMiddleware(container, fn),
// 		)
// 	}
// }

func (this *ActivateController) Activator() IActivator {

	return this.activator
}
