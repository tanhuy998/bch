package presenter

type GetSingleCampaignRequest struct {
	UUID string `param:"uuid" validate:"required"`
}
