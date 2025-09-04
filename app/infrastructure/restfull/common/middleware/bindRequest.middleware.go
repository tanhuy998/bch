package middleware

import (
	"app/infrastructure/restfull/common/middleware/hook"
)

type ()

func BindRequest[Request_T any](
	//container *hero.Container,
	hooks ...hook.PresenterBindingHook[Request_T, EmptyPresenter],
) ContainerDependentMiddleware {

	//return BindPresenters(container, initializer...)
	return BindPresenters(hooks...)
}
