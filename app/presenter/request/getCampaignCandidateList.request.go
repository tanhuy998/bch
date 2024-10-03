package requestPresenter

import "github.com/google/uuid"

type GetCampaignCandidateListRequest struct {
	CampaignUUID  *uuid.UUID `param:"campaignUUID"`
	PivotID       string     `url:"p_pivot"`
	PageSizeLimit int        `url:"p_limit" validate:"required"`
	IsPrev        bool       `url:"p_prev"`
	ExposeHeader  bool       `url:"r_exp_header"`
}
