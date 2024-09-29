package candidateService

import (
	"app/domain/model"
	"app/repository"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ICheckSigningExistence interface {
		Serve(campaignUUID string, candidateUUID string) (bool, error)
		RetrievePendingCandidateSigning(campaignUUID string, candidateUUID string) (*model.Candidate, error)
	}

	CheckSigningExistenceService struct {
		CampaignRepository  repository.ICampaignRepository
		CandidateRepository repository.ICandidateRepository
	}
)

func (this *CheckSigningExistenceService) Serve(
	str_campaignUUID string,
	str_candidateUUID string,
) (bool, error) {

	candidate, err := this.RetrievePendingCandidateSigning(str_campaignUUID, str_candidateUUID)

	if err != nil {

		return false, err
	}

	if candidate == nil {

		return false, nil
	}

	return true, nil
}

func (this *CheckSigningExistenceService) RetrievePendingCandidateSigning(
	str_campaignUUID string,
	str_candidateUUID string,
) (*model.Candidate, error) {

	campaignUUID, err := uuid.Parse(str_campaignUUID)

	if err != nil {

		return nil, nil
	}

	isCampaignPending, err := this.checkIsCampaignPending(campaignUUID)

	if err != nil {

		return nil, err
	}

	if !isCampaignPending {

		return nil, nil
	}

	candidateUUID, err := uuid.Parse(str_candidateUUID)

	if err != nil {
		fmt.Println(err.Error(), str_candidateUUID)
		return nil, nil
	}

	// candidate, err := this.CandidateRepository.Find(
	// 	bson.D{
	// 		{"uuid", candidateUUID},
	// 		{"campaignUUID", campaignUUID},
	// 	},
	// 	nil,
	// )

	res, err := this.CandidateRepository.Aggregate(
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"uuid", candidateUUID},
						{"campaignUUID", campaignUUID},
					},
				},
			},
			bson.D{
				{
					"$lookup", bson.D{
						{"from", repository.CAMPAIGN_COLLECTION_NAME},
						{"localField", "campaignUUID"},
						{"foreignField", "uuid"},
						{"as", "campaigns"},
					},
				},
			},
			bson.D{
				{
					"$match", bson.D{
						{"campaigns.0.expire", bson.D{{"$gt", time.Now()}}},
					},
				},
			},
			bson.D{
				{
					"$project", bson.D{
						{"campaigns", 0},
					},
				},
			},
		},
		nil,
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if res == nil || len(res) == 0 {

		return nil, nil
	}

	return res[0], nil
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
