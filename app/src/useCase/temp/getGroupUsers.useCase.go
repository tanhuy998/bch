package usecase

import (
	authServiceAdapter "app/adapter/auth"
	actionResultService "app/service/actionResult"
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	requestPresenter "app/src/presenter/request"
	responsePresenter "app/src/presenter/response"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetGroupUsers interface {
		Execute(
			input *requestPresenter.GetGroupUsersRequest,
			output *responsePresenter.GetGroupUsersResponse,
		) (mvc.Result, error)
	}

	GetGroupUsersUseCase struct {
		GetCommandGroupUsersService authServiceAdapter.IGetCommandGroupUsers //authService.IGetCommandGroupUsers
		ActionResult                actionResultService.IActionResult
	}
)

func (this *GetGroupUsersUseCase) Execute(
	input *requestPresenter.GetGroupUsersRequest,
	output *responsePresenter.GetGroupUsersResponse,
) (mvc.Result, error) {

	groupUUID, err := uuid.Parse(input.GroupUUID)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	searchModel := &model.CommandGroup{
		UUID:       libCommon.PointerPrimitive(groupUUID),
		TenantUUID: libCommon.PointerPrimitive(input.GetAuthority().GetTenantUUID()),
	}

	data, err := this.GetCommandGroupUsersService.SearchAndRetrieveByModel(searchModel, input.GetContext()) //this.GetCommandGroupUsersService.Serve(input.GroupUUID)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Data = data

	return this.ActionResult.ServeResponse(output)
}
