package serviceConfig

import "github.com/kataras/iris/v12/hero"

type (
	IServiceConfigurator interface {
		ConfigureServices(container *hero.Container)
	}
)
