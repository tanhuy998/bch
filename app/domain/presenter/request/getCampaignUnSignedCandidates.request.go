package requestPresenter

type (
	GetCampaignUnSignedCandidates struct {
		CampaignUUID string `param:"campaignUUID"`
		PivotID      string `url:"p_pivot"` // objectID of a mongodb document
		//PivotTime     *time.Time `url:"p_pivot_time"`
		PageSizeLimit int  `url:"p_limit" validate:"required,gt=0"`
		Direction     int  `url:"p_dir"`
		IsPrev        bool `url:"p_prev"`
	}
)
