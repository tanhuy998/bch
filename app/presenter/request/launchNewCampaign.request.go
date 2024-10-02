package requestPresenter

import "app/model"

type LaunchNewCampaignRequest struct {
	Data *model.Campaign `json:"data" validate:"required"`
}
