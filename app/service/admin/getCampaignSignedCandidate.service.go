package adminService

import (
	"app/domain/model"
	"app/repository"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IGetCampaignSignedCandidates interface {
		Serve(
			campaignUUID string,
			candiatePivot_id string,
			limit int,
			isPrevDir bool,
		) (*repository.PaginationPackWithHeader[model.Candidate, model.Campaign], error)
	}

	GetCampaignSignedCandidates struct {
		CandidateRepository repository.ICandidateRepository
	}
)

func (this *GetCampaignSignedCandidates) Serve(
	campaignUUID_str string,
	candiatePivot_id string,
	limit int,
	isPrevDir bool,
) (*repository.PaginationPackWithHeader[model.Candidate, model.Campaign], error) {

	campaignUUID, err := uuid.Parse(campaignUUID_str)

	if err != nil {

		return nil, err
	}

	candidatePivotObjID, err := primitive.ObjectIDFromHex(candiatePivot_id)

	if err != nil {

		return nil, err
	}

}
