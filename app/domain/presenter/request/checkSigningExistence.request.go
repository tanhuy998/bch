package requestPresenter

type (
	CheckSigningExistenceRequest struct {
		CampaignUUID  string `url:"campaignUUID"`
		CandidateUUID string `url:"canidateUUID"`
	}
)
