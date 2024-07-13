package requestPresenter

type GetCampaignSignedCandidatesRequest struct {
	CampaignUUID  string `param:"campaignUUID"`
	PivotID       string `url:"p_pivot"`
	PageSizeLimit int    `url:"p_limit" validate:"required"`
	Direction     int    `url:"p_dir"`
	IsPrev        bool   `url:"p_prev"`
}
