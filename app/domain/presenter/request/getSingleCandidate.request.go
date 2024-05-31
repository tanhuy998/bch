package requestPresenter

type GetSingleCandidateRequest struct {
	UUID string `param:"uuid" validate:"required,uuid_rfc4122"`
}
