package presenter

type GetCampaignCandidateListRequest struct {
	CampaignUUID  string `param:"campaignUUID"`
	PivotID       string `url:"p_pivot"`
	PageSizeLimit int    `url:"p_limit" validate:"required"`
	IsPrev        bool   `url:"p_prev"`
	ExposeHeader  bool   `url:"r_exp_header"`
}
