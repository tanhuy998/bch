package passwordService

import (
	"app/domain/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	COST = 12
)

var (
	ERR_EMPTY_UNNAME_OR_PW = errors.New("empty username or password")
)

type (
	IPassword interface {
		Generate(plain string) ([]byte, error)
		Resolve(model *model.User) error
		Compare(model *model.User, secret []byte) error
	}

	PasswordService struct {
	}
)

func (this *PasswordService) Generate(plain string) ([]byte, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), COST)

	if err != nil {

		return nil, err
	}

	return hashed, nil
}

func (this *PasswordService) Compare(model *model.User, secret []byte) error {

	gen_secret, err := merge([]byte(model.Username), []byte(model.PassWord))

	if err != nil {

		return err
	}

	return bcrypt.CompareHashAndPassword(gen_secret, secret)
}

func (this *PasswordService) Resolve(model *model.User) error {

	secret, err := merge([]byte(model.Username), []byte(model.PassWord))

	if err != nil {

		return err
	}

	model.Secret = secret
	return nil
}

func merge(uname []byte, pw []byte) ([]byte, error) {

	if len(uname) == 0 || len(pw) == 0 {

		return nil, ERR_EMPTY_UNNAME_OR_PW
	}

	var (
		minSlice []byte
		maxSlice []byte
		oddPart  []byte
		ret      []byte
	)

	minSlice = uname
	maxSlice = pw

	if len(uname) < len(pw) {

		maxSlice, minSlice = swap(minSlice, maxSlice)
	}

	oddPart = maxSlice[len(minSlice)-1:]
	ret = make([]byte, 0)

	for i, val := range minSlice {

		ret = append(ret, val, maxSlice[i])
	}

	ret = append(ret, oddPart...)

	return ret, nil
}

func swap[T any](a T, b T) (T, T) {

	return b, a
}
