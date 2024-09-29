package requestPresenter

import "app/domain/model"

type LaunchNewCampaignRequest struct {
	Data *model.Campaign `json:"data" validate:"required"`
}
