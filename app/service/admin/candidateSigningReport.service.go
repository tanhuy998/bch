package adminService

import (
	"app/domain/valueObject"
	"app/internal/common"
	"app/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ICandidateSigningReport interface {
		Serve(campaignUUID_str string) (*valueObject.CandidateSigningReport, error)
		CountCampaignCandidates(campaignUUID_str string) (int64, error)
		CountSignedCandidates(campaignUUID_str string) (int64, error)
	}

	CandidateSigningReportService struct {
		CandidateSigningCommitRepository repository.ICandidateSigningCommit
		CandidateRepository              repository.ICandidateRepository
		CampaignRepository               repository.ICampaignRepository
	}
)

func (this *CandidateSigningReportService) Serve(
	campaignUUID_str string,
) (*valueObject.CandidateSigningReport, error) {

	campaignUUID, err := uuid.Parse(campaignUUID_str)

	if err != nil {

		return nil, common.ERR_BAD_REQUEST
	}

	candidateSigningReport := valueObject.CandidateSigningReport{}

	totalCount, err := this.countCampaignCandidates(campaignUUID)

	if err != nil {
		fmt.Println("count total err", err)
		return nil, err
	}

	candidateSigningReport.TotalCount = totalCount

	if totalCount == 0 {

		return &candidateSigningReport, nil
	}

	signedCount, err := this.countSignedCandidates(campaignUUID)

	if err != nil {

		return nil, err
	}

	candidateSigningReport.SignedCount = signedCount

	return &candidateSigningReport, nil
}

func (this *CandidateSigningReportService) CountCampaignCandidates(
	campaignUUID_str string,
) (int64, error) {

	campaignUUID, err := uuid.Parse(campaignUUID_str)

	if err != nil {

		return 0, err
	}

	return this.countCampaignCandidates(campaignUUID)
}

func (this *CandidateSigningReportService) countCampaignCandidates(
	campaignUUID uuid.UUID,
) (int64, error) {

	data, err := repository.Aggregate[valueObject.CampaignCandidateCount](
		this.CampaignRepository.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "candidates"},
						{"localField", "uuid"},
						{"foreignField", "campaignUUID"},
						{"as", "candidates"},
					},
				},
			},
			bson.D{
				{
					"$match", bson.D{
						{"uuid", campaignUUID},
					},
				},
			},
			bson.D{{"$project", bson.D{{"totalCount", bson.D{{"$size", "$candidates"}}}}}},
		},
		context.TODO(),
	)

	if err != nil {

		return 0, err
	}

	if len(data) == 0 {

		return 0, common.ERR_HTTP_NOT_FOUND
	}

	return data[0].TotalCount, nil
}

func (this *CandidateSigningReportService) CountSignedCandidates(
	campaignUUID_str string,
) (int64, error) {

	campaignUUID, err := uuid.Parse(campaignUUID_str)

	if err != nil {

		return 0, err
	}

	return this.countSignedCandidates(campaignUUID)
}

func (this *CandidateSigningReportService) countSignedCandidates(
	campaignUUID uuid.UUID,
) (int64, error) {

	data, err := repository.Aggregate[valueObject.CampaignSignedCandidateCount](
		this.CandidateRepository.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "candidateSigningCommits"},
						{"localField", "uuid"},
						{"foreignField", "candidateUUID"},
						{"as", "commits"},
					},
				},
			},
			bson.D{
				{"$unwind",
					bson.D{
						{"path", "$commits"},
						{"includeArrayIndex", "index"},
					},
				},
			},
			bson.D{
				{"$group",
					bson.D{
						{"_id",
							bson.D{
								{"campaignUUID", "$campaignUUID"},
								{"index", "$index"},
							},
						},
						{"signedCount", bson.D{{"$count", bson.D{}}}},
					},
				},
			},
			bson.D{
				{
					"$match", bson.D{
						{"_id.campaignUUID", campaignUUID},
						{"_id.index", 0},
					},
				},
			},
		},
		context.TODO(),
	)

	if err != nil {

		return 0, err
	}

	if len(data) == 0 {

		return 0, nil
	}

	return data[0].SignedCount, nil
}
