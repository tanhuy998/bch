package requestPresenter

type GetCampaignCandidateListRequest struct {
	CampaignUUID  string `param:"campaignUUID"`
	PivotID       string `url:"p_pivot"`
	PageSizeLimit int    `url:"p_limit" validate:"required"`
	IsPrev        bool   `url:"p_prev"`
}
