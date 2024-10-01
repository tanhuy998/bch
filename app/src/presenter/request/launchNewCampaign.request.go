package requestPresenter

import "app/src/model"

type LaunchNewCampaignRequest struct {
	Data *model.Campaign `json:"data" validate:"required"`
}
