package passwordServicePort

type (
	IPasswordDispatcher interface {
		GetRawUsername() []byte
		GetRawPasword() []byte
		GetSecret() []byte
		SetSecret(rawSecret []byte)
	}

	IPassword interface {
		// Generate(plain string) ([]byte, error)
		// Resolve(model *model.User) error
		// Compare(model *model.User, secret []byte) error

		Generate(plain string) ([]byte, error)
		Resolve(model IPasswordDispatcher) error
		Compare(model IPasswordDispatcher, secret []byte) error
	}
)
