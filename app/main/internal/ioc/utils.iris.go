package ioc

import (
	"app/infrastructure/http/common"
	"log"

	irisIoc "app/internal/lib/iris/ioc"

	accessLogServicePort "app/port/accessLog"
	actionResultServicePort "app/port/actionResult"
	"app/port/loggerServicePort"

	passwordServicePort "app/port/passwordService"
	"app/port/responsePresetPort"
	actionResultService "app/service/actionResult"
	irisAccessLoggerService "app/service/irisAccessLogger"
	passwordService "app/service/password"
	"app/service/responsePresetService"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"

	internalLog "app/internal/log"

	restfullCommon "app/infrastructure/restfull/common"
	restfullMiddleware "app/infrastructure/restfull/common/middleware"
)

func RegisterUtilServices(container *hero.Container) {

	container.Register(log.Default()).Explicitly()
	irisIoc.BindDependency[context.Validator, validator.Validate](container, validator.New())

	irisIoc.BindDependency[actionResultServicePort.IActionResult, actionResultService.ResponseResultService](container, nil)
	irisIoc.BindDependency[responsePresetPort.IResponsePreset, responsePresetService.ResponsePresetService](container, nil)
	irisIoc.BindDependency[passwordServicePort.IPassword, passwordService.PasswordService](container, nil)
	irisIoc.BindDependency[common.IMiddlewareErrorHandler, common.ErrorHandler](container, nil)
	irisIoc.BindDependency[restfullMiddleware.IErrorHandler, restfullCommon.ErrorHandler](container, nil)

	irisIoc.BindDependency[loggerServicePort.ILogger, internalLog.LogBroker](container, nil)

	/*
		access logger must be initialized before repositories in order to trace db query
	*/
	irisIoc.BindDependency[
		accessLogServicePort.IAccessLogger,
		irisAccessLoggerService.IrisAccessLoggerService,
	](container, nil)
	// container.Register(new(common.Controller)).Explicitly().EnableStructDependents()
}
