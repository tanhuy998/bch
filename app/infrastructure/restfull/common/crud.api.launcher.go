package common

import (
	"app/infrastructure/restfull/common/endpoint"
	libCommon "app/internal/lib/common"
	"reflect"

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

func instantiateCurator[T endpoint.IAPIEndpointCurator](ptr *T) {

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

func (this *CrudAPILauncher[Read_Curator_T, Write_Controller_T, Create_Curator_T, Delete_Curator_T]) BeforeLaunch(
	parentApp *mvc.Application, options []mvc.Option,
) {

	//parentApp.EnableStructDependents().Handle(&this, options...)

	childApp := mvc.New(parentApp.Router.Party("/")).EnableStructDependents()

	this.init()

	childApp.Handle(
		this.create, options...,
	).Handle(
		this.read, options...,
	).Handle(
		this.update, options...,
	).Handle(
		this.delete, options...,
	)
}

func (this *CrudAPILauncher[Read_Curator_T, Write_Controller_T, Create_Curator_T, Delete_Curator_T]) init() {

	instantiateCurator(&this.read)
	instantiateCurator(&this.update)
	instantiateCurator(&this.create)
	instantiateCurator(&this.delete)
}

func (this *CrudAPILauncher[Read_Curator_T, Create_Curator_T, Update_Curator_T, Delete_Curator_T]) Child() []endpoint.IAPICurator {

	return []endpoint.IAPICurator{
		this.read, this.create, this.update, this.delete,
	}
}
