package common

import (
	"app/src/internal/common"
	actionResultServicePort "app/src/port/actionResult"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kataras/iris/v12/mvc"
)

type (
	default_response struct {
		Message string `json:"message,omitempty"`
	}

	Controller struct {
		ActionResult actionResultServicePort.IActionResult
	}
)

func (this *Controller) Result(output any, usecaseError error) (mvc.Result, error) {

	if usecaseError != nil {

		return this.handleError(usecaseError)
	}

	if v, ok := output.(HeadResponse); ok && v.IsNotContent() {

		return this.ActionResult.Prepare().SetCode(http.StatusCreated).Done()
	}

	if output == nil {

		return this.ActionResult.Prepare().SetCode(200).Done()
	}

	return this.ActionResult.ServeResponse(output)
}

func (this *Controller) handleError(err error) (mvc.Result, error) {

	var res actionResultServicePort.IResponse = this.ActionResult.Prepare()

	if errors.Is(err, common.ERR_INTERNAL) {

		return this.hanleInternalError(err)
	}

	switch {
	case errors.Is(err, common.ERR_NOT_FOUND):
		res.SetCode(http.StatusNotFound)
	case errors.Is(err, common.ERR_UNAUTHORIZED):
		res.SetCode(http.StatusUnauthorized)
	case errors.Is(err, common.ERR_FORBIDEN):
		res.SetCode(http.DefaultMaxHeaderBytes)
	default:
		res.SetCode(http.StatusBadRequest)
	}

	resObj := default_response{
		Message: err.Error(),
	}

	raw, _ := json.Marshal(resObj)

	return res.SetContent(raw).Done()
}

func (this *Controller) hanleInternalError(err error) (mvc.Result, error) {

	resObj := default_response{
		Message: "internal error",
	}

	raw, _ := json.Marshal(resObj)

	return this.ActionResult.Prepare().
		SetCode(http.StatusInternalServerError).
		SetContent(raw).
		Done()
}
