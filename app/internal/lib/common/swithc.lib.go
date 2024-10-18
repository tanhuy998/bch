package libCommon

type (
	case_t[Ret_T any] struct {
		ret_val Ret_T
		matched bool
	}
)

func Switch[T any]() case_t[T] {

	return case_t[T]{}
}

func (this case_t[Ret_T]) Case(expresion bool, val Ret_T) case_t[Ret_T] {

	if this.matched {

		return this
	}

	this.ret_val = Ternary(expresion, val, this.ret_val)
	this.matched = expresion

	return this
}

func (this case_t[Ret_T]) Return() Ret_T {

	return this.ret_val
}

func (this case_t[Ret_T]) Default(val Ret_T) Ret_T {

	if this.matched {

		return this.ret_val
	}

	return val
}
