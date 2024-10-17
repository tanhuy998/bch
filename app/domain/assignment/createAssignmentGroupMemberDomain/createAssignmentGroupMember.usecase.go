package createAssignmentGroupMemberDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	accessTokenServicePort "app/port/accessToken"

	assignmentServicePort "app/port/assignment"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/repository"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	AuthData = accessTokenServicePort.IAccessTokenAuthData

	CreateAssignmentGroupMemberUseCase struct {
		usecasePort.UseCase[requestPresenter.CreateAssignmentGroupMember, responsePresenter.CreateAssignmentGroupMemeber]
		AssignmentGroupRepo                repository.IAssignmentGroup
		CommandGroupUserRepo               repository.ICommandGroupUser
		CreateAssignmentGroupMemberService assignmentServicePort.ICreateAssignmentGroupMember
	}
)

func (this *CreateAssignmentGroupMemberUseCase) Execute(
	input *requestPresenter.CreateAssignmentGroupMember,
) (*responsePresenter.CreateAssignmentGroupMemeber, error) {

	if !input.IsValidTenantUUID() {

		return nil, this.ErrorWithContext(
			input, common.ERR_UNAUTHORIZED,
		)
	}

	dataModel, err := this.validateCommandGroupOwnnershipAndGenerateModel(input)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	data, err := this.CreateAssignmentGroupMemberService.Serve(
		input.GetTenantUUID(),
		*input.AssignmentGroupUUID,
		dataModel,
		&domain_context{input.GetContext()},
	)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	output := this.GenerateOutput()
	output.Message = "success"
	output.Data = data

	return output, nil
}

func (this *CreateAssignmentGroupMemberUseCase) validateCommandGroupOwnnershipAndGenerateModel(
	input *requestPresenter.CreateAssignmentGroupMember,
) ([]*model.AssignmentGroupMember, error) {

	auth := input.GetAuthority()

	if auth == nil {

		return nil, common.ERR_UNAUTHORIZED
	}

	switch existingAssignmentGroup, err := this.AssignmentGroupRepo.FindOneByUUID(*input.AssignmentGroupUUID, input.GetContext()); {
	case err != nil:
		return nil, err
	case existingAssignmentGroup == nil:
		return nil, errors.Join(
			common.ERR_NOT_FOUND, fmt.Errorf("assignment group not found"),
		)
	case auth.IsTenantAgent(): // if user is tenant agent, full authority on assignment group
		return this.generateDataModels(input, *existingAssignmentGroup.CommandGroupUUID)
	case *existingAssignmentGroup.TenantUUID != auth.GetTenantUUID():
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("command group not in tenant"))
	case !auth.QueryCommandGroup(*existingAssignmentGroup.CommandGroupUUID).HasRoles("COMMANDER").Done():
		return nil, errors.Join(
			common.ERR_FORBIDEN,
			fmt.Errorf("the current user is not leader of the command group that assigned with the requested assignment group"),
		)
	default:
		return this.generateDataModels(input, *existingAssignmentGroup.CommandGroupUUID)
	}
}

func (this *CreateAssignmentGroupMemberUseCase) generateDataModels(
	input *requestPresenter.CreateAssignmentGroupMember,
	commandGroupUUID uuid.UUID,
) ([]*model.AssignmentGroupMember, error) {

	if len(input.ComandGroupUserUUIDList) == 0 {

		return nil, errors.Join(
			common.ERR_BAD_REQUEST, fmt.Errorf("empty data given"),
		)
	}

	auth := input.GetAuthority()
	ret := make([]*model.AssignmentGroupMember, len(input.ComandGroupUserUUIDList))
	inputCommandGroupUserUUIDs := make(bson.A, len(input.ComandGroupUserUUIDList))

	for i, v := range input.ComandGroupUserUUIDList {

		inputCommandGroupUserUUIDs[i] = v

		ret[i] = &model.AssignmentGroupMember{
			CreatedBy:            libCommon.PointerPrimitive(auth.GetUserUUID()),
			CommandGroupUserUUID: &v,
		}
	}

	cUsers, err := this.CommandGroupUserRepo.FindMany(
		bson.D{
			{
				"uuid", bson.D{
					{"$in", inputCommandGroupUserUUIDs},
				},
			},
			{"tenantUUID", input.GetTenantUUID()},
		},
		input.GetContext(),
	)

	switch {
	case err != nil:
		return nil, err
	case len(cUsers) != len(input.ComandGroupUserUUIDList):

		return nil, errors.Join(
			common.ERR_FORBIDEN, fmt.Errorf("(tenant agent error) any of requested users is not in the tenant"),
		)
	}

	if auth.IsTenantAgent() {

		return ret, nil
	}

	for _, v := range cUsers {

		if *v.CommandGroupUUID != commandGroupUUID {

			return nil, errors.Join(
				common.ERR_FORBIDEN, fmt.Errorf("any of requested users is not participated in the group whose current user is leading"),
			)
		}
	}

	return ret, nil
}
