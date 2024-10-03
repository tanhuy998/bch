package requestPresenter

import "github.com/google/uuid"

type GetCampaignSignedCandidatesRequest struct {
	CampaignUUID *uuid.UUID `param:"campaignUUID"`
	PivotID      string     `url:"p_pivot"`
	//PivotTime     *time.Time `url:"p_pivot_time"`
	PageSizeLimit int  `url:"p_limit" validate:"required,gt=0"`
	Direction     int  `url:"p_dir"`
	IsPrev        bool `url:"p_prev"`
}
