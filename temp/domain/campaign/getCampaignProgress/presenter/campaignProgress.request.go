package presenter

type (
	CampaignProgressRequestPresenter struct {
		CampaignUUID string `param:"uuid" validate:"required"`
	}
)
