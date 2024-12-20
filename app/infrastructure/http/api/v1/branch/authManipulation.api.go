package api

import (
	"app/infrastructure/http/api/v1/branch/auth/authGenApi"
	authManipulationApi "app/infrastructure/http/api/v1/branch/auth/manipulation"
	authReportApi "app/infrastructure/http/api/v1/branch/auth/report"
	authSignaturesApi "app/infrastructure/http/api/v1/branch/auth/signatures"

	"github.com/kataras/iris/v12/core/router"
)

func initAuthApi(app router.Party) {

	genericRouter := app.Party("/auth")

	authGenApi.RegisterGenAPI(genericRouter).EnableStructDependents()
	authSignaturesApi.RegisterSignaturesApi(genericRouter).EnableStructDependents()

	reportRouter := genericRouter.Party("/report")
	authReportApi.RegisterAuthReportApi(reportRouter).EnableStructDependents()

	// add authorization for this router
	manipulationRouter := genericRouter.Party("/man")

	//userManipulationRouter := manipulationRouter.Party("/users")
	authManipulationApi.RegisterUserApi(manipulationRouter).EnableStructDependents()
	//registerUserApi(userManipulationRouter).EnableStructDependents()

	authManipulationApi.RegisterRoleApi(manipulationRouter).EnableStructDependents()

	commandRouter := manipulationRouter.Party("/command")
	authManipulationApi.RegisterCommandGroupApi(commandRouter).EnableStructDependents()
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
