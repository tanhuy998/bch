package libCommon

func Ternary[T any](criteria bool, valIfTrue T, valIfFalse T) T {

	if criteria {

		return valIfTrue
	}

	return valIfFalse
}

func PointerPrimitive[T any](val T) *T {

	ret := val
	return &ret
}
