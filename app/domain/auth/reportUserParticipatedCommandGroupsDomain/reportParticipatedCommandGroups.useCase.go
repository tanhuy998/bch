package reportUserParticipatedCommandGroupsDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ReportParticipatedCommandGroupsUseCase struct {
		usecasePort.UseCase[requestPresenter.ReportParticipatedGroups, responsePresenter.ReportParticipatedGroups]
		UserReppo                       repository.IUser
		ReportParticipatedCommandGroups authServicePort.IReportParticipatedCommandGroups
	}
)

func (this *ReportParticipatedCommandGroupsUseCase) Execute(
	input *requestPresenter.ReportParticipatedGroups,
) (*responsePresenter.ReportParticipatedGroups, error) {

	if !input.IsValidTenantUUID() {

		return nil, common.ERR_UNAUTHORIZED
	}

	err := this.validateAuthority(input)

	if err != nil {

		return nil, err
	}

	executionCtx := libCommon.Ternary[context.Context](
		input.GetAuthority().IsTenantAgent(),
		input.GetContext(),
		&domain_context{input.GetContext()},
	)

	report, err := this.ReportParticipatedCommandGroups.Serve(
		input.GetTenantUUID(), *input.UserUUID, executionCtx,
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	if report == nil {

		return nil, this.ErrorWithContext(
			input, common.ERR_UNAUTHORIZED,
		)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = report

	return output, nil
}

/*
Check the authority of access token that passed along with the requested
in order to do specific operation.
If the access token authority is tenant agent, requests to all users that
belonged to the tenant with be approved.
Otherwise, Only access token's authority who is command group members that has role "COMMANDER"
corresponding to the group which the requested user has participated is approved.
*/
func (this *ReportParticipatedCommandGroupsUseCase) validateAuthority(
	input *requestPresenter.ReportParticipatedGroups,
) error {

	auth := input.GetAuthority()

	if auth.IsTenantAgent() || *input.UserUUID == auth.GetUserUUID() {
		// when authority is tenant agent, let rest of validation
		// to domain service
		return nil
	}

	authLeadingCommandGroupsCriterias := bson.A{}

	for _, v := range auth.GetParticipatedGroups() {
		// get groups that the authority has role "COMMANDER"
		if !v.HasRoles("COMMANDER") {

			continue
		}
		// compact the command group uuid as criteria for the query
		authLeadingCommandGroupsCriterias = append(
			authLeadingCommandGroupsCriterias,
			bson.D{
				{"commandGroupUUID", v.GetCommandGroupUUID()},
			},
		)
	}

	pipeline := mongo.Pipeline{
		bson.D{
			{
				"$match", bson.D{
					{"uuid", input.UserUUID},
				},
			},
		},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "commandGroupUsers"},
					{"localField", "uuid"},
					{"foreignField", "userUUID"},
					{"as", "groups"},
					{
						"pipeline", mongo.Pipeline{
							bson.D{
								{
									"$match", bson.D{
										{"$or", authLeadingCommandGroupsCriterias},
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
				"$set", bson.D{
					{
						"isCommandGroupMember", bson.D{
							{
								"$cond", bson.D{
									{
										"if", bson.D{
											{
												"$and", bson.A{
													bson.D{{"$isArray", "$groups"}},
													bson.D{
														{
															"$ne", bson.A{
																"$groups", bson.A{},
															},
														},
													},
												},
											},
										},
									},
									{"then", true},
									{"else", false},
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{
				"$project", bson.D{
					{"uuid", 1},
					{"tenantUUID", 1},
					{"isCommandGroupMember", 1},
					{"groups", 1},
				},
			},
		},
	}

	type QeuryResult struct {
		TenantUUID           uuid.UUID                 `bson:"tenantUUID"`
		IsCommandGroupMember bool                      `bson:"isCommandGroupMember"`
		Groups               []*model.CommandGroupUser `bson:"groups"`
	}

	switch requestedUser, err := repository.AggregateOne[QeuryResult](this.UserReppo.GetCollection(), pipeline, input.GetContext()); {
	case err != nil:
		return err
	case requestedUser == nil:
		return errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("user not found"))
	case requestedUser.TenantUUID != input.GetTenantUUID():
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("requested user not in tenant"))
	case !requestedUser.IsCommandGroupMember:
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("the requested user is not participated one of the command group you are leading"))
	default:
	}

	return nil
}
