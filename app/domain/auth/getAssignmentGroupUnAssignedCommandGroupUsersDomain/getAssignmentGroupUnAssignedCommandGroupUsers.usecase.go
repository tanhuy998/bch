package getAssignmentGroupUnAssignedCommandGroupUsersDomain

import (
	"app/internal/common"
	"app/model"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/repository"
	"errors"
	"fmt"
)

type (
	GetAssignmentGroupUnAssignedCommandGroupUsersUseCase struct {
		usecasePort.UseCase[requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers, responsePresenter.GetAssignmentGroupUnAssignedCommandGroupUsers[model.CommandGroupUser]]
		AssignmentGroupRepo                             repository.IAssignmentGroup
		GetAssignmentUnAssignedCommandGroupUsersService authServicePort.IGetAssignmentGroupUnAssignedCommandGroupUsers[model.CommandGroupUser]
	}
)

func (this *GetAssignmentGroupUnAssignedCommandGroupUsersUseCase) Execute(
	input *requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers,
) (*responsePresenter.GetAssignmentGroupUnAssignedCommandGroupUsers[model.CommandGroupUser], error) {

	switch {
	case !input.IsValidTenantUUID():
		return nil, common.ERR_UNAUTHORIZED
	case input.GetAuthority() == nil:
		return nil, common.ERR_UNAUTHORIZED
	}

	err := this.validateAuthority(input)

	if err != nil {

		return nil, err
	}

	// executionCtx := libCommon.Ternary[context.Context](
	// 	input.IsTenantAgent(),
	// 	input.GetContext(),
	// 	&non_tenant_agent_context{input.GetContext()},
	// )

	var (
		data        []model.CommandGroupUser
		domainInput = mapDomainInput(input)
		auth        = input.GetAuthority()
	)

	if auth.IsTenantAgent() {

		data, err = this.GetAssignmentUnAssignedCommandGroupUsersService.Serve(
			//input.GetTenantUUID(), *input.AssignmentGroupUUID, executionCtx,
			domainInput,
		)
	} else {

		data, err = this.GetAssignmentUnAssignedCommandGroupUsersService.Serve(
			//input.GetTenantUUID(), *input.AssignmentGroupUUID, executionCtx, input.GetUserUUID(),
			domainInput, auth.GetUserUUID(),
		)
	}

	fmt.Println(err)

	if err != nil {

		return nil, err
	}

	output := this.GenerateOutput()
	output.Message = "success"
	output.Data = data

	return output, nil
}

func (this *GetAssignmentGroupUnAssignedCommandGroupUsersUseCase) validateAuthority(
	input *requestPresenter.GetAssignmentGroupUnAssignedCommandGroupUsers,
) error {

	auth := input.GetAuthority()

	if auth.IsTenantAgent() {

		return nil
	}

	switch existingAssignmentGroup, err := this.AssignmentGroupRepo.FindOneByUUID(*input.AssignmentGroupUUID, input.GetContext()); {
	case err != nil:
		return err
	case existingAssignmentGroup == nil:
		return errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("assignment group not found"))
	case *existingAssignmentGroup.TenantUUID != input.GetTenantUUID():
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("assignment group not in tenant"))
	case !auth.QueryCommandGroup(*existingAssignmentGroup.CommandGroupUUID).HasRoles("COMMANDER").Done():
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("the current user is not leader of the requested command group"))
	default:
		return nil
	}
}
