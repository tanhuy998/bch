package namespace

import (
	irisIoc "app/internal/lib/iris/ioc"

	"github.com/kataras/iris/v12/hero"
)

type (
	DependenyBindingFunc func(container *hero.Container)
)

func BindDependency[Abstract_T, Concrete_T any](
	obj *Concrete_T,
) DependenyBindingFunc {

	return func(container *hero.Container) {

		irisIoc.BindDependency[Abstract_T, Concrete_T](container, obj)
	}
}
