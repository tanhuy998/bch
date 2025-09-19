package endpoint

import (
	"app/infrastructure/restfull/common/endpoint/internal/session"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type (
	action_endpoint_t struct {
		end_point_t
		actionFn interface{}
	}
)

func _assertActionFn(actionFn interface{}) {

	reflectTypeFunc := reflect.TypeOf(actionFn)

	switch {
	case reflectTypeFunc.Kind() != reflect.Func:
		panic(
			fmt.Sprintf(
				`actionFn must be a function that returns 2 values that implement (%s/mvc.Result, error)`,
				reflect.TypeFor[mvc.Result]().PkgPath(),
			),
		)
	case reflectTypeFunc.NumOut() != 2:
		panic(
			fmt.Sprintf(
				`actionFn must be a function that return 2 values that implement (%s/mvc.Result, error), %d value given`,
				reflect.TypeFor[mvc.Result]().PkgPath(),
				reflectTypeFunc.NumOut(),
			),
		)
	}

	const (
		result = 0
		err    = 1
	)

	switch {
	case !reflectTypeFunc.Out(result).Implements(reflect.TypeFor[mvc.Result]()):
		panic(
			fmt.Sprintf(
				`type of first returned value of action function must implements %s/mvc.Result, (%s, %s) given`,
				reflect.TypeFor[mvc.Result]().PkgPath(),
				reflectTypeFunc.Out(result).Name(),
				reflectTypeFunc.Out(err).Name(),
			),
		)
	case !reflectTypeFunc.Out(err).Implements(reflect.TypeFor[error]()):
		panic(
			fmt.Sprintf(
				`type of second returned value of action function must implements error, (%s, %s) given`,
				reflectTypeFunc.Out(result).Name(),
				reflectTypeFunc.Out(err).Name(),
			),
		)
	}
}

func newActionEndpoint(builder *endpoint_builder_t, actionfn interface{}) *action_endpoint_t {

	return &action_endpoint_t{
		end_point_t: end_point_t{
			builder: builder,
		},
		actionFn: actionfn,
	}
}

func (this *action_endpoint_t) _register() {

	builder := this.builder

	builder.acitvator.Router().ConfigureContainer(

		func(api *router.APIContainer) {

			this.route = api.Handle(
				builder.method, builder.path, this.actionFn,
			).SetName(session.RegisteredControllerMethod())
		},
	)

	this.builder = nil
}
