package getCommandGroupUsersDomain

import (
	"app/domain"
	"app/internal/common"
	libCommon "app/internal/lib/common"
	authServicePort "app/port/auth"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/repository"
	"context"
	"errors"
	"fmt"
)

type (
	GetCommandGroupUsersUseCase struct {
		//usecasePort.UseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse]
		domain.PaginatableUseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse]
		CommandGroupRepo           repository.ICommandGroup
		GetCommandGroupUserService authServicePort.IGetCommandGroupUsers[domain.PaginateCursorType]
	}
)

func (this *GetCommandGroupUsersUseCase) Execute(
	input *requestPresenter.GetGroupUsersRequest,
) (*responsePresenter.GetGroupUsersResponse, error) {

	if !input.IsValidTenantUUID() {

		return nil, common.ERR_UNAUTHORIZED
	}

	err := this.validateAuthority(input)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	executionCtx := libCommon.Ternary[context.Context](
		input.GetAuthority().IsTenantAgent(),
		input.GetContext(),
		&command_group_leader_context{input.GetContext()},
	)

	data, err := this.GetCommandGroupUserService.Serve(
		input.GetTenantUUID(), *input.RequestedGroupUUID, input, executionCtx,
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()
	output.Message = "success"
	output.Data = data

	return output, nil
}

func (this *GetCommandGroupUsersUseCase) validateAuthority(
	input *requestPresenter.GetGroupUsersRequest,
) error {

	if !input.HasAuthority() {

		return common.ERR_FORBIDEN
	}

	auth := input.GetAuthority()

	if auth.IsTenantAgent() {

		return nil
	}

	if !auth.QueryCommandGroup(*input.RequestedGroupUUID).HasRoles("COMMANDER").Done() {

		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("the current user doesn't has authority to access this enpoint"))
	}

	return nil
}
