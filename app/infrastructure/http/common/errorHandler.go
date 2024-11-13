package common

import (
	"app/internal/common"
	accessLogServicePort "app/port/accessLog"
	actionResultServicePort "app/port/actionResult"
	contextHolderPort "app/port/contextHolder"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

type ()

type (
	IMiddlewareErrorHandler interface {
		HandleContextError(iris.Context, error)
	}

	ErrorHandler struct {
		AccessLogger accessLogServicePort.IAccessLogger
		ActionResult actionResultServicePort.IActionResult
	}
)

func (this *ErrorHandler) HandleContextError(ctx iris.Context, err error) {

	if err == nil {

		return
	}

	res := this.HandleError(err, ctx)

	if errors.Is(err, common.ERR_INTERNAL) {

		this.AccessLogger.PushError(ctx, err)
	}

	res.Dispatch(ctx)
}

func (this *ErrorHandler) HandleError(err error, ctx context.Context) hero.Result {

	defer func() {

		if errors.Is(err, common.ERR_INTERNAL) {

			this.logError(err, ctx)
		}
	}()

	res := this.ActionResult.Prepare()

	if err == nil {

		return res
	}

	switch {
	case errors.Is(err, common.ERR_INTERNAL):
		res.SetCode(http.StatusInternalServerError)
	case errors.Is(err, common.ERR_NOT_FOUND):
		res.SetCode(http.StatusNotFound) // 404
	case errors.Is(err, common.ERR_UNAUTHORIZED):
		res.SetCode(http.StatusUnauthorized) // 401
	case errors.Is(err, common.ERR_FORBIDEN):
		res.SetCode(http.StatusForbidden) // 403
	case errors.Is(err, common.ERR_CONFLICT):
		res.SetCode(http.StatusConflict) // 409
	default:
		res.SetCode(http.StatusBadRequest) // 400
	}

	resBody := default_response{}

	if errors.Is(err, common.ERR_INTERNAL) {

		resBody.Message = "internal error"
	} else {
		//fmt.Println("-------------", err == nil)
		resBody.Message = err.Error()
	}

	raw, _ := json.Marshal(resBody)

	res.SetContent(raw)

	return res
}

func (this *ErrorHandler) logError(err error, ctx context.Context) {

	if ctx == nil {

		errOutput, ok := any(err).(contextHolderPort.IContextHolder)

		if !ok {

			return
		}

		ctx = errOutput.GetContext()

		if ctx == nil {

			return
		}
	}

	this.AccessLogger.PushError(ctx, err)
}
