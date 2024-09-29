package requestPresenter

type GetSingleCandidateSigningInfoRequest struct {
	CandidateUUID string `param:"candidateUUID" validate:"required,uuid_rfc4122"`
}
