package requestPresenter

type (
	CampaignProgressRequestPresenter struct {
		CampaignUUID string `param:"uuid" validate:"required"`
	}
)
