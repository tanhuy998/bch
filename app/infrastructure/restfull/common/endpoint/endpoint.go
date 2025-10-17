package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"fmt"

	"github.com/kataras/iris/v12/context"
)

type (
	endpoint_t struct {
		//
		action_endpoint_t
	}
)

func newEndpoint(builder *endpoint_builder_t, middlewares []context.Handler) *endpoint_t {

	return &endpoint_t{
		action_endpoint_t: action_endpoint_t{
			endpoint_default_t: endpoint_default_t{
				builder:   builder,
				activator: builder.acitvator,
			},
			actionFn: nil,
		},
	}
}

func (this *endpoint_t) _register() {

	switch {
	case this.actionFn == nil:
		this.endpoint_default_t._register()
	default:
		this.action_endpoint_t._register()
	}

	this.route.SetName(
		fmt.Sprintf("%s %s %s", this.builder.method, this.builder.path, session.RegisteredControllerMethod()),
	)
}

func (this *endpoint_t) _done() {

	switch {
	case this.actionFn == nil:
		panic(
			fmt.Sprintf(
				"The endpoint %s is built as custom action but have no action func attached to it",
				this.route.Name,
			),
		)
	}
}
