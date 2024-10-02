package authServicePort

import "github.com/google/uuid"

type (
	ICheckCommandGroupUserRole interface {
		//Serve(groupUUID string, userUUID string, roleUUIDs []string) error
		Compare(groupUUID uuid.UUID, userUUID uuid.UUID, roleUUID []uuid.UUID) (unAssignedRoles []uuid.UUID, err error)
	}
)
