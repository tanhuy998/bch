package getUserParticipatedCommandGroupDomain

import (
	"app/src/repository"
	"app/src/valueObject"
	"context"

	"app/src/model"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IGetParticipatedCommandGroups interface {
		Serve(userUUID string) (*valueObject.ParticipatedCommandGroupReport, error)
	}

	GetParticipatedCommandGroupsService struct {
		CommandGroupUserRepo repository.ICommandGroupUser
	}
)

func (this *GetParticipatedCommandGroupsService) Serve(userUUID_str string) (*valueObject.ParticipatedCommandGroupReport, error) {

	userUUID, err := uuid.Parse(userUUID_str)

	if err != nil {

		return nil, err
	}

	res, err := repository.Aggregate[valueObject.ParticipatedCommandGroupDetail](
		this.CommandGroupUserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"userUUID", userUUID},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroups"},
						{"localField", "commandGroupUUID"},
						{"foreignField", "uuid"},
						{"as", "commandGroups"},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroupUserRoles"},
						{"localField", "uuid"},
						{"foreignField", "commandGroupUUID"},
						{"as", "roles"},
					},
				},
			},
			bson.D{{"$unwind", bson.D{{"path", "$commandGroups"}}}},
			bson.D{
				{"$unwind",
					bson.D{
						{"path", "$roles"},
						{"preserveNullAndEmptyArrays", true},
					},
				},
			},
			bson.D{
				{"$set",
					bson.D{
						{"commandGroup", "$commandGroups"},
						{"role", "$roles"},
					},
				},
			},
			bson.D{
				{"$project",
					bson.D{
						{"commandGroup", 1},
						{"role", 1},
					},
				},
			},
		},
		context.TODO(),
	)

	if err != nil {

		return nil, err
	}

	report := &valueObject.ParticipatedCommandGroupReport{
		UserUUID: userUUID,
	}

	if len(res) > 0 {

		report.Details = res
	}

	return report, nil
}

func (this *GetParticipatedCommandGroupsService) SearchAndRetrieveByModel(
	searchModel *model.User, ctx context.Context,
) (*valueObject.ParticipatedCommandGroupReport, error) {

	res, err := repository.Aggregate[valueObject.ParticipatedCommandGroupDetail](
		this.CommandGroupUserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", searchModel,
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroups"},
						{"localField", "commandGroupUUID"},
						{"foreignField", "uuid"},
						{"as", "commandGroups"},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroupUserRoles"},
						{"localField", "uuid"},
						{"foreignField", "commandGroupUUID"},
						{"as", "roles"},
					},
				},
			},
			bson.D{{"$unwind", bson.D{{"path", "$commandGroups"}}}},
			bson.D{
				{"$unwind",
					bson.D{
						{"path", "$roles"},
						{"preserveNullAndEmptyArrays", true},
					},
				},
			},
			bson.D{
				{"$set",
					bson.D{
						{"commandGroup", "$commandGroups"},
						{"role", "$roles"},
					},
				},
			},
			bson.D{
				{"$project",
					bson.D{
						{"commandGroup", 1},
						{"role", 1},
					},
				},
			},
		},
		context.TODO(),
	)

	if err != nil {

		return nil, err
	}

	if len(res) == 0 {

		return nil, nil
	}

	report := &valueObject.ParticipatedCommandGroupReport{
		UserUUID: *searchModel.UUID,
		Details:  res,
	}

	return report, nil
}
