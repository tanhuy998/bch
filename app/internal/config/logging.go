package config

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
)

func ConfigureLogger(app *iris.Application) *accesslog.AccessLog {

	ac := accesslog.New(os.Stdout)
	ac.RequestBody = true
	ac.PanicLog = accesslog.LogStack

	app.UseRouter(ac.Handler)

	return ac
}
