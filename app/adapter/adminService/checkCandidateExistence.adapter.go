package adminServiceAdapter

import "github.com/google/uuid"

type (
	ICheckCandidateExistence interface {
		Serve(candidateUUID uuid.UUID) (bool, error)
	}
)
