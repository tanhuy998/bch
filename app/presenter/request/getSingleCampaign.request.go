package requestPresenter

type GetSingleCampaignRequest struct {
	UUID string `param:"uuid" validate:"required"`
}
