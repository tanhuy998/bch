package authServiceAdapter

import "app/domain/model"

type (
	IGetSingleUserService interface {
		Serve(uuid string) (*model.User, error)
		SearchByUsername(username string) (*model.User, error)
		CheckUsernameExistence(username string) (bool, error)
	}
)
