package common

import (
	"app/infrastructure/restfull/common/endpoint"
	libCommon "app/internal/lib/common"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type (
	IReadEndpointCurator[Initiator_T any] interface {
		endpoint.IAPIEndpointCurator
		IActivator
		GET(path string) Initiator_T
		HEAD(path string) Initiator_T
	}
	ICreateEndpointCurator[Initiator_T any] interface {
		endpoint.IAPIEndpointCurator
		IActivator
		POST(path string) Initiator_T
	}
	IUpdateEndpointCurator[Initiator_T any] interface {
		endpoint.IAPIEndpointCurator
		IActivator
		PUT(path string) Initiator_T
		PATCH(path string) Initiator_T
	}
	IDeleteEndpointCurator[Initiator_T any] interface {
		endpoint.IAPIEndpointCurator
		IActivator
		DELETE(path string) Initiator_T
	}
)

type (
	IActivator interface {
		comparable
		BeforeActivation(activator mvc.BeforeActivation)
	}
)

type (
	// IMvcAppLauncher interface {
	// }

	IAPILauncher interface {
		//IMvcAppLauncher
		_launch(app *mvc.Application, options []mvc.Option) *mvc.Application
	}
)

type (
	CrudAPILauncher[
		Read_Curator_T IReadEndpointCurator[endpoint.IEndpointBuilder],
		Create_Curator_T ICreateEndpointCurator[endpoint.IEndpointBuilder],
		Update_Curator_T IUpdateEndpointCurator[endpoint.IEndpointBuilder],
		Delete_Curator_T IDeleteEndpointCurator[endpoint.IEndpointBuilder],
	] struct {
		endpoint.APIEndpointCurator
		read   Read_Curator_T
		update Update_Curator_T
		create Create_Curator_T
		delete Delete_Curator_T
	}
)

type (
	controller_launcher[Launcher_T IAPILauncher] struct {
		app *mvc.Application
		c   Launcher_T
	}
)

func (this *controller_launcher[Controller_t]) Launch(
	options ...mvc.Option,
) *mvc.Application {

	return this.c._launch(this.app, options)
}

func NewControllerLauncher[Controller_T IAPILauncher](
	party router.Party,
) interface {
	Launch(options ...mvc.Option) *mvc.Application
} {

	if party == nil {
		panic("bad party value, nil given")
	}

	c := &controller_launcher[Controller_T]{
		app: mvc.New(party).EnableStructDependents(),
	}

	return c
}

func InstantiateCRUDControllerActivator[T IActivator](ptr *T) {

	if ptr == nil {

		panic("nil arg given")
	}

	t := reflect.TypeFor[T]()

	switch {
	case t.Kind() == reflect.Pointer:
		fallthrough
	case t.Elem().Kind() == reflect.Struct:
		// obj param is double pointer of the actual T
		// ICRUDControllerActivator expects T implemented it
		// as pointer receiver
		libCommon.EvaluateIndirectValueOfPointer(ptr)
	default:
		panic("CRUDController just accept struct that implements ICRUDControllerActivator's methods by pointer recevier")
	}

}

func (copy CrudAPILauncher[Read_Curator_T, Write_Controller_T, Create_Curator_T, Delete_Curator_T]) _launch(
	app *mvc.Application, options []mvc.Option,
) *mvc.Application {

	fmt.Println("------------- _launch()", mvc.IgnoreEmbeddedControllers)

	defer copy.registerEndpoints()

	app.Handle(&copy, options...)

	childApp := mvc.New(app.Router.Party("/"))

	copy.init()

	childApp.Handle(
		copy.create, options...,
	).Handle(
		copy.read, options...,
	).Handle(
		copy.update, options...,
	).Handle(
		copy.delete, options...,
	)

	return app
}

func (this *CrudAPILauncher[Read_Curator_T, Write_Controller_T, Create_Curator_T, Delete_Curator_T]) init() {

	InstantiateCRUDControllerActivator(&this.read)
	InstantiateCRUDControllerActivator(&this.update)
	InstantiateCRUDControllerActivator(&this.create)
	InstantiateCRUDControllerActivator(&this.delete)
}

func (this *CrudAPILauncher[Read_Curator_T, Write_Controller_T, Create_Curator_T, Delete_Curator_T]) registerEndpoints() {

	endpoint.RegisterEndpointsOf(this)
	endpoint.RegisterEndpointsOf(this.read)
	endpoint.RegisterEndpointsOf(this.create)
	endpoint.RegisterEndpointsOf(this.update)
	endpoint.RegisterEndpointsOf(this.delete)
}
