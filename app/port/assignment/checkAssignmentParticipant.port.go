package assignmentServicePort

import (
	"app/unitOfWork/aggregate"
	"context"

	"github.com/google/uuid"
)

type (
	ICheckAssignmentParticipantContext interface {
		context.Context
		aggregate.IDomainContext
		aggregate.IDomainUser
		GetAssignmentUUID() uuid.UUID
	}
)

type (
	ICheckAssigmnetParticipant interface {
		Serve(ctx ICheckAssignmentParticipantContext) error
	}
)
