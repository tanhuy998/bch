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
	IReadController[Initiator_T any] interface {
		endpoint.IEndpointBuilder
		IActivator
		GET(path string) Initiator_T
		HEAD(path string) Initiator_T
	}
	ICreateController[Initiator_T any] interface {
		endpoint.IEndpointBuilder
		IActivator
		POST(path string) Initiator_T
	}
	IUpdateController[Initiator_T any] interface {
		endpoint.IEndpointBuilder
		IActivator
		PUT(path string) Initiator_T
		PATCH(path string) Initiator_T
	}
	IDeleteController[Initiator_T any] interface {
		endpoint.IEndpointBuilder
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
	IMvcAppLauncher interface {
		_launch(app *mvc.Application, options []mvc.Option) *mvc.Application
	}

	ILauncher interface {
		IMvcAppLauncher
	}
)

type (
	CRUDController[
		Read_Controller_T IReadController[endpoint.IEndpointInitiator],
		Create_Contrller_T ICreateController[endpoint.IEndpointInitiator],
		Update_Controller_T IUpdateController[endpoint.IEndpointInitiator],
		Delete_Controller_T IDeleteController[endpoint.IEndpointInitiator],
	] struct {
		Controller
		endpoint.EndpointBuilder
		read   Read_Controller_T
		update Update_Controller_T
		create Create_Contrller_T
		delete Delete_Controller_T
	}
)

type (
	controller_launcher[Controller_t ILauncher] struct {
		app *mvc.Application
		c   Controller_t
	}
)

func (this *controller_launcher[Controller_t]) Launch(
	options ...mvc.Option,
) *mvc.Application {

	return this.c._launch(this.app, options)
}

func NewControllerLauncher[Controller_T ILauncher](
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

func (copy CRUDController[Read_Controller_T, Write_Controller_T, Create_Contrller_T, Delete_Controller_T]) _launch(
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

// func (this *CRUDController[Read_Controller_T, Write_Controller_T, Create_Contrller_T, Delete_Controller_T]) BeforeActivation(
// 	activator mvc.BeforeActivation,
// ) {

// 	this.init()
// }

func (this *CRUDController[Read_Controller_T, Write_Controller_T, Create_Contrller_T, Delete_Controller_T]) init() {

	InstantiateCRUDControllerActivator(&this.read)
	InstantiateCRUDControllerActivator(&this.update)
	InstantiateCRUDControllerActivator(&this.create)
	InstantiateCRUDControllerActivator(&this.delete)
}

func (this *CRUDController[Read_Controller_T, Write_Controller_T, Create_Contrller_T, Delete_Controller_T]) registerEndpoints() {

	endpoint.RegisterEndpointsOf(this)
	endpoint.RegisterEndpointsOf(this.read)
	endpoint.RegisterEndpointsOf(this.create)
	endpoint.RegisterEndpointsOf(this.update)
	endpoint.RegisterEndpointsOf(this.delete)
}
