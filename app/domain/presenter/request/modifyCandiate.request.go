package requestPresenter

import "app/domain/model"

type ModifyCandidateRequest struct {
	UUID      string           `param:"uuid" validate:"required,uuid_rfc4122"`
	Candidate *model.Candidate `jsom:"data' validate:"required"`
}
