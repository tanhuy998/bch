package getCommandGroupUsersDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/repository"
	"context"
	"errors"
	"fmt"
)

type (
	GetCommandGroupUsersUseCase struct {
		usecasePort.UseCase[requestPresenter.GetGroupUsersRequest, responsePresenter.GetGroupUsersResponse]
		CommandGroupRepo           repository.ICommandGroup
		GetCommandGroupUserService authServicePort.IGetCommandGroupUsers
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
		input.GetTenantUUID(), *input.GroupUUID, executionCtx,
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

	if input.IsTenantAgent() {

		return nil
	}

	if !input.QueryCommandGroup(*input.GroupUUID).HasRoles("COMMANDER").Done() {

		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("the current user doesn't has authority to access this enpoint"))
	}

	return nil
}
