package endpoint

import (
	"app/infrastructure/restfull/common/activator"
	"app/infrastructure/restfull/common/middleware"
	"app/shared/common/variable"
)

type (
	IAPICurator interface {
		_Curator() *APIEndpointCurator
		Activator() activator.IActivator
	}

	IAPIEndpointCurator interface {
		comparable
		IAPICurator
	}
)

type (
	APIEndpointCurator struct {
		activator.ActivateController
	}
)

func AccquiredRegisteredMethod(in interface{}) string {

	switch v := in.(type) {
	case *variable.Constant[string]:
		return variable.ValueOfConst(*v)
	default:
		panic("(endpoint) invalid type of registered method")
	}
}

func (this *APIEndpointCurator) _Curator() *APIEndpointCurator {

	return this
}

func (this *APIEndpointCurator) UseMiddleware(middlewares ...interface{}) {

	container := this.Activator().Dependencies()

	for _, fn := range middlewares {

		this.Activator().Router().Use(
			middleware.TransformMiddleware(container, fn),
		)
	}
}
