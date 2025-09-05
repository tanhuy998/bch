package endpoint

import (
	"app/infrastructure/restfull/common/middleware"
	"app/infrastructure/restfull/common/middleware/hook"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/hero"
)

type (
	EndpointAction string

	AuthorityConstraint = hook.AuthorityConstraint
)

type (
	IEndpointUseMiddleware interface {
		Middleware(middlewares ...interface{})
	}

	IEndpointAuthenticate interface {
		Authenticate() IEndpoint
	}

	IEndpointAuthorize interface {
		Authorize(
			constraints ...AuthorityConstraint,
		) IEndpoint
	}

	IEndpoint interface {
		IEndpointUseMiddleware
		getRoute() *router.Route
	}

	IEndpointBuilder interface {
		IEndpointUseMiddleware

		UseMiddleware(middlewares ...interface{}) IEndpointBuilder
		Build() IEndpoint
	}
)

type (
	endpoint_builder_t struct {
		route *router.Route
		auth  context.Handler
	}
)

func NewEnpointBuilder(r *router.Route) *endpoint_builder_t {

	ret := &endpoint_builder_t{
		route: r,
	}

	return ret
}

func (this *endpoint_builder_t) getContainer() *hero.Container {

	return this.route.Party.ConfigureContainer().EnableStructDependents().Container
}

func (this *endpoint_builder_t) UseMiddleware(middlewares ...interface{}) IEndpointBuilder {

	this._middleware(middlewares)

	return this
}

func (this *endpoint_builder_t) Middleware(middlewares ...interface{}) {

	this._middleware(middlewares)
}

func (this *endpoint_builder_t) _middleware(middlewares []interface{}) {

	for _, fn := range middlewares {

		this.route.Use(
			middleware.TransformMiddleware(this.getContainer(), fn),
		)
	}
}

func (this *endpoint_builder_t) Authenticate() IEndpoint {

	this.auth = middleware.Auth()(this.getContainer())

	return this
}

func (this *endpoint_builder_t) Authorize(constraints ...AuthorityConstraint) IEndpoint {

	this.auth = middleware.Auth(constraints...)(this.getContainer())

	return this
}

func (this *endpoint_builder_t) getRoute() *router.Route {

	return this.route
}

func (this *endpoint_builder_t) Build() IEndpoint {

	return this
}
