package authServicePort

import (
	"app/src/model"

	"github.com/google/uuid"
)

type (
	ICheckUserInCommandGroup interface {
		Serve(groupUUID, userUUID string) (bool, error)
		Detail(groupUUID uuid.UUID, userUUID uuid.UUID) (*model.CommandGroupUser, error)
	}
)
