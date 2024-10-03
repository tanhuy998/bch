package requestPresenter

import "github.com/google/uuid"

type DeleteCampaignRequest struct {
	UUID *uuid.UUID `param:"uuid"`
}
