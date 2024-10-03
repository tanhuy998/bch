package requestPresenter

import "github.com/google/uuid"

type GetSingleCandidateSigningInfoRequest struct {
	CandidateUUID *uuid.UUID `param:"candidateUUID" validate:"required,uuid_rfc4122"`
}
