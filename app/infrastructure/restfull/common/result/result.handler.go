package result

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

type (
	IDisposer interface {
		Dispose(obj ...interface{})
	}
)

type (
	default_response_body_t struct {
		HttpStatusCode int    `json:"status,omitempty"`
		Message        string `json:"message,omitempty"`
	}

	IController interface {
		BindDependencies(container *hero.Container) IController
	}

	Handler struct {
		ActionResult actionResultServicePort.IActionResult
		ErrorLogger  loggerPort.ErrorLogger
		ErrorHandler
	}
)

func (this *Handler) ResultAndDispose(
	output any, usecaseError error, disposers ...IDisposer,
) (mvc.Result, error) {

	defer this.disposeOutput(disposers, output, usecaseError)

	res, err := this.ResultOf(output, usecaseError)

	return res, err
}

func (this *Handler) disposeOutput(disposser []IDisposer, objectsToBeDisposed ...interface{}) {

	for _, obj := range disposser {

		obj.Dispose(objectsToBeDisposed...)
	}
}

func (this *Handler) ResultOf(output any, usecaseError error) (mvc.Result, error) {

	defer this.logResult(output)

	if usecaseError != nil {

		return this.dispatchError(usecaseError)
	}

	if output == nil {

		return this.ActionResult.Prepare().SetCode(http.StatusNoContent).Done()
	}

	// if v, ok := output.(responseOutput.IHTTPStatusResponse); ok {

	// 	raw, err := json.Marshal(output)

	// 	if err != nil {

	// 		return this.dispatchError(libError.NewInternal(err))
	// 	}

	// 	return this.ActionResult.Prepare().
	// 		SetCode(
	// 			v.GetHTTPStatus(),
	// 		).
	// 		SetContent(raw).
	// 		Done()
	// }

	// if v, ok := output.(responseOutput.INoContentOutput); ok && v.IsNotContent() {

	// 	return this.ActionResult.Prepare().SetCode(http.StatusNoContent).Done()
	// }

	// if v, ok := output.(responseOutput.ICreatedOutput); ok && v.IsCreatedStatus() {

	// 	raw, err := json.Marshal(output)

	// 	if err != nil {

	// 		return this.dispatchError(libError.NewInternal(err))
	// 	}

	// 	return this.ActionResult.Prepare().SetCode(http.StatusCreated).SetContent(raw).Done()
	// }

	// if v, ok := output.(responseOutput.IAcceptedOuput); ok && v.IsAccepptedStatus() {

	// 	raw, err := json.Marshal(output)

	// 	if err != nil {

	// 		return this.dispatchError(libError.NewInternal(err))
	// 	}

	// 	return this.ActionResult.Prepare().SetCode(http.StatusAccepted).SetContent(raw).Done()
	// }

	// return this.ActionResult.ServeResponse(output)

	switch out := output.(type) {
	case responseOutput.IHTTPStatusResponse:
		raw, err := json.Marshal(output)

		if err != nil {

			return this.dispatchError(libError.NewInternal(err))
		}

		return this.ActionResult.Prepare().
			SetCode(
				out.GetHTTPStatus(),
			).
			SetContent(raw).
			Done()
	case responseOutput.INoContentOutput:
		return this.ActionResult.Prepare().SetCode(http.StatusNoContent).Done()
	case responseOutput.ICreatedOutput:
		raw, err := json.Marshal(output)

		if err != nil {

			return this.dispatchError(libError.NewInternal(err))
		}

		return this.ActionResult.Prepare().SetCode(http.StatusCreated).SetContent(raw).Done()
	case responseOutput.IAcceptedOuput:
		raw, err := json.Marshal(output)

		if err != nil {

			return this.dispatchError(libError.NewInternal(err))
		}

		return this.ActionResult.Prepare().SetCode(http.StatusAccepted).SetContent(raw).Done()
	default:
		return this.ActionResult.ServeResponse(output)
	}
}

func (this *Handler) logResult(res any) {

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

func (this *Handler) dispatchError(err error) (mvc.Result, error) {

	res := this.ErrorHandler.HandleError(err, nil)

	return res, nil
}
