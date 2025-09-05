package input

import (
	"app/infrastructure/restfull/common/endpoint"
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/middleware/hook/binding"
	"app/shared/common/variable"
	"time"
)

var (
	input_annotation_accumulator_key = variable.NewConstain(time.Now().String())
)

type (
	binding_func func(hooks ...binding.Hook) binding.ContainerDependentMiddleware
)

type (
	binding_t struct {
		funcs []binding_func
		hooks []binding.Hook
	}
)

type (
	Bind[Input_T any] struct {
		annotation.Annotation
	}
)

func (Bind[Input_T]) Accumulate(asset interface{}) interface{} {

	switch payload := asset.(type) {
	case *binding_t:
		payload.funcs = append(payload.funcs, binding.BindPresenters[Input_T, struct{}])
		return payload
	default:
		ret := new(binding_t)
		ret.funcs = append(ret.funcs, binding.BindPresenters[Input_T, struct{}])
		return ret
	}
}

func (Bind[Input_T]) GetAccumulatorKey() interface{} {

	return input_annotation_accumulator_key
}

func (Bind[Input_T]) Apply(endpoint endpoint.IEndpointUseMiddleware, asset interface{}) {

	switch payload := asset.(type) {
	case *binding_t:
		for _, fn := range payload.funcs {

			endpoint.Middleware(
				fn(payload.hooks...),
			)
		}
	default:
		panic("error while retrieving Input annotation binding")
	}
}
