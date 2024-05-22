package requestPresenter

type GetSingleCampaignRequest struct {
	UUID string `param:"uuid" validate:"required,uuid_rfc4122"`
}
