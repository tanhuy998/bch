package env

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

type (
	env_assertion_val_t string
)

func _parse_int[T constraints.Signed](str string, base int, bitSize int) (val T, ok bool) {

	switch v, err := strconv.ParseInt(string(str), base, bitSize); {
	case err != nil:
		return
	default:
		return T(v), true
	}
}

func _parse_uint[T constraints.Unsigned](str string, bitSize int, base int) (val T, ok bool) {

	switch v, err := strconv.ParseUint(string(str), base, bitSize); {
	case err != nil:
		return
	default:
		return T(v), true
	}
}

func _parse_float[T constraints.Float](str string, bitSize int) (val T, ok bool) {

	switch v, err := strconv.ParseFloat(string(str), bitSize); {
	case err != nil:
		return
	default:
		return T(v), true
	}
}

func (this env_assertion_val_t) ToInt8(base int) (val int8, ok bool) {

	return _parse_int[int8](string(this), base, 8)
}

func (this env_assertion_val_t) ToUint8(base int) (val uint8, ok bool) {

	return _parse_uint[uint8](string(this), base, 8)
}

func (this env_assertion_val_t) ToInt16(base int) (val int16, ok bool) {

	return _parse_int[int16](string(this), base, 16)
}

func (this env_assertion_val_t) ToUint16(base int) (val uint16, ok bool) {

	return _parse_uint[uint16](string(this), base, 16)
}

func (this env_assertion_val_t) ToInt(base int) (val int, ok bool) {

	return _parse_int[int](string(this), base, 32)
}

func (this env_assertion_val_t) ToUint(base int) (val uint, ok bool) {

	return _parse_uint[uint](string(this), base, 32)
}

func (this env_assertion_val_t) ToFloat32() (val float32, ok bool) {

	return _parse_float[float32](string(this), 32)
}

func (this env_assertion_val_t) ToFloat64() (val float64, ok bool) {

	return _parse_float[float64](string(this), 64)
}

func (this env_assertion_val_t) ToBoolean() (val bool, ok bool) {

	switch v, err := strconv.ParseBool(string(this)); {
	case err != nil:
		return
	default:
		return v, true
	}
}

func (this env_assertion_val_t) Default() string {

	return string(this)
}

func (this env_assertion_val_t) String() string {

	return string(this)
}
