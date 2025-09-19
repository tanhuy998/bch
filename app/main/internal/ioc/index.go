package ioc

import (
	"app/internal/bootstrap"
	"app/main/internal/dependencies/boundedContext"
	"app/main/internal/dependencies/log"

	"reflect"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/hero"
)

var global_ioc_container *hero.Container

func init() {

	bootstrap.Boot()
}

func WireDependencies(obj interface{}) {

	if global_ioc_container == nil {

		panic("could wire dependencies, global ioc container is absent")
	}

	global_ioc_container.Struct(obj, 0)
}

func RegisterServices(app router.Party) {

	app.ConfigureContainer(
		func(api *router.APIContainer) {
			api.SetDependencyMatcher(
				func(dep *hero.Dependency, reflector reflect.Type) bool {

					matchDefault := hero.DefaultDependencyMatcher(dep, reflector)

					switch {
					case !matchDefault:
						return false
					default:

					}

					return true
				},
			)
			api.EnableStructDependents()

			var container *hero.Container = api.Container

			defer func() {

				global_ioc_container = container
			}()

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

func RegisterNamespaceDependentServices(app router.Party) {

}

func RegisterDomainServices(app router.Party) {

}
