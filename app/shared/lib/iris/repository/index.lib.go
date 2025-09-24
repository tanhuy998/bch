package repositoryConfig

import (
	"app/infrastructure/http/common/config/repository/binding"
	"fmt"

	"github.com/kataras/iris/v12/hero"
)

func BindRepositories(
	container *hero.Container,
	fns ...binding.RepoRegisterFunc,
) {

	var debugCounter int

	defer func() {

		if r := recover(); r != nil {

			panic(
				fmt.Sprintf(`Failed to bind mongo repository at dependency index %d: %S`, debugCounter, r),
			)
		}
	}()

	for i, fn := range fns {

		debugCounter = i

		fn(container)
	}
}
