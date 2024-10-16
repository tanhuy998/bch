package getSingleAssignmentDomain

import (
	"app/internal/common"
	libError "app/internal/lib/error"
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
	GetSingleAssignmentService struct {
		AssignmentRepo repository.IAssignment
	}
)

func (this *GetSingleAssignmentService) Serve(
	tenantUUID uuid.UUID, uuid uuid.UUID, ctx context.Context,
) (*model.Assignment, error) {

	ret, err := repository.AggregateOne[model.Assignment](
		this.AssignmentRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"tenantUUID", tenantUUID},
						{"uuid", uuid},
					},
				},
			},
			bson.D{
				{
					"$lookup", bson.D{
						{"from", "users"},
						{"localField", "createdBy"},
						{"foreignField", "uuid"},
						{"as", "createdUser"},
					},
				},
			},
			bson.D{
				{
					"$set", bson.D{
						{
							"createdUser", bson.D{
								{"$arrayElemAt", bson.A{"$createdUser", 0}},
							},
						},
					},
				},
			},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	if ret == nil {

		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("assignment not found"))
	}

	if ret.TenantUUID == nil {

		return nil, libError.NewInternal(fmt.Errorf("wrong data"))
	}

	if *ret.TenantUUID != tenantUUID {

		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("assignment not in tenant"))
	}

	return ret, nil
}

func (this *GetSingleAssignmentService) Search(
	model *model.Assignment, ctx context.Context,
) (*model.Assignment, error) {

	ret, err := this.AssignmentRepo.Find(
		bson.D{
			{"title", model.Title},
			{"tenantUUID", model.TenantUUID},
		},
		ctx,
	)

	if err != nil {

		return nil, errors.Join(common.ERR_INTERNAL, err)
	}

	return ret, nil
}
