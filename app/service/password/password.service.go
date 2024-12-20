package passwordService

import (
	"app/internal/common"
	libError "app/internal/lib/error"
	passwordServicePort "app/port/passwordService"
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
		Resolve(model passwordServicePort.IPasswordDispatcher) error
		Compare(model passwordServicePort.IPasswordDispatcher, secret []byte) error
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

func (this *PasswordService) Compare(model passwordServicePort.IPasswordDispatcher, secret []byte) error {

	pw_token, err := merge(model.GetRawUsername(), model.GetRawPasword())

	if err != nil {

		return err
	}

	switch err := bcrypt.CompareHashAndPassword(secret, pw_token); {
	case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
		return errors.Join(
			common.ERR_UNAUTHORIZED, errors.New("invalid username or password"),
		)
	case err != nil:
		return libError.NewInternal(err)
	default:
		return nil
	}
}

func (this *PasswordService) Resolve(model passwordServicePort.IPasswordDispatcher) error {

	pw_token, err := merge(model.GetRawUsername(), model.GetRawPasword())

	if err != nil {

		return err
	}

	secret, err := bcrypt.GenerateFromPassword(pw_token, COST)

	if err != nil {

		return err
	}

	model.SetSecret(secret)
	return nil
}

func merge(uname []byte, pw []byte) ([]byte, error) {

	if len(uname) == 0 || len(pw) == 0 {

		return nil, ERR_EMPTY_UNNAME_OR_PW
	}

	var (
		minSlice []byte
		maxSlice []byte
		//oddPart  []byte
		ret []byte
	)

	minSlice = uname
	maxSlice = pw

	if len(minSlice) > len(maxSlice) {

		maxSlice, minSlice = minSlice, maxSlice
	}

	ret = make([]byte, 0)

	for i, val := range minSlice {

		ret = append(ret, val, maxSlice[i])
	}

	if len(minSlice) != len(maxSlice) {

		//oddPart = maxSlice[len(minSlice)-1:]
		//ret = append(ret, oddPart...)
		ret = append(ret, maxSlice[len(minSlice)-1:]...)
	}

	return ret, nil
}
