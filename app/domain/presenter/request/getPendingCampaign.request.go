package requestPresenter

type GetPendingCampaignRequest struct {
	PageNumber int `param:"pageNumber" validate:"required"`
}
