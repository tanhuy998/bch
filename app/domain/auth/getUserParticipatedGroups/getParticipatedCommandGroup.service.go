package getUserParticipatedCommandGroupDomain

import (
	"app/internal/common"
	"app/repository"
	"app/valueObject"
	"context"
	"errors"
	"fmt"

	"app/model"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// IGetParticipatedCommandGroups interface {
	// 	Serve(userUUID string) (*valueObject.ParticipatedCommandGroupReport, error)
	// }

	GetParticipatedCommandGroupsService struct {
		CommandGroupUserRepo repository.ICommandGroupUser
		UserRepo             repository.IUser
	}
)

func (this *GetParticipatedCommandGroupsService) Serve(
	tenantUUID uuid.UUID, userUUID uuid.UUID, ctx context.Context,
) (*valueObject.ParticipatedCommandGroupReport, error) {

	switch existingUser, err := this.UserRepo.FindOneByUUID(userUUID, ctx); {
	case err != nil:
		return nil, err
	case existingUser == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("user not found"))
	case *existingUser.TenantUUID != tenantUUID:
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("user not in tenant"))
	}

	res, err := repository.Aggregate[valueObject.ParticipatedCommandGroupDetail](
		this.CommandGroupUserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"userUUID", userUUID},
						{"tenantUUID", tenantUUID},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroupUserRoles"},
						{"localField", "uuid"},
						{"foreignField", "commandGroupUserUUID"},
						{"as", "roles"},
						{"pipeline",
							bson.A{
								bson.D{
									{"$lookup",
										bson.D{
											{"from", "roles"},
											{"localField", "roleUUID"},
											{"foreignField", "uuid"},
											{"as", "detail"},
										},
									},
								},
								bson.D{{"$unwind", "$detail"}},
								bson.D{{"$replaceWith", "$detail"}},
								bson.D{{
									"$project", bson.D{
										{"name", 1},
										{"uuid", 1},
									},
								}},
							},
						},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroups"},
						{"localField", "commandGroupUUID"},
						{"foreignField", "uuid"},
						{"as", "detail"},
						{"pipeline",
							bson.A{
								bson.D{{"$project", bson.D{{"name", 1}}}},
							},
						},
					},
				},
			},
			bson.D{{"$unwind", "$detail"}},
			bson.D{{"$set", bson.D{{"name", "$detail.name"}}}},
			bson.D{
				{"$project",
					bson.D{
						{"tenantUUID", 0},
						{"userUUID", 0},
						{"uuid", 0},
						{"detail", 0},
					},
				},
			},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	report := &valueObject.ParticipatedCommandGroupReport{
		UserUUID:   userUUID,
		TenantUUID: tenantUUID,
		Details:    res,
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
		ctx,
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
