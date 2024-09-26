package responsePresetService

import (
	actionResultService "app/service/actionResult"
	"net/http"

	"github.com/kataras/iris/v12/mvc"
)

type (
	ResponsePresetService struct {
		ActionResult actionResultService.IActionResult
	}
)

func (this *ResponsePresetService) InternalError() (mvc.Result, error) {

	return this.ActionResult.Prepare().
		SetCode(http.StatusInternalServerError).
		SetContent([]byte(`{"message": "500 internal error"}`)).
		Done()
}

func (this *ResponsePresetService) UnAuthorizedResource() (mvc.Result, error) {

	return this.ActionResult.Prepare().
		SetCode(http.StatusForbidden).
		SetContent([]byte(`{"message": "unauthorized resource"}`)).
		Done()
}

func (this *ResponsePresetService) NotFound() (mvc.Result, error) {

	return this.ActionResult.Prepare().
		SetCode(http.StatusNotFound).
		SetContent([]byte(`{"message": "not found"}`)).
		Done()
}
