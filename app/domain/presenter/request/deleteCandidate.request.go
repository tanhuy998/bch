package requestPresenter

type DeleteCandidateRequest struct {
	CandidateUUID string `param:"uuid" validate="required,uuid_rfc4122"`
}
