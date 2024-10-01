package addUserToCommandGroup

import (
	actionResultServicePort "app/src/port/actionResult"
	authServicePort "app/src/port/auth"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"
	"app/src/repository"
	"context"
	"encoding/json"

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
		AddUserToCommandGroupService authServicePort.IAddUserToCommandGroup
		CommandGroupUserRepo         repository.ICommandGroupUser
		ActionResult                 actionResultServicePort.IActionResult
	}
)

func (this *AddUserToCommandGroupUseCase) Execute(
	input *requestPresenter.AddUserToCommandGroupRequest,
	output *responsePresenter.AddUserToCommandGroupResponse,
) (mvc.Result, error) {

	err := this.AddUserToCommandGroupService.Serve(input.GroupUUID, input.UserUUID)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	ret, err := this.CommandGroupUserRepo.Find(
		bson.D{},
		context.TODO(),
	)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Message = "success"
	output.Data = ret

	rawContent, _ := json.Marshal(output)

	return this.ActionResult.Prepare().SetCode(201).SetContent(rawContent), nil
}
