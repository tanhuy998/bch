package usecase

import (
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	actionResultService "app/service/actionResult"
	authService "app/service/auth"
	"encoding/json"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ICreateCommandGroup interface {
		Execute(
			input *requestPresenter.CreateCommandGroupRequest,
			output *responsePresenter.CreateCommandGroupResponse,
		) (mvc.Result, error)
	}

	CreateCommandGroupUseCase struct {
		CreateCommandGroupService    authService.ICreateCommandGroup
		GetSingleCommandGroupService authService.IGetSingleCommandGroup
		ActionResult                 actionResultService.IActionResult
	}
)

func (this *CreateCommandGroupUseCase) Execute(
	input *requestPresenter.CreateCommandGroupRequest,
	output *responsePresenter.CreateCommandGroupResponse,
) (mvc.Result, error) {

	err := this.CreateCommandGroupService.Serve(input.Data.Name)

	if err != nil {

		return nil, err
	}

	newGroup, err := this.GetSingleCommandGroupService.SearchByName(input.Data.Name)

	if err != nil {

		return nil, err
	}

	output.Message = "success"
	output.Data.UUID = newGroup.UUID

	rawContent, err := json.Marshal(output)

	return this.ActionResult.Prepare().SetCode(201).SetContent(rawContent), nil
}
