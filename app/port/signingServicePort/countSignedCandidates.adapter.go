package signingServicePort

import (
	"github.com/google/uuid"
)

type (
	ICountSignedCandidates interface {
		Serve(camapignUUID uuid.UUID) (int64, error)
	}
)
