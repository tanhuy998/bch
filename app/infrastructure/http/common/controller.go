package common

import (
	"app/internal/common"
	"app/internal/responseOutput"
	actionResultServicePort "app/port/actionResult"
	loggerPort "app/port/logger"
	"encoding/json"
	"errors"
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
		ActionResult actionResultServicePort.IActionResult
		ErrorLogger  loggerPort.ErrorLogger
	}
)

func (this *Controller) ResultOf(output any, usecaseError error) (mvc.Result, error) {

	if usecaseError != nil {

		return this.handleError(usecaseError)
	}

	if output == nil {

		return this.ActionResult.Prepare().SetCode(200).Done()
	}

	if v, ok := output.(responseOutput.INoContentOutput); ok && v.IsNotContent() {

		return this.ActionResult.Prepare().SetCode(200).Done()
	}

	if v, ok := output.(responseOutput.ICreatedOutput); ok && v.IsCreatedStatus() {

		raw, err := json.Marshal(output)

		if err != nil {

			return this.hanleInternalError(err)
		}

		return this.ActionResult.Prepare().SetCode(http.StatusCreated).SetContent(raw).Done()
	}

	return this.ActionResult.ServeResponse(output)
}

func (this *Controller) handleError(err error) (mvc.Result, error) {

	if errors.Is(err, common.ERR_INTERNAL) {

		return this.hanleInternalError(err)
	}

	res := this.ActionResult.Prepare()

	switch {
	case errors.Is(err, common.ERR_NOT_FOUND):
		res.SetCode(http.StatusNotFound)
	case errors.Is(err, common.ERR_UNAUTHORIZED):
		res.SetCode(http.StatusUnauthorized)
	case errors.Is(err, common.ERR_FORBIDEN):
		res.SetCode(http.StatusForbidden)
	default:
		res.SetCode(http.StatusBadRequest)
	}

	resObj := default_response{
		Message: err.Error(),
	}

	raw, _ := json.Marshal(resObj)

	res.SetContent(raw)

	return res.Done()
}

func (this *Controller) hanleInternalError(err error) (mvc.Result, error) {

	res := this.ActionResult.Prepare()

	if this.ErrorLogger != nil {

		this.ErrorLogger.Error(err.Error())
	}

	resObj := default_response{
		Message: "internal error",
	}

	raw, _ := json.Marshal(resObj)

	return res.SetCode(http.StatusInternalServerError).SetContent(raw).Done()
}
