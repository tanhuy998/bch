package api

import (
	"app/src/infrastructure/http/api/v1/controller"
	"app/src/infrastructure/http/middleware"
	"app/src/infrastructure/http/middleware/middlewareHelper"

	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func initAssignmentApi(app router.Party) *mvc.Application {

	router := app.Party("/assigns")

	container := router.ConfigureContainer().Container

	wrapper := mvc.New(router)

	wrapper.Router.Use(
		middleware.Auth(
			container,
			middlewareHelper.AuthRequireTenantAgent,
		),
	)

	wrapper.Handle(
		new(controller.AssignmentController),
		// applyRoutes(func(activator *mvc.ControllerActivator) {

		// 	activator.Handle(
		// 		"GET", "/{uuid:uuid}", "GetSingleAssignment",
		// 		middleware.BindPresenters[requestPresenter.GetSingleAssignmentRequest, responsePresenter.GetSingleAssignmentResponse](
		// 			container,
		// 			middlewareHelper.UseAuthority,
		// 		),
		// 	)

		// 	activator.Handle(
		// 		"POST", "/", "CreateAssignment",
		// 		middleware.BindPresenters[requestPresenter.CreateAssigmentRequest, responsePresenter.CreateAssignmentResponse](
		// 			container,
		// 			middlewareHelper.UseAuthority,
		// 		),
		// 	)

		// 	activator.Handle(
		// 		"POST", "/{assignmentUUID:uuid}/group", "CreateAssignmentGroup",
		// 		middleware.BindPresenters[requestPresenter.CreateAssignmentGroupRequest, responsePresenter.CreateAssignmentGroupResponse](
		// 			container,
		// 			middlewareHelper.UseAuthority,
		// 		),
		// 	)
		// }),
	)

	return wrapper
}
