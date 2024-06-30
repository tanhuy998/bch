package candidateService

import (
	"app/repository"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	ICheckSigningExistence interface {
		Serve(campaignUUID string, candidateUUID string) (bool, error)
	}

	CheckSigningExistenceService struct {
		CampaignRepository  repository.ICampaignRepository
		CandidateRepository repository.ICandidateRepository
	}
)

func (this *CheckSigningExistenceService) Serve(str_campaignUUID string, str_candidateUUID string) (bool, error) {

	campaignUUID, err := uuid.Parse(str_campaignUUID)

	if err != nil {

		return false, nil
	}

	isCampaignPending, err := this.checkIsCampaignPending(campaignUUID)

	if err != nil {

		return false, err
	}

	if !isCampaignPending {

		return false, nil
	}

	candidateUUID, err := uuid.Parse(str_candidateUUID)

	if err != nil {
		fmt.Println(err.Error(), str_candidateUUID)
		return false, nil
	}

	candidate, err := this.CandidateRepository.Find(
		bson.D{
			{"uuid", candidateUUID},
			{"campaignUUID", campaignUUID},
		},
		nil,
	)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println(4)
	if candidate == nil {

		return false, nil
	}

	return true, nil
}

func (this *CheckSigningExistenceService) checkIsCampaignPending(campaignUUID uuid.UUID) (bool, error) {

	campaign, err := this.CampaignRepository.FindByUUID(campaignUUID, nil)

	if err != nil {

		return false, err
	}

	if campaign == nil {

		return false, nil
	}

	return campaign.Expire.After(time.Now()), nil
}
