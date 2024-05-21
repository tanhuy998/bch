package requestPresenter

type GetCampaignListRequest struct {
	PivotID       int `query:"p_pivot" validate:"required"`
	PageSizeLimit int `query:"p_limit" validate:"required"`
	Direction     int `query:"p_dir" validate:"required"`
	PageNumber    int
}
