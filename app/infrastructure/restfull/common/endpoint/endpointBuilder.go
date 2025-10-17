package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"app/infrastructure/restfull/common/middleware"
	"app/infrastructure/restfull/common/middleware/hook"
	"io"

	"slices"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
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

	IEndpointBuilder interface {
		//IEndpointUseMiddleware
		UseMiddleware(middlewares ...interface{}) IEndpointBuilder
		BuildAction(action interface{}) IEndpoint
		Build() IEndpoint
	}

	IRoute interface {
		Use(handlers ...context.Handler)
		Done(handlers ...context.Handler)
		IsStatic() bool
		IsOnline() bool
		StaticPath() string
		GetTitle() string
		Trace(w io.Writer, stoppedIndex int)
	}
)

type (
	endpoint_builder_t struct {
		acitvator   mvc.BeforeActivation
		method      string
		path        string
		middlewares []context.Handler
	}
)

func NewEnpointBuilder(
	activator mvc.BeforeActivation,
	method, path string,
) *endpoint_builder_t {

	ret := &endpoint_builder_t{
		acitvator: activator,
		method:    method,
		path:      path,
	}

	return ret
}

func (this *endpoint_builder_t) getContainer() *hero.Container {

	return this.acitvator.Dependencies()
}

func (this *endpoint_builder_t) UseMiddleware(middlewares ...interface{}) IEndpointBuilder {

	this._middleware(middlewares)

	return this
}

func (this *endpoint_builder_t) Middleware(middlewares ...interface{}) {

	session.InReservation()
	this._middleware(middlewares)
}

func (this *endpoint_builder_t) _middleware(middlewares []interface{}) {

	switch {
	case this.middlewares == nil:
		this.middlewares = make([]context.Handler, len(middlewares))
	case len(this.middlewares) > 0:
		this.middlewares = slices.Grow(this.middlewares, len(middlewares))
	}

	for _, fn := range middlewares {
		this.middlewares = append(this.middlewares, middleware.TransformMiddleware(this.getContainer(), fn))
	}
}

func (this *endpoint_builder_t) Build() IEndpoint {

	return newDefaultEndpoint(this)
}

func (this *endpoint_builder_t) BuildAction(actionFn interface{}) IEndpoint {

	return newActionEndpoint(this, actionFn)
}
