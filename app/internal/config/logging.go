package config

import (
	"os"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/middleware/accesslog"
)

func ConfigureLogger(app router.Party) *accesslog.AccessLog {

	ac := accesslog.New(os.Stdout)
	ac.RequestBody = true
	ac.PanicLog = accesslog.LogStack

	app.UseRouter(ac.Handler)

	return ac
}
