package requestPresenter

import "github.com/google/uuid"

type GetSingleCampaignRequest struct {
	UUID *uuid.UUID `param:"uuid" validate:"required"`
}
