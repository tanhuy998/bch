package requestPresenter

import "app/domain/model"

type UpdateCampaignRequest struct {
	UUID string          `json:"uuid" param:"uuid"`
	Data *model.Campaign `json:"data"`
}
