package addUserToCommandGroupDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/repository"
	"errors"
	"fmt"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IAddUserToCommandGroup interface {
		Execute(
			input *requestPresenter.AddUserToCommandGroupRequest,
			output *responsePresenter.AddUserToCommandGroupResponse,
		) (mvc.Result, error)
	}

	AddUserToCommandGroupUseCase struct {
		usecasePort.UseCase[requestPresenter.AddUserToCommandGroupRequest, responsePresenter.AddUserToCommandGroupResponse]
		AddUserToCommandGroupService authServicePort.IAddUserToCommandGroup
		CommandGroupUserRepo         repository.ICommandGroupUser
	}
)

func (this *AddUserToCommandGroupUseCase) Execute(
	input *requestPresenter.AddUserToCommandGroupRequest,
) (*responsePresenter.AddUserToCommandGroupResponse, error) {

	if !input.IsValidTenantUUID() {

		return nil, errors.Join(common.ERR_BAD_REQUEST, fmt.Errorf("invalid tenant"))
	}

	dataModel := &model.CommandGroupUser{
		UserUUID:         input.UserUUID,
		CommandGroupUUID: input.GroupUUID,
		CreatedBy:        libCommon.PointerPrimitive(input.GetAuthority().GetUserUUID()),
	}

	err := this.AddUserToCommandGroupService.Serve(input.GetTenantUUID(), dataModel, input.GetContext())

	if err != nil {

		return nil, err
	}

	// ret, err := this.CommandGroupUserRepo.Find(
	// 	bson.D{},
	// 	context.TODO(),
	// )

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	//output.Data = ret

	return output, nil
}
