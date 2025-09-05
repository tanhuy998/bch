package common

import (
	"app/infrastructure/restfull/common/endpoint"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IAPILauncher interface {
		_launch(app *mvc.Application, options []mvc.Option) *mvc.Application
	}
)

type (
	controller_launcher[Launcher_T IAPILauncher] struct {
		app *mvc.Application
		l   Launcher_T
	}
)

func (this *controller_launcher[Controller_t]) Launch(
	options ...mvc.Option,
) *mvc.Application {

	return this.l._launch(this.app, options)
}

type (
	APILauncher[Curator_t endpoint.IAPIEndpointCurator] struct {
		curator Curator_t
	}
)

func (copy APILauncher[Curator_t]) _launch(app *mvc.Application, options []mvc.Option) *mvc.Application {

	instantiateCurator(&copy.curator)

	return app.Handle(&copy, options...)
}
