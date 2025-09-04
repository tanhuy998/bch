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
		UseMiddleware(middlewares ...interface{})
	}

	IEndpointAuthenticate interface {
		Authenticate() IEndpoint
	}

	IEndpointAuthorize interface {
		Authorize(
			constraints ...AuthorityConstraint,
		) IEndpoint
	}

	IEndpointInitiator interface {
		IEndpoint
		IEndpointAuthenticate
		IEndpointAuthorize
	}

	IEndpoint interface {
		//IEndpointUseMiddleware
		//getEndpointInititor() IEndpointInitiator
		//Middleware(middlewares ...interface{})
		IEndpointUseMiddleware
		Middleware(middlewares ...interface{}) IEndpoint
		getRoute() *router.Route
	}
)

type (
	controller_endpoint_t struct {
		route *router.Route
		auth  context.Handler
	}
)

func NewEnpoint(r *router.Route) *controller_endpoint_t {

	ret := &controller_endpoint_t{
		route: r,
	}

	return ret
}

func (this *controller_endpoint_t) getContainer() *hero.Container {

	return this.route.Party.ConfigureContainer().EnableStructDependents().Container
}

func (this *controller_endpoint_t) UseMiddleware(middlewares ...interface{}) {

	this._middleware(middlewares)

	//return this
}

func (this *controller_endpoint_t) Middleware(middlewares ...interface{}) IEndpoint {

	this._middleware(middlewares)

	return this
}

func (this *controller_endpoint_t) _middleware(middlewares []interface{}) {

	for _, fn := range middlewares {

		this.route.Use(
			middleware.TransformMiddleware(this.getContainer(), fn),
		)
	}
}

func (this *controller_endpoint_t) Authenticate() IEndpoint {

	this.auth = middleware.Auth()(this.getContainer())

	return this
}

func (this *controller_endpoint_t) Authorize(constraints ...AuthorityConstraint) IEndpoint {

	this.auth = middleware.Auth(constraints...)(this.getContainer())

	return this
}

func (this *controller_endpoint_t) getRoute() *router.Route {

	return this.route
}
