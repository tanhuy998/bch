package adminService

import (
	"app/app/model"

	"github.com/google/uuid"
)

type (
	ICampaignAdminManipulationService interface {
		LaunchNewCampaign(*model.Campaign) error
		ModifyExistingCampaign(uuid.UUID, *model.Campaign) error
	}

	IAdminCampaignDelectionService interface {
		DeleteCampaign(uuid.UUID) error
	}

	ICandidateAdminManipulationService interface {
		ModifyExistingCandidate(uuid.UUID, *model.Candidate) error
		DeleteCandidate(uuid.UUID) error
	}
)
