package authServicePort

import "app/src/model"

type (
	IModifyUser interface {
		Serve(userUUID string, data *model.User) error
	}
)
