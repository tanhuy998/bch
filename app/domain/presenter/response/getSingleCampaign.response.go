package responsePresenter

import "app/domain/model"

type (
	GetSingleCampaignResponse struct {
		Message string          `json:"message"`
		Data    *model.Campaign `json:"data"`
	}
)
