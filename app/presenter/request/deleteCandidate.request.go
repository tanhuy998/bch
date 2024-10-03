package requestPresenter

import "github.com/google/uuid"

type DeleteCandidateRequest struct {
	CandidateUUID *uuid.UUID `param:"uuid" validate="required,uuid_rfc4122"`
}
