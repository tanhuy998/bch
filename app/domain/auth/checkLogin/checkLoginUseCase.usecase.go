package checkLoginDomain

import (
	"app/internal/common"
	generalTokenClientServicePort "app/port/generalTokenClient"
	usecasePort "app/port/usecase"
	requestPresenter "app/presenter/request"
	responsePresenter "app/presenter/response"
)

type (
	CheckLoginUseCase struct {
		usecasePort.UseCase[requestPresenter.CheckLogin, responsePresenter.CheckLogin]
		GeneralTokenClient generalTokenClientServicePort.IGeneralTokenClient
	}
)

func (this *CheckLoginUseCase) Execute(
	input *requestPresenter.CheckLogin,
) (*responsePresenter.CheckLogin, error) {

	generalToken, err := this.GeneralTokenClient.Read(input.GetContext())

	if err != nil {

		return nil, err
	}

	if generalToken == nil {

		return nil, this.ErrorWithContext(
			input, common.ERR_UNAUTHORIZED,
		)
	}

	output := this.GenerateOutput()

	return output, nil
}
