package presenter

type GetPendingCampaignRequest struct {
	PivotID       string `url:"p_pivot"`
	PageSizeLimit int    `url:"p_limit" validate:"required"`
	Direction     int    `url:"p_dir"`
	IsPrev        bool   `url:"p_prev"`
	PageNumber    int
}
