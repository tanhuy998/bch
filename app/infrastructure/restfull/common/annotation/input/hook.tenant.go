package input

import (
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/middleware/hook/binding"
)

type (
	BindInputTenant struct{ annotation.Annotation }
)

func (BindInputTenant) Accumulate(asset interface{}) interface{} {

	switch payload := asset.(type) {
	case *binding_t:
		payload.hooks = append(payload.hooks, binding.UseTenantMapping)
		return payload
	default:
		ret := new(binding_t)
		ret.hooks = append(ret.hooks, binding.UseTenantMapping)
		return ret
	}
}

func (BindInputTenant) GetAccumulatorKey() interface{} {

	return input_annotation_accumulator_key
}

func (BindInputTenant) Once() {}
