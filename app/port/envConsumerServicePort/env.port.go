package envConsumerServicePort

type (
	IENVConsumer interface {
		Get(key string) ITypeAssertionValue
	}

	ITypeAssertionValue interface {
		ToInt8(base int) (val int8, ok bool)
		ToUint8(base int) (val uint8, ok bool)
		ToInt16(base int) (val int16, ok bool)
		ToUint16(base int) (val uint16, ok bool)
		ToInt(base int) (val int, ok bool)
		ToUint(base int) (val uint, ok bool)
		ToFloat32() (val float32, ok bool)
		ToFloat64() (val float64, ok bool)
		ToBoolean() (val bool, ok bool)
		Default() string
	}
)
