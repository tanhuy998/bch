package adminService

import (
	"app/domain/model"
	libCommon "app/lib/common"
	"app/repository"
	"errors"

	"github.com/google/uuid"
)

type (
	IAddNewCandidate interface {
		Execute(inputCampaign string, candidateModel *model.Candidate) error
	}

	AdminAddNewCandidateToCampaign struct {
		CampaignRepo  repository.ICampaignRepository
		CandidateRepo repository.ICandidateRepository
	}
)

func (this *AdminAddNewCandidateToCampaign) Execute(inputCampaignUUID string, model *model.Candidate) error {

	camUUID, err := uuid.Parse(inputCampaignUUID)

	if err != nil {

		return err
	}

	campaignExists, err := this.checkCampaignExistence(camUUID)

	if err != nil {

		return err
	}

	if !campaignExists {

		return errors.New("Invalid Campaign")
	}

	model.UUID = libCommon.PointerPrimitive(uuid.New())
	model.CampaignUUID = libCommon.PointerPrimitive(camUUID)

	return this.CandidateRepo.Create(model, nil)
}

func (this *AdminAddNewCandidateToCampaign) checkCampaignExistence(campaignUUID uuid.UUID) (bool, error) {

	campaign, err := this.CampaignRepo.FindByUUID(campaignUUID, nil)

	if err != nil {

		return false, err
	}

	return campaign != nil, nil
}
