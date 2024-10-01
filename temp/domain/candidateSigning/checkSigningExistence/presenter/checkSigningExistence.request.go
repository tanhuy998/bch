package presenter

type (
	CheckSigningExistenceRequest struct {
		CampaignUUID  string `param:"campaignUUID" validate:"required,uuid_rfc4122"`
		CandidateUUID string `param:"candidateUUID" validate:"required,uuid_rfc4122"`
	}
)
