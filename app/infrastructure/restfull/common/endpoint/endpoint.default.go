package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"app/infrastructure/restfull/common/middleware"
	"fmt"
	"io"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type (
	IEndpoint interface {
		IEndpointUseMiddleware
		IRoute
		IActionBuilder
		//SetAction(actionFn interface{})
		getRoute() *router.Route
		_register()
		_buildDone()
	}
)

type (
	endpoint_default_t struct {
		builder   *endpoint_builder_t
		activator mvc.BeforeActivation
		route     *router.Route
	}
)

func newDefaultEndpoint(
	builder *endpoint_builder_t,
) *endpoint_default_t {

	return &endpoint_default_t{builder: builder}
}

func (this *endpoint_default_t) _register() {

	builder := this.builder

	this.route = builder.acitvator.Handle(
		builder.method, builder.path, session.RegisteredControllerMethod(),
		this.builder.middlewares...,
	).SetName(session.RegisteredControllerMethod())
}

func (this *endpoint_default_t) getRoute() *router.Route {

	return this.route
}

func (this *endpoint_default_t) Middleware(middlewares ...interface{}) {

	for _, m := range middlewares {

		this.route.Use(
			middleware.TransformMiddleware(this.route.Party.ConfigureContainer().Container, m),
		)
	}
}

func (this *endpoint_default_t) Use(handlers ...context.Handler) {

	this.route.Use(handlers...)
}

func (this *endpoint_default_t) Done(handlers ...context.Handler) {

	this.route.Done(handlers...)
}

func (this *endpoint_default_t) IsStatic() bool {

	return this.route.IsStatic()
}

func (this *endpoint_default_t) StaticPath() string {

	return this.route.StaticPath()
}

func (this *endpoint_default_t) GetTitle() string {

	return this.route.GetTitle()
}

func (this *endpoint_default_t) Trace(w io.Writer, stoppedIndex int) {

	this.route.Trace(w, stoppedIndex)
}

func (this *endpoint_default_t) IsOnline() bool {

	return this.route.IsOnline()
}

func (this *endpoint_default_t) BuildAction(actionFn interface{}) {

	panic(
		fmt.Sprintf(
			"The endpoint %s are built as controller method mapping action, could not set action",
			this.route.Name,
		),
	)
}

func (this *endpoint_default_t) _buildDone() {

}
