package common

import (
	"app/internal/common"
	"app/internal/responseOutput"
	actionResultServicePort "app/port/actionResult"
	contextHolderPort "app/port/contextHolder"
	loggerPort "app/port/logger"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
)

type (
	default_response struct {
		Message string `json:"message,omitempty"`
	}

	IController interface {
		BindDependencies(container *hero.Container) IController
	}

	Controller struct {
		//AccessLogger accessLogServicePort.IAccessLogger
		ActionResult actionResultServicePort.IActionResult
		ErrorLogger  loggerPort.ErrorLogger
		ErrorHandler
	}
)

func (this *Controller) ResultOf(output any, usecaseError error) (mvc.Result, error) {

	if usecaseError != nil {

		return this.handleError(usecaseError)
	}

	if output == nil {

		return this.ActionResult.Prepare().SetCode(http.StatusNoContent).Done()
	}

	if v, ok := output.(responseOutput.IHTTPStatusResponse); ok {

		raw, err := json.Marshal(output)

		if err != nil {

			return this.hanleInternalError(err)
		}

		return this.ActionResult.Prepare().
			SetCode(
				v.GetHTTPStatus(),
			).
			SetContent(raw).
			Done()
	}

	if v, ok := output.(responseOutput.INoContentOutput); ok && v.IsNotContent() {

		return this.ActionResult.Prepare().SetCode(http.StatusNoContent).Done()
	}

	if v, ok := output.(responseOutput.ICreatedOutput); ok && v.IsCreatedStatus() {

		raw, err := json.Marshal(output)

		if err != nil {

			return this.hanleInternalError(err)
		}

		return this.ActionResult.Prepare().SetCode(http.StatusCreated).SetContent(raw).Done()
	}

	if v, ok := output.(responseOutput.IAcceptedOuput); ok && v.IsAccepptedStatus() {

		raw, err := json.Marshal(output)

		if err != nil {

			return this.hanleInternalError(err)
		}

		return this.ActionResult.Prepare().SetCode(http.StatusAccepted).SetContent(raw).Done()
	}

	return this.ActionResult.ServeResponse(output)
}

func (this *Controller) handleError(err error) (mvc.Result, error) {

	// if errors.Is(err, common.ERR_INTERNAL) {

	// 	return this.hanleInternalError(err)
	// }

	// res := this.ActionResult.Prepare()

	// switch {
	// case errors.Is(err, common.ERR_NOT_FOUND):
	// 	res.SetCode(http.StatusNotFound) // 404
	// case errors.Is(err, common.ERR_UNAUTHORIZED):
	// 	res.SetCode(http.StatusUnauthorized) // 401
	// case errors.Is(err, common.ERR_FORBIDEN):
	// 	res.SetCode(http.StatusForbidden) // 403
	// case errors.Is(err, common.ERR_CONFLICT):
	// 	res.SetCode(http.StatusConflict) // 409
	// default:
	// 	res.SetCode(http.StatusBadRequest) // 400
	// }

	// resObj := default_response{
	// 	Message: err.Error(),
	// }

	// raw, _ := json.Marshal(resObj)

	// res.SetContent(raw)

	res := this.HandleError(err)

	if errors.Is(err, common.ERR_INTERNAL) {

		this.logError(err)
	}

	return res, nil
}

func (this *Controller) hanleInternalError(err error) (mvc.Result, error) {

	res := this.ActionResult.Prepare()

	resObj := default_response{
		Message: "internal error",
	}
	fmt.Println("handle error")
	this.logError(err)

	raw, _ := json.Marshal(resObj)

	return res.SetCode(http.StatusInternalServerError).SetContent(raw).Done()
}

func (this *Controller) logError(err error) {

	errOutput, ok := any(err).(contextHolderPort.IContextHolder)

	if !ok {

		return
	}

	ctx := errOutput.GetContext()

	if ctx == nil {

		return
	}

	this.AccessLogger.PushError(ctx, err)
}
