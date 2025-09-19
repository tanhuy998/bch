package input

import (
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/middleware/hook/binding"
)

type (
	UseInputAuthorityMapping struct{ annotation.Annotation }
)

func (UseInputAuthorityMapping) Accumulate(asset interface{}) interface{} {

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

func (UseInputAuthorityMapping) GetAccumulatorKey() interface{} {

	return input_annotation_accumulator_key
}

func (m *UseInputAuthorityMapping) Once() {}
