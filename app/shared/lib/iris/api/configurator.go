package api

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type (
	Configurator func(parent iris.Party)
)

func (fn Configurator) Configure(parent router.Party) {
	fn(parent)
}
