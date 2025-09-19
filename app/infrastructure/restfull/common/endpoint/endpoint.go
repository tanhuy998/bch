package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"app/infrastructure/restfull/common/middleware"
	"io"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
)

type (
	IEndpoint interface {
		IEndpointUseMiddleware
		IRoute
		getRoute() *router.Route
		_register()
	}
)

type (
	end_point_t struct {
		builder *endpoint_builder_t
		route   *router.Route
	}
)

func newEndpoint(
	builder *endpoint_builder_t,
) *end_point_t {

	return &end_point_t{builder: builder}
}

func (this *end_point_t) _register() {

	builder := this.builder

	this.route = builder.acitvator.Handle(
		builder.method, builder.path, session.RegisteredControllerMethod(),
		builder.middlewares...,
	).SetName(session.RegisteredControllerMethod())

	this.builder = nil
}

func (this *end_point_t) getRoute() *router.Route {

	return this.route
}

func (this *end_point_t) Middleware(middlewares ...interface{}) {

	for _, m := range middlewares {

		this.route.Use(
			middleware.TransformMiddleware(this.route.Party.ConfigureContainer().Container, m),
		)
	}
}

func (this *end_point_t) Use(handlers ...context.Handler) {

	this.route.Use(handlers...)
}

func (this *end_point_t) Done(handlers ...context.Handler) {

	this.route.Done(handlers...)
}

func (this *end_point_t) IsStatic() bool {

	return this.route.IsStatic()
}

func (this *end_point_t) StaticPath() string {

	return this.route.StaticPath()
}

func (this *end_point_t) GetTitle() string {

	return this.route.GetTitle()
}

func (this *end_point_t) Trace(w io.Writer, stoppedIndex int) {

	this.route.Trace(w, stoppedIndex)
}

func (this *end_point_t) IsOnline() bool {

	return this.route.IsOnline()
}
