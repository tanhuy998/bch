package presenter

import "time"

type UpdatedCampaignData struct {
	Title  *string    `json:"title" validate:"required_without_all"`
	Expire *time.Time `json:"expire" validate:"required_without_all"`
}

type UpdateCampaignRequest struct {
	UUID string               `param:"uuid" validate:"required"`
	Data *UpdatedCampaignData `json:"data" validate:"required"`
}
