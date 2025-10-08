package ioc

import (
	"app/internal/bootstrap"
	"app/main/internal/dependencies/boundedContext"
	"app/main/internal/dependencies/log"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/hero"
)

func init() {

	bootstrap.Boot()
}

func RegisterServices(app router.Party) {

	app.ConfigureContainer(
		func(api *router.APIContainer) {

			api.EnableStructDependents()

			// api.SetDependencyMatcher(
			// 	func(dep *hero.Dependency, reflector reflect.Type) bool {

			// 		matchDefault := hero.DefaultDependencyMatcher(dep, reflector)

			// 		switch {
			// 		case !matchDefault:
			// 			return false
			// 		default:

			// 		}

			// 		return true
			// 	},
			// )

			var container *hero.Container = api.Container

			log.Main().Println("Wiring dependencies...")

			InitializeENV(container)
			/*
				Database independent services
			*/
			RegisterUtilServices(container)
			// RegisterAdapters(container)
			RegisterCaches(container)
			RegisterAuthDependencies(container)

			// InitializeDatabase(app)
			InitializeDatabase(container)
			/*
				Database dependent services
			*/
			boundedContext.RegisterAuthBoundedContext(container)
			boundedContext.RegisterTenantBoundedContext(container)
			boundedContext.RegisterAuthGenBoundedContext(container)
			boundedContext.RegisterAuthSignaturesBoundedContext(container)
			boundedContext.RegisterAssignmentBoundedContext(container)

			log.Main().Println("Wiring dependencies successully.")
		},
	)

}

func RegisterUtils(app router.Party) {

	app.ConfigureContainer(
		func(api *router.APIContainer) {

			log.Main().Println("Wiring utils services")
			defer log.Main().Println("Wiring utils services successfully.")

			api.EnableStructDependents()

			container := api.Container

			InitializeENV(container)
			RegisterUtilServices(container)
			RegisterCaches(container)
			RegisterAuthDependencies(container)
		},
	)
}

func RegisterDatabases(app router.Party) {

	app.ConfigureContainer(
		func(api *router.APIContainer) {

			log.Main().Println("Wiring database and repositories.")
			defer log.Main().Println("Database and repositories wired successfully.")

			api.EnableStructDependents()

			container := api.Container

			InitializeDatabase(container)
		},
	)
}

func RegisterBoundedContext(app router.Party) {

	app.ConfigureContainer(
		func(api *router.APIContainer) {

			log.Main().Println("Wiring domain bounded contexts.")
			defer log.Main().Println("Domain bounded contexts Wired successfully.")

			api.EnableStructDependents()

			container := api.Container

			boundedContext.RegisterAuthBoundedContext(container)
			boundedContext.RegisterTenantBoundedContext(container)
			boundedContext.RegisterAuthGenBoundedContext(container)
			boundedContext.RegisterAuthSignaturesBoundedContext(container)
			boundedContext.RegisterAssignmentBoundedContext(container)
		},
	)
}

func RegisterNamespaceDependentServices(app router.Party) {

}

func RegisterDomainServices(app router.Party) {

}
