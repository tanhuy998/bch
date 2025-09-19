package common

import (
	"app/infrastructure/restfull/common/endpoint"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type (
	IBeforeLaucnhEvent interface {
		BeforeLaunch(parentApp *mvc.Application, options []mvc.Option)
	}
)

type (
	APILauncher struct {
		router router.Party
	}
)

func NewAPILauncher(router router.Party) *APILauncher {

	return &APILauncher{router: router}
}

func (this *APILauncher) LaunchAPIOf(curator endpoint.IAPICurator, options ...mvc.Option) *mvc.Application {

	app := mvc.New(this.router).EnableStructDependents()

	app.Handle(curator, options...)

	switch o := curator.(type) {
	case IBeforeLaucnhEvent:
		o.BeforeLaunch(app, options)
	}

	endpoint.LaunchApiOf(curator)

	return app
}
