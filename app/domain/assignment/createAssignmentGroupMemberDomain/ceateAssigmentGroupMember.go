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
	"context"
	"errors"
	"fmt"

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

	dataModel, err := this.validateCommandGroupOwnnershipAndGenerateModel(*input)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	executionContext := libCommon.Ternary[context.Context](
		input.GetAuthority().IsTenantAgent(),
		input.GetContext(),
		&domain_context{input.GetContext()},
	)

	err = this.CreateAssignmentGroupMemberService.Serve(
		input.GetTenantUUID(),
		*input.AssignmentGroupUUID,
		dataModel,
		executionContext,
	)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	output := this.GenerateOutput()
	output.Message = "success"

	return output, nil
}

func (this *CreateAssignmentGroupMemberUseCase) generateDataModels(
	input requestPresenter.CreateAssignmentGroupMember,
	assignmentGroup model.AssignmentGroup,
) ([]*model.AssignmentGroupMember, error) {

	if len(input.Data) == 0 {

		return nil, errors.Join(
			common.ERR_BAD_REQUEST, fmt.Errorf("empty data given"),
		)
	}

	auth := input.GetAuthority()
	ret := make([]*model.AssignmentGroupMember, len(input.Data))
	criterias := make(bson.A, len(input.Data))

	for i, v := range input.Data {

		criterias[i] = bson.D{
			{"userUUID", v},
			{"tenantUUID", input.GetTenantUUID()},
		}

		ret[i] = &model.AssignmentGroupMember{
			CreatedBy:            libCommon.PointerPrimitive(auth.GetUserUUID()),
			CommandGroupUserUUID: &v,
		}
	}

	cUsers, err := this.CommandGroupUserRepo.FindMany(
		bson.D{
			{"$or", criterias},
		},
		input.GetContext(),
	)

	switch {
	case err != nil:
		return nil, err
	case len(cUsers) != len(input.Data):
		return nil, errors.Join(
			common.ERR_FORBIDEN, fmt.Errorf("(tenant agent error) any of requested users is not in the tenant"),
		)
	}

	if auth.IsTenantAgent() {

		return ret, nil
	}

	for _, v := range cUsers {

		if v.CommandGroupUUID != assignmentGroup.CommandGroupUUID {

			return nil, errors.Join(
				common.ERR_FORBIDEN, fmt.Errorf("any of requested users is not participated in the group whose current user is leading"),
			)
		}
	}

	return ret, nil
}

func (this *CreateAssignmentGroupMemberUseCase) validateCommandGroupOwnnershipAndGenerateModel(
	input requestPresenter.CreateAssignmentGroupMember,
) ([]*model.AssignmentGroupMember, error) {

	auth := input.GetAuthority()

	existingAssignmentGroup, err := this.AssignmentGroupRepo.FindOneByUUID(*input.AssignmentGroupUUID, input.GetContext())

	if auth.IsTenantAgent() {
		// if user is tenant agent, full authority on assignment group
		return this.generateDataModels(input, *existingAssignmentGroup)
	}

	switch {
	case err != nil:
		return nil, err
	case existingAssignmentGroup == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("command group not found"))
	case *existingAssignmentGroup.TenantUUID != auth.GetTenantUUID():
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("command group not in tenant"))
	}

	for _, commandGroup := range auth.GetParticipatedGroups() {

		if commandGroup.GetCommandGroupUUID() == existingAssignmentGroup.CommandGroupUUID {
			// check whether or not the user is holding the COMMAND role of the command group
			// that is assigned to the assignment group
			// * Role checking is done by middleware.Auth() on the endpoint, no need no check again
			return this.generateDataModels(input, *existingAssignmentGroup)
		}
	}

	return nil, errors.Join(
		common.ERR_FORBIDEN,
		fmt.Errorf("the current user is not leader of the command group that assigned with the requested assignment group"),
	)
}
