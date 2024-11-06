package getTenantUsersDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	authServicePort "app/port/auth"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	GetTenantUserUseCase struct {
		usecasePort.UseCase[requestPresenter.GetTenantUsers, responsePresenter.GetTenantUsers[model.User]]
		GetTenantUsersService authServicePort.IGetTenantUsers[model.User]
	}
)

func (this *GetTenantUserUseCase) Execute(
	input *requestPresenter.GetTenantUsers,
) (output *responsePresenter.GetTenantUsers[model.User], err error) {

	if !input.IsValidTenantUUID() {

		return nil, this.ErrorWithContext(
			input, errors.Join(
				common.ERR_UNAUTHORIZED,
				fmt.Errorf("no match tenant access token"),
			),
		)
	}

	output = this.GenerateOutput()
	output.Data, err = this.GetTenantUsersService.Serve(
		input.GetTenantUUID(),
		input.PageNumber,
		input.PageSize,
		libCommon.Ternary[*primitive.ObjectID](
			input.HasCursor(), libCommon.PointerPrimitive(input.GetCursor()), nil,
		),
		input.IsPrevious(),
		input.GetContext(),
	)

	if err != nil {

		return nil, this.ErrorWithContext(
			input, err,
		)
	}

	return output, nil
}
