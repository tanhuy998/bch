package common

import (
	libError "app/internal/lib/error"
	"app/internal/responseOutput"
	actionResultServicePort "app/port/actionResult"
	loggerPort "app/port/loggerServicePort"
	"encoding/json"
	"net/http"

	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/mvc"
)

type ()

type (
	default_response_body_t struct {
		HttpStatusCode int    `json:"status,omitempty"`
		Message        string `json:"message,omitempty"`
	}

	IController interface {
		BindDependencies(container *hero.Container) IController
	}

	Controller struct {
		ActionResult actionResultServicePort.IActionResult
		ErrorLogger  loggerPort.ErrorLogger
		ErrorHandler
	}
)

func (this *Controller) ResultOf(output any, usecaseError error) (mvc.Result, error) {

	defer this.logResult(output)

	if usecaseError != nil {

		return this.dispatchError(usecaseError)
	}

	if output == nil {

		return this.ActionResult.Prepare().SetCode(http.StatusNoContent).Done()
	}

	if v, ok := output.(responseOutput.IHTTPStatusResponse); ok {

		raw, err := json.Marshal(output)

		if err != nil {

			return this.dispatchError(libError.NewInternal(err))
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

			return this.dispatchError(libError.NewInternal(err))
		}

		return this.ActionResult.Prepare().SetCode(http.StatusCreated).SetContent(raw).Done()
	}

	if v, ok := output.(responseOutput.IAcceptedOuput); ok && v.IsAccepptedStatus() {

		raw, err := json.Marshal(output)

		if err != nil {

			return this.dispatchError(libError.NewInternal(err))
		}

		return this.ActionResult.Prepare().SetCode(http.StatusAccepted).SetContent(raw).Done()
	}

	return this.ActionResult.ServeResponse(output)
}

func (this *Controller) logResult(res any) {

	if res == nil {

		return
	}

	v, ok := res.(responseOutput.ILoggableResponse)

	if !ok {

		return
	}

	ctx := v.GetContext()

	if ctx == nil {

		return
	}

	this.AccessLogger.WriteMessage(ctx, v.GetMessage())
}

func (this *Controller) dispatchError(err error) (mvc.Result, error) {

	res := this.ErrorHandler.HandleError(err, nil)

	return res, nil
}
