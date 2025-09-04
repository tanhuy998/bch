package mvc

import (
	libCommon "app/internal/lib/common"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type (
	IMvcLauncher interface {
		LaunchMvc(
			party router.Party, option ...mvc.Option,
		) *mvc.Application
	}
)

func Launch[Controller_T IMvcLauncher](party router.Party, options ...mvc.Option) *mvc.Application {

	var controller Controller_T

	libCommon.EvaluateIndirectValueOf(&controller)

	return controller.LaunchMvc(party, options...)
}
