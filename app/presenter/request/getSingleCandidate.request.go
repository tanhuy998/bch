package requestPresenter

import "github.com/google/uuid"

type GetSingleCandidateRequest struct {
	UUID *uuid.UUID `param:"uuid" validate:"required,uuid_rfc4122"`
}
