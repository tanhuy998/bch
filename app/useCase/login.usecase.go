package usecase

import (
	authServiceAdapter "app/adapter/auth"
	refreshTokenServicePort "app/adapter/refreshToken"
	refreshTokenIdServicePort "app/adapter/refreshTokenidServicePort"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/bootstrap"
	actionResultService "app/service/actionResult"
	"errors"
	"net/http"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
)

var (
	ERR_NIL_CONTEXT = errors.New("login usecase error: nil context")
)

type (
	ILogIn interface {
		Execute(*requestPresenter.LoginRequest, *responsePresenter.LoginResponse) (mvc.Result, error)
	}

	LogInUseCase struct {
		LogInService            authServiceAdapter.ILogIn
		ActionResult            actionResultService.IActionResult
		RefreshTokenManipulator refreshTokenServicePort.IRefreshTokenManipulator
	}
)

func (this *LogInUseCase) Execute(
	input *requestPresenter.LoginRequest,
	output *responsePresenter.LoginResponse,
) (mvc.Result, error) {

	reqContext := input.GetContext()

	if reqContext == nil {

		this.ActionResult.ServeErrorResponse(ERR_NIL_CONTEXT)
	}

	accessToken, refreshToken, err := this.LogInService.Serve(input.Data.Username, input.Data.Password, reqContext)

	if err != nil {

		return this.ActionResult.ServeErrorResponse(err)
	}

	output.Data.AccessToken = accessToken
	output.Message = "success"

	for _, hostname := range bootstrap.GetHostNames() {

		if hostname == "*" {

			continue
		}

		reqContext.SetCookieKV(
			refreshTokenIdServicePort.REFRESH_TOKEN_COOKIE,
			refreshToken,
			context.CookieExpires(this.RefreshTokenManipulator.DefaultExpireDuration()),
			context.CookieDomain(hostname),
			context.CookiePath("/refresh"),
			context.CookieHTTPOnly(true),
			context.CookieSameSite(http.SameSiteStrictMode),
		)
	}

	return this.ActionResult.ServeResponse(output)
}
