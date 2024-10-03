package requestPresenter

import "github.com/google/uuid"

type (
	CheckSigningExistenceRequest struct {
		CampaignUUID  *uuid.UUID `param:"campaignUUID" validate:"required,uuid_rfc4122"`
		CandidateUUID *uuid.UUID `param:"candidateUUID" validate:"required,uuid_rfc4122"`
	}
)
