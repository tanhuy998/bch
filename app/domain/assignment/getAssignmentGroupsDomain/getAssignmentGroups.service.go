package getAssignmentGroupsDomain

import (
	"app/internal/common"
	"app/model"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	GetAssignmentGroupsService struct {
		AssignmentGroupRepo repository.IAssignmentGroup
	}
)

func (this *GetAssignmentGroupsService) Serve(
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID, ctx context.Context,
) ([]*model.AssignmentGroup, error) {

	switch {
	case tenantUUID == uuid.Nil:
		return nil, common.ERR_UNAUTHORIZED
	case assignmentUUID == uuid.Nil:
		return nil, errors.Join(common.ERR_BAD_REQUEST, fmt.Errorf("invalid assignment uuid"))
	}

	return repository.Aggregate[model.AssignmentGroup](
		this.AssignmentGroupRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "users"},
						{"localField", "createdBy"},
						{"foreignField", "uuid"},
						{"as", "createdUser"},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroups"},
						{"localField", "commandGroupUUID"},
						{"foreignField", "uuid"},
						{"as", "commandGroup"},
					},
				},
			},
			bson.D{
				{"$set",
					bson.D{
						{"commandGroup",
							bson.D{
								{"$arrayElemAt",
									bson.A{
										"$commandGroup",
										0,
									},
								},
							},
						},
						{"createdUser",
							bson.D{
								{"$arrayElemAt",
									bson.A{
										"$createdUser",
										0,
									},
								},
							},
						},
					},
				},
			},
			bson.D{
				{"$project",
					bson.D{
						{"user",
							bson.D{
								{"password", 0},
								{"secret", 0},
							},
						},
					},
				},
			},
		},
		ctx,
	)
}
