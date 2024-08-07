package signingService

import (
	"app/domain/model"
	libCommon "app/lib/common"
	"app/repository"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IGetCampaignSignedCandidates interface {
		Serve(
			campaignUUID string,
			pivotObjID_str string,
			limit int,
			isPrevDir bool,
		) (*repository.PaginationPack[model.Candidate], error)
	}

	GetCampaignSignedCandidates struct {
		CandidateSigningCommitRepository repository.ICandidateSigningCommit
		CandidateRepository              repository.ICandidateRepository
		SigningInfoRepo                  repository.ICandidateSigningInfo
	}
)

func (this *GetCampaignSignedCandidates) Serve(
	campaignUUID_str string,
	pivotObjID_str string,
	limit int,
	isPrevDir bool,
) (*repository.PaginationPack[model.Candidate], error) {

	var (
		pivotObjID primitive.ObjectID
		err        error
	)

	campaignUUID, err := uuid.Parse(campaignUUID_str)

	if err != nil {

		return nil, err
	}

	if pivotObjID_str == "" {

		pivotObjID = primitive.NilObjectID
		err = nil

	} else {

		pivotObjID, err = primitive.ObjectIDFromHex(pivotObjID_str)
	}

	if err != nil {

		return nil, err
	}

	pack, err := this.Query(campaignUUID, pivotObjID, limit, isPrevDir)

	return pack, err
}

func (this *GetCampaignSignedCandidates) Query(
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

	pipelineAfterPivot := mongo.Pipeline{
		bson.D{
			{
				"$match", bson.D{
					{"campaignUUID", campaignUUID},
				},
			},
		},
		bson.D{
			{
				"$lookup", bson.D{
					{"from", "candidateSigningInfos"},
					{"localField", "uuid"},
					{"foreignField", "candidateUUID"},
					{"as", "signingInfos"},
				},
			},
		},
		// bson.D{
		// 	{
		// 		"$unwind", "$signingInfos",
		// 	},
		// },
		bson.D{
			{
				"$match", bson.D{
					{
						"signingInfos", bson.D{
							{"$ne", bson.A{}},
						},
					},
				},
			},
		},
		bson.D{
			{
				"$project", bson.D{
					{"signingInfos", 0},
				},
			},
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
		//this.CandidateSigningCommitRepository.GetCollection(),
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
