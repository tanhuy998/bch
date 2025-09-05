package assignment

import (
	"app/infrastructure/restfull/api/v1/branch/assignment/controller"
	"app/infrastructure/restfull/common"

	"github.com/kataras/iris/v12"
)

func Api(app iris.Party) {

	router := app.Party("/assigns")

	// wrapper := mvc.New(router)

	// wrapper.Router.Use(
	// 	middleware.Auth(container),
	// )

	//wrapper.Router.ConfigureContainer()

	// container := router.ConfigureContainer().Container

	// controller := new(controller.AssignmentController).BindDependencies(container)

	// wrapper.Handle(
	// 	new(controller.AssignmentController),
	// )

	launcher := common.NewControllerLauncher[controller.AssignmentController](router)

	launcher.Launch()
}
