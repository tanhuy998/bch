package signingService

import (
	"app/repository"
	"app/valueObject"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ICountSignedCandidate interface {
		Serve(camapignUUID uuid.UUID) (int64, error)
	}

	CountSignedCandidateService struct {
		CandidateRepo repository.ICandidateRepository
	}
)

func (this *CountSignedCandidateService) Serve(campaignUUID uuid.UUID) (int64, error) {

	data, err := repository.Aggregate[valueObject.CampaignSignedCandidateCount](
		this.CandidateRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "candidateSigningInfos"},
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
	/*
		The signing info is now migrated to seperated collection.
		When the aggregation query returns nothing, it doesn't meaning the campaign is unknown
		and there are two scenerios corresponding to when the campaign does not exist or there
		are no candiate signed to that campaign, it depends on the the existence of the campaign.
	*/
	if len(data) == 0 {

		return 0, nil
	}

	return data[0].SignedCount, nil
}
