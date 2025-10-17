package result

import (
	"app/internal/common"
	"app/internal/errorMap"
	accessLogServicePort "app/port/accessLog"
	actionResultServicePort "app/port/actionResult"
	contextHolderPort "app/port/contextHolder"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

type (
	error_response_body_t struct {
		HttpStatusCode int    `json:status,omitempty`
		ErrorCode      string `json:"code,omitempty"`
		Message        string `json:"message,omitempty"`
	}
)

type (
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

	var httpStatusCode int

	switch {
	case errors.Is(err, common.ERR_INTERNAL):
		res.SetCode(http.StatusInternalServerError)
		httpStatusCode = http.StatusInternalServerError
	case errors.Is(err, common.ERR_NOT_FOUND):
		res.SetCode(http.StatusNotFound) // 404
		httpStatusCode = http.StatusNotFound
	case errors.Is(err, common.ERR_UNAUTHORIZED):
		res.SetCode(http.StatusUnauthorized) // 401
		httpStatusCode = http.StatusUnauthorized
	case errors.Is(err, common.ERR_FORBIDEN):
		res.SetCode(http.StatusForbidden) // 403
		httpStatusCode = http.StatusForbidden
	case errors.Is(err, common.ERR_CONFLICT):
		res.SetCode(http.StatusConflict) // 409
		httpStatusCode = http.StatusConflict
	default:
		res.SetCode(http.StatusBadRequest) // 400
		httpStatusCode = http.StatusBadRequest
	}

	resBody := error_response_body_t{}
	resBody.HttpStatusCode = httpStatusCode

	if errors.Is(err, common.ERR_INTERNAL) {
		resBody.Message = "internal error"
	} else {
		resBody.Message = err.Error()
	}

	var codeErr errorMap.ICodeError

	if errors.As(err, &codeErr) {

		resBody.ErrorCode = fmt.Sprintf("%X", codeErr.GetCode())
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
