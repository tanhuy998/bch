package authServiceAdapter

import "app/domain/model"

type (
	ICreateUserService interface {
		Serve(username string, password string, name string) (*model.User, error)
	}
)
