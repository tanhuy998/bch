package common

import (
	"app/internal/common"
	accessLogServicePort "app/port/accessLog"
	actionResultServicePort "app/port/actionResult"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

type ()

type (
	IMiddlewareErrorHandler interface {
		Handle(iris.Context, error)
	}

	ErrorHandler struct {
		AccessLogger accessLogServicePort.IAccessLogger
		ActionResult actionResultServicePort.IActionResult
	}
)

func (this *ErrorHandler) Handle(ctx iris.Context, err error) {

	res := this.HandleError(err)

	if errors.Is(err, common.ERR_INTERNAL) {

		this.AccessLogger.PushError(ctx, err)
	}

	res.Dispatch(ctx)
}

func (this *ErrorHandler) HandleError(err error) hero.Result {

	res := this.ActionResult.Prepare()

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

	resObj := default_response{}

	if errors.Is(err, common.ERR_INTERNAL) {

		resObj.Message = "internal error"
	} else {

		resObj.Message = err.Error()
	}

	raw, _ := json.Marshal(resObj)

	res.SetContent(raw)

	return res
}
