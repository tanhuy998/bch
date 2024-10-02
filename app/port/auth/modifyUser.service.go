package authServicePort

import "app/model"

type (
	IModifyUser interface {
		Serve(userUUID string, data *model.User) error
	}
)
