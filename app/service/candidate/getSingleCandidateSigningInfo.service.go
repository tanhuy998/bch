package candidateService

import (
	"app/model"
	"app/repository"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	IGetSingleCandidateSigningInfo interface {
		Serve(campaignUUID_str string, candidateUUID_str string) (*model.CandidateSigningInfo, error)
	}

	GetSingleCandidateSigningInfoService struct {
		CandidateRepo repository.ICandidateRepository
		CampaignRepo  repository.ICampaignRepository
	}
)

func (this *GetSingleCandidateSigningInfoService) Serve(
	campaignUUID_str string, candidateUUID_str string,
) (*model.CandidateSigningInfo, error) {

	campaignUUID, err := uuid.Parse(campaignUUID_str)

	if err != nil {

		return nil, err
	}

	candidateUUID, err := uuid.Parse(candidateUUID_str)

	if err != nil {

		return nil, err
	}

	candidate, err := this.CandidateRepo.Find(
		bson.D{
			{"uuid", candidateUUID},
			{repository.CAMPAIGN_UUID_KEY, campaignUUID},
		},
		nil,
	)

	if err != nil {

		return nil, err
	}

	if *(candidate.CampaignUUID) != campaignUUID {

		return nil, errors.New("cannot find candidate corresponding to given campaign") // common.ERR_HTTP_NOT_FOUND
	}

	return candidate.SigningInfo, nil
}
