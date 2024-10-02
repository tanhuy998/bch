package responsePresenter

import (
	"app/model"
)

type (
	GetSingleCampaignResponse struct {
		Message string          `json:"message"`
		Data    *model.Campaign `json:"data"`
	}
)
