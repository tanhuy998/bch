package responsePresetPort

import "github.com/kataras/iris/v12/mvc"

type (
	IResponsePreset interface {
		UnAuthorizedResource() (mvc.Result, error)
		NotFound() (mvc.Result, error)
	}
)
