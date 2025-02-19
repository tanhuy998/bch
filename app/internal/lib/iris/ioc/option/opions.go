package iocOption

import (
	"app/internal/lib/iris/ioc/internal"
	"reflect"

	"github.com/kataras/iris/v12/hero"
)

type (
	DependencyOptionFunc func(container *hero.Container, dep *hero.Dependency)
)

func AsAbstracts(abstracts ...reflect.Type) DependencyOptionFunc {

	return func(container *hero.Container, dep *hero.Dependency) {

		concreteObj := dep.OriginalValue

		for _, t := range abstracts {

			internal.Register(container, t, concreteObj)
		}
	}
}

func BindAs[Abstract_T any]() DependencyOptionFunc {

	return func(container *hero.Container, dep *hero.Dependency) {

		internal.Register(container, reflect.TypeFor[Abstract_T](), dep.OriginalValue)
	}
}

func StructDependents(val bool) DependencyOptionFunc {

	return func(container *hero.Container, dep *hero.Dependency) {

		dep.StructDependents = val
	}
}
