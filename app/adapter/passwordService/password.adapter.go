package passwordServiceAdapter

type (
	IPassword interface {
		Generate(plain string) ([]byte, error)
		Compare(hashed []byte, actualPassword []byte) error
	}
)
