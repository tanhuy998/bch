package authService

import "app/repository"

type (
	FetchAuthDataService struct {
		UserRepo repository.IUser
	}
)
