package addUserToCommandGroupDomain

import (
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"app/repository"
	"context"

	"github.com/kataras/iris/v12/mvc"
	"go.mongodb.org/mongo-driver/bson"
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

	err := this.AddUserToCommandGroupService.Serve(*input.GroupUUID, *input.UserUUID, input.GetContext())

	if err != nil {

		return nil, err
	}

	ret, err := this.CommandGroupUserRepo.Find(
		bson.D{},
		context.TODO(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(input, err)
	}

	output := this.GenerateOutput()

	output.Message = "success"
	output.Data = ret

	return output, nil
}
