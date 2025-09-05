package session

import (
	"reflect"
	"sync"
)

var builder_session *struct {
	sync.Mutex
	builder                      reflect.Value
	accumulators_bag             map[interface{}]interface{}
	controller_registered_method string
}

type (
	IAccumulator interface {
		Accumulate(interface{}) interface{}
		GetAccumulatorKey() interface{}
	}
)

func __bag() map[interface{}]interface{} {

	return builder_session.accumulators_bag
}

func __getAssetOf(accumulator IAccumulator) (v interface{}, ok bool) {

	key := accumulator.GetAccumulatorKey()

	m := __bag()

	switch len(m) {
	case 0:
		return
	default:
		v, ok = m[key]
		return
	}
}

func AssetOf(accumulator IAccumulator) (asset interface{}) {

	asset, _ = __getAssetOf(accumulator)
	return
}

func Adopt(accumulator IAccumulator) {

	asset, _ := __getAssetOf(accumulator)

	if __bag() == nil {

		builder_session.accumulators_bag = make(map[interface{}]interface{})
	}

	key := accumulator.GetAccumulatorKey()

	builder_session.accumulators_bag[key] = accumulator.Accumulate(asset)
}

func Start(
	builder reflect.Value,
	registered_controller_method string,
) {

	switch {
	case registered_controller_method == "":
		panic("invalid endpoint initialization session, zero value of controller registered method name")
	}

	builder_session = &struct {
		sync.Mutex
		builder                      reflect.Value
		accumulators_bag             map[interface{}]interface{}
		controller_registered_method string
	}{
		builder:                      builder,
		controller_registered_method: registered_controller_method,
	}

	builder_session.Lock()
}

func End() {

	switch builder_session {
	case nil:
		return
	default:
		builder_session.Unlock()
		builder_session = nil
	}
}

func In() bool {

	return builder_session != nil
}

func RegisteredControllerMethod() string {

	return builder_session.controller_registered_method
}
