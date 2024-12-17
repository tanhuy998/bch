package libConfig

import "github.com/kataras/iris/v12/hero"

type (
	DependencyOptionFunc = func(dep *hero.Dependency)
)

func StructDependents(val bool) DependencyOptionFunc {

	return func(dep *hero.Dependency) {

		dep.StructDependents = val
	}
}
