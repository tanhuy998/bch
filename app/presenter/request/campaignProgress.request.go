package requestPresenter

import "github.com/google/uuid"

type (
	CampaignProgressRequestPresenter struct {
		CampaignUUID *uuid.UUID `param:"uuid" validate:"required"`
	}
)
