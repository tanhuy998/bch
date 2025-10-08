package middleware

import (
	"app/infrastructure/restfull/common/middleware/hook/binding"
)

type ()

func BindRequest[Request_T any](
	hooks ...binding.Hook,
) ContainerDependentMiddleware {

	return binding.BindPresenters[Request_T, EmptyPresenter](hooks...)
}
