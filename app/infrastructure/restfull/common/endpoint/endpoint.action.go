package endpoint

import (
	libCommon "app/internal/lib/common"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IActionBuilder interface {
		BuildAction(actionFn interface{})
	}
)

type (
	action_endpoint_t struct {
		endpoint_default_t
		actualActionFn interface{}
		actionFn       interface{}
	}
)

func emptyAction() {

}

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

func newActionEndpoint(builder *endpoint_builder_t, actionFn interface{}) *action_endpoint_t {

	switch {
	case actionFn != nil:
		_assertActionFn(actionFn)
	}

	ret := &action_endpoint_t{
		endpoint_default_t: endpoint_default_t{
			builder: builder,
		},
	}

	ret.actionFn = actionFn

	return ret
}

func (this *action_endpoint_t) BuildAction(actionFn interface{}) {

	switch this.actionFn.(type) {
	case nil:
		_assertActionFn(actionFn)

		this.route.Party.RemoveRoute(
			this.route.Name,
		)

		this.actionFn = actionFn

		this._register()
	default:
		panic(
			fmt.Sprintf(
				"The endpoint %s are built with a valid action function, could not override",
				this.route.Name,
			),
		)
	}
}

func (this *action_endpoint_t) _register() {

	builder := this.builder

	container := builder.acitvator.Router().ConfigureContainer().Container

	registeredActionFn := container.Handler(
		libCommon.Ternary[interface{}](this.actionFn == nil, emptyAction, this.actionFn),
	)

	this.route = builder.acitvator.Router().Handle(
		builder.method, builder.path, append(this.builder.middlewares, registeredActionFn)...,
	)
}

func (this *action_endpoint_t) _buildDone() {

	switch {
	case this.actionFn == nil:
		panic(
			fmt.Sprintf(
				"The enpoint %s is built as action endpoint but no valid action function passed to",
				this.route.Name,
			),
		)
	}
}
