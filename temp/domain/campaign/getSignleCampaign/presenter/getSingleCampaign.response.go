package presenter

import (
	"app/src/model"
)

type (
	GetSingleCampaignResponse struct {
		Message string          `json:"message"`
		Data    *model.Campaign `json:"data"`
	}
)
