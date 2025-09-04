package endpoint

import (
	"app/infrastructure/restfull/common/activator"
	"app/infrastructure/restfull/common/middleware"
	"app/shared/common/variable"
)

type (
	IEndpointBuilder interface {
		_EndpointBuilder() *EndpointBuilder
		Endpoint(
			httpMethod, path, funcName string,
		) IEndpointInitiator
	}
)

type (
	EndpointBuilder struct {
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

func (this *EndpointBuilder) _EndpointBuilder() *EndpointBuilder {

	return this
}

func (this *EndpointBuilder) UseMiddleware(middlewares ...interface{}) {

	container := this.Activator().Dependencies()

	for _, fn := range middlewares {

		this.Activator().Router().Use(
			middleware.TransformMiddleware(container, fn),
		)
	}
}

func (this *EndpointBuilder) Endpoint(
	httpMethod, path, funcName string,
) IEndpointInitiator {

	return NewEnpoint(
		this.Activator().Handle(
			httpMethod, path, funcName,
		),
	)
}
