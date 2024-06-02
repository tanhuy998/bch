package candidateService

import (
	"app/domain/model"
	"app/internal/common"
	"app/repository"
	"fmt"

	"github.com/google/uuid"
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

	candidate, err := this.CandidateRepo.FindByUUID(candidateUUID, nil)

	if err != nil {

		return nil, err
	}

	if *(candidate.CampaignUUID) != campaignUUID {

		return nil, common.ERR_HTTP_NOT_FOUND
	}
	fmt.Println(candidate)
	return candidate.SigningInfo, nil
}
