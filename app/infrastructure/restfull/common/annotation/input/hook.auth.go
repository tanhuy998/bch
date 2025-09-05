package input

import (
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/middleware/hook/binding"
)

type (
	BindInputAuthority struct{ annotation.Annotation }
)

func (BindInputAuthority) Accumulate(asset interface{}) interface{} {

	switch payload := asset.(type) {
	case *binding_t:
		payload.hooks = append(payload.hooks, binding.UseAuthority)
		return payload
	default:
		ret := new(binding_t)
		ret.hooks = append(ret.hooks, binding.UseAuthority)
		return ret
	}
}

func (BindInputAuthority) GetAccumulatorKey() interface{} {

	return input_annotation_accumulator_key
}

func (m *BindInputAuthority) Once() {}
