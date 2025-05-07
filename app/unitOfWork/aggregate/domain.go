package aggregate

import "github.com/google/uuid"

type (
	IDomainUser interface {
		GetUserUUID() uuid.UUID
	}

	IDomainContext interface {
		GetTenantUUID() uuid.UUID
	}

	IAssignmentGroupDomain interface {
		GetAssignmentGroupUUID() uuid.UUID
	}
)
