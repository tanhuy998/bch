package api

import (
	authManipulationApi "app/internal/api/auth/manipulation"

	"github.com/kataras/iris/v12"
)

func initAuthApi(app *iris.Application) {

	genericRouter := app.Party("/auth")

	// add authorization for this router
	manipulationRouter := genericRouter.Party("/man")

	//userManipulationRouter := manipulationRouter.Party("/users")
	authManipulationApi.RegisterUserApi(manipulationRouter)
	//registerUserApi(userManipulationRouter).EnableStructDependents()

	commandRouter := manipulationRouter.Party("/command")
	authManipulationApi.RegisterCommandGroupApi(commandRouter)
	// commandGroupRouter := commandRouter.Party("/groups")
	// registerCommandGroupApi(commandGroupRouter)
}

// func initAuthApi(parentRouter router.Party) *mvc.Application {

// 	router := parentRouter.Party("/man")

// 	router.ConfigureContainer(func(api *iris.APIContainer) {

// 		api.Use(middleware.Authentication())
// 	})

// 	registerUserApi(router)

// }

// func registerUserApi(parentRoute router.Party) *mvc.Application {

// 	router := parentRoute.Party("/users")

// 	container := router.ConfigureContainer().Container
// 	controller := new(controller.AuthManipulationController)

// 	wrapper := mvc.New(router)

// 	wrapper.Handle(
// 		controller,
// 		applyRoutes(func(activator *mvc.ControllerActivator) {

// 			activator.Handle(
// 				"POST", "/", "CreateUser",
// 				middleware.BindPresenters[requestPresenter.CreateUserRequestPresenter, responsePresenter.CreateUserPresenter](container),
// 			)

// 		}),
// 	)

// 	return wrapper
// }

// func registerCommandGroupApi(parentRoute router.Party) *mvc.Application {

// }
