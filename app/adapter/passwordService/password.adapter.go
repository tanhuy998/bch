package passwordServiceAdapter

import "app/domain/model"

type (
	IPassword interface {
		Generate(plain string) ([]byte, error)
		Resolve(model *model.User) error
		Compare(model *model.User, secret []byte) error
	}
)
