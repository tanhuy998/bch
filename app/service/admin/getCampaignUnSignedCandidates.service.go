package adminService

import (
	libCommon "app/internal/lib/common"
	"app/model"
	"app/port/signingServicePort"
	"app/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IGetCampaignUnSignedCandidates interface {
		Serve(
			campaignUUID string,
			pivotObjID_str string,
			limit int,
			isPrevDir bool,
		) (*repository.PaginationPack[model.Candidate], error)
	}

	GetCampaignUnSignedCandidatesService struct {
		CandidateSigningCommitRepository repository.ICandidateSigningCommit
		CandidateRepository              repository.ICandidateRepository
		GetUnSignedCandidateAdapter      signingServicePort.IGetCampaignUnSignedCandidates
	}
)

func (this *GetCampaignUnSignedCandidatesService) Serve(
	campaignUUID string,
	pivotObjID_str string,
	limit int,
	isPrevDir bool,
) (*repository.PaginationPack[model.Candidate], error) {

	return this.GetUnSignedCandidateAdapter.Serve(campaignUUID, pivotObjID_str, limit, isPrevDir)
}

// func (this *GetCampaignUnSignedCandidatesService) Serve(
// 	campaignUUID_str string,
// 	pivotObjID_str string,
// 	limit int,
// 	isPrevDir bool,
// ) (*repository.PaginationPack[model.Candidate], error) {

// 	var (
// 		pivotObjID primitive.ObjectID
// 		err        error
// 	)

// 	campaignUUID, err := uuid.Parse(campaignUUID_str)

// 	if err != nil {

// 		return nil, err
// 	}

// 	if pivotObjID_str == "" {

// 		pivotObjID = primitive.NilObjectID
// 		err = nil

// 	} else {

// 		pivotObjID, err = primitive.ObjectIDFromHex(pivotObjID_str)
// 	}

// 	if err != nil {

// 		return nil, err
// 	}

// 	pack, err := this.Query(campaignUUID, pivotObjID, limit, isPrevDir)

// 	return pack, err
// }

func (this *GetCampaignUnSignedCandidatesService) Query(
	campaignUUID uuid.UUID,
	pivotObjID primitive.ObjectID,
	limit int,
	isPrevDir bool,
) (*repository.PaginationPack[model.Candidate], error) {

	var sortOrder repository.MongoDBCursorSortOrder = repository.SORT_DESC

	if isPrevDir {

		sortOrder = repository.SORT_ASC
	}

	mainPipeLine := mongo.Pipeline{
		bson.D{
			{
				"$sort", bson.D{
					{"_id", sortOrder},
				},
			},
		},
	}

	// pipelineAfterPivot := mongo.Pipeline{
	// 	bson.D{
	// 		{"$group",
	// 			bson.D{
	// 				{"_id", "$candidateUUID"},
	// 				{"lastCommit", bson.D{{"$last", "$$ROOT"}}},
	// 			},
	// 		},
	// 	},
	// 	bson.D{
	// 		{"$lookup",
	// 			bson.D{
	// 				{"from", "candidates"},
	// 				{"localField", "_id"},
	// 				{"foreignField", "uuid"},
	// 				{"as", "candidates"},
	// 			},
	// 		},
	// 	},
	// 	bson.D{
	// 		{"$unwind", "$candidates"},
	// 	},
	// 	bson.D{
	// 		{"$project",
	// 			bson.D{
	// 				{"lastCommit", 1},
	// 				{"candidate", "$candidates"},
	// 			},
	// 		},
	// 	},
	// 	bson.D{
	// 		{
	// 			"$match", bson.D{
	// 				{"candidate.campaignUUID", campaignUUID},
	// 			},
	// 		},
	// 	},
	// 	bson.D{
	// 		{"$limit", limit},
	// 	},
	// 	bson.D{
	// 		{
	// 			"$replaceRoot", bson.D{
	// 				{
	// 					"newRoot", bson.D{
	// 						{
	// 							"$mergeObjects", bson.A{
	// 								"$$ROOT.candidate",
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	bson.D{
	// 		{
	// 			"$project", bson.D{
	// 				{"signingInfo", 0},
	// 			},
	// 		},
	// 	},
	// }

	pipelineAfterPivot := mongo.Pipeline{
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
			{"$match",
				bson.D{
					{
						"campaignUUID", bson.D{
							{"$eq", campaignUUID},
						},
					},
					{"commits",
						bson.D{
							{"$type", "array"},
							{"$eq", bson.A{}},
						},
					},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"commits", 0},
					{"signingInfo", 0},
					{"campaignUUID", 0},
					{"version", 0},
				},
			},
		},
		bson.D{
			{"$limit", limit},
		},
	}

	if !pivotObjID.IsZero() {

		mainPipeLine = append(
			mainPipeLine,
			bson.D{
				{"$match", repository.PrepareObjIDFilterPaginationQuery(pivotObjID, isPrevDir, nil)},
			},
		)
	}

	mainPipeLine = append(mainPipeLine, pipelineAfterPivot...)

	data, err := repository.Aggregate[model.Candidate](
		this.CandidateRepository.GetCollection(),
		mainPipeLine,
		context.TODO(),
	)

	//fmt.Println("len", len(data), *data[0].ObjectID)

	if err != nil {

		return nil, err
	}

	if isPrevDir && len(data) > 0 {

		libCommon.ReverseSlice(data)
	}

	return &repository.PaginationPack[model.Candidate]{
		Data: data,
	}, nil
}
