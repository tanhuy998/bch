package genericUseCase

import (
	"app/internal/common"
	"app/model"
	"app/repository"
	"app/valueObject/requestInput"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MongodAuthDomainUseCase[Input_T requestInput.ITenantDomainInput] struct {
		CommandGroupUserRepo repository.ICommandGroupUser
	}
)

func (this *MongodAuthDomainUseCase[Input_T]) CheckUserJoinedAssignment(
	input Input_T, assignmentUUID uuid.UUID,
) error {

	auth := input.GetAuthority()

	if auth == nil {

		return common.ERR_UNAUTHORIZED
	}

	if auth.IsTenantAgent() {

		return nil
	}

	query := mongo.Pipeline{
		bson.D{
			{
				"$match", bson.D{
					{"userUUID", auth.GetUserUUID()},
					{"tenantUUID", input.GetTenantUUID()},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "assignmentGroupMembers"},
					{"localField", "uuid"},
					{"foreignField", "commandGroupUserUUID"},
					{"as", "assignmentGroupMembers"},
					{"pipeline",
						mongo.Pipeline{
							bson.D{
								{
									"$match", bson.D{
										{"tenantUUID", input.GetTenantUUID()},
									},
								},
							},
							bson.D{
								{"$lookup",
									bson.D{
										{"from", "assignmentGroups"},
										{"localField", "assignmentGroupUUID"},
										{"foreignField", "uuid"},
										{"as", "assignmentGroups"},
										{
											"pipeline", mongo.Pipeline{
												bson.D{
													{
														"$match", bson.D{
															{"tenantUUID", input.GetTenantUUID()},
															{"assignmentUUID", assignmentUUID},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{
				"$match", bson.D{
					{
						"assignmentGroupMembers", bson.D{
							{"$ne", bson.A{}},
						},
					},
				},
			},
		},
		bson.D{
			{"$unwind", "$assignmentGroupMembers"},
		},
		bson.D{
			{"$unwind", "$assignmentGroups"},
		},
		//bson.D{{"$replaceWith", bson.D{{"newWith", bson.D{}}}}},
	}

	type T struct {
		AssignmentGroup model.AssignmentGroup `bson:"assignmentGroups"`
	}

	switch ret, err := repository.AggregateOne[T](this.CommandGroupUserRepo.GetCollection(), query, input.GetContext()); {
	case err != nil:
		return err
	case ret == nil:
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("the current user cannot access the resource"))
	default:
		fmt.Println(ret.AssignmentGroup.UUID)
	}

	return nil
}
