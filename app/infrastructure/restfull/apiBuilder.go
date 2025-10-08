package restfull

import (
	v1 "app/infrastructure/restfull/api/v1"
	"app/infrastructure/restfull/common/log"
	"app/infrastructure/restfull/common/middleware"
	libCommon "app/internal/lib/common"

	"github.com/kataras/iris/v12"
)

type (
	IApplicationAPIBuilder interface {
		BeforeInitializtion(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder
		AfterInitialization(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder
		BeforeLocalDependencies(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder
		AfterLocalDependencies(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder

		BeforeEndpoints(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder
		AfterEndpoints(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder
		Build(endpointOptions ...RestFullAPIInitializationOption) *API
	}
)

type (
	API APIBuilder
)

type (
	APIBuilder struct {
		*iris.Application
		messure          func()
		before_init      []RestFullAPIInitializationOption
		after_init       []RestFullAPIInitializationOption
		before_local_dep []RestFullAPIInitializationOption
		after_local_dep  []RestFullAPIInitializationOption
		before_endpoints []RestFullAPIInitializationOption
		after_endpoints  []RestFullAPIInitializationOption
	}
)

func NewAPI(options ...RestFullAPIInitializationOption) /**iris.Application*/ IApplicationAPIBuilder {

	apiBuilder := new(APIBuilder)

	apiBuilder.messure = libCommon.LMessureTime("Restfull API initialization time", log.Logger())

	app := iris.New()

	app.ConfigureContainer().EnableStructDependents()

	apiBuilder.Application = app

	applyOptions(app, options...)

	app.UseRouter(
		middleware.InternalAccessLog(
			app.ConfigureContainer().Container,
		),
	)

	v1.Initialize(app)

	log.Logger().Println("Registered endpoints")

	for _, route := range app.GetRoutes() {

		log.Logger().Println(route)
	}

	return apiBuilder
}

func (this *APIBuilder) Build(endpointOptions ...RestFullAPIInitializationOption) *API {

	defer this._dispose()

	if this.messure != nil {

		defer this.messure()
	}

	app := this.Application

	applyOptions(app, this.before_init...)

	applyOptions(app, this.before_local_dep...)

	applyOptions(app, configureLocalDependencies)

	applyOptions(app, this.after_local_dep...)

	applyOptions(app, this.before_endpoints...)

	//v1.Initialize(app)

	applyOptions(app, endpointOptions...)

	applyOptions(app, this.after_endpoints...)

	applyOptions(app, this.after_init...)

	for _, route := range app.GetRoutes() {

		log.Logger().Println(route)
	}

	return (*API)(this)
}

func (this *APIBuilder) _dispose() {

	this.before_init = nil
	this.before_endpoints = nil
	this.before_local_dep = nil

	this.after_endpoints = nil
	this.after_init = nil
	this.after_local_dep = nil
}

func (this *APIBuilder) BeforeInitializtion(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder {

	this.before_init = append(this.before_init, options...)

	return this
}

func (this *APIBuilder) AfterInitialization(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder {

	this.after_init = append(this.after_init, options...)

	return this
}

func (this *APIBuilder) BeforeLocalDependencies(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder {

	this.before_local_dep = append(this.before_local_dep, options...)

	return this
}

func (this *APIBuilder) AfterLocalDependencies(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder {

	this.after_local_dep = append(this.after_local_dep, options...)

	return this
}

func (this *APIBuilder) BeforeEndpoints(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder {

	this.before_endpoints = append(this.before_endpoints, options...)

	return this
}

func (this *APIBuilder) AfterEndpoints(options ...RestFullAPIInitializationOption) IApplicationAPIBuilder {

	this.after_endpoints = append(this.after_endpoints, options...)

	return this
}
