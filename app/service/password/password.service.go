package passwordService

import "golang.org/x/crypto/bcrypt"

type (
	IPassword interface {
		Generate(plain string) ([]byte, error)
		Compare(hashed []byte, actualPassword []byte) error
	}

	PasswordService struct {
	}
)

func (this *PasswordService) Generate(plain string) ([]byte, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), 0)

	if err != nil {

		return nil, err
	}

	return hashed, nil
}

func (this *PasswordService) Compare(hashed []byte, actualPassword []byte) error {

	return bcrypt.CompareHashAndPassword(hashed, actualPassword)
}
