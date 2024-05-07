package api

import (
	"app/app/controller"
	"app/app/middleware"
	"app/app/model"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func applyRoutes(f func(*mvc.ControllerActivator)) mvc.OptionFunc {

	return f
}

func Init(app *iris.Application) {

	initAdminApi(app)
	inituserApi(app)
}

func initAdminApi(app *iris.Application) {

	router := app.Party("/admin")

	// router.Use(middleware.Authentication())
	// router.Use(middleware.Authorization([]model.UserGroup{}))

	router.ConfigureContainer(func(api *iris.APIContainer) {

		api.Use(middleware.Authentication())
		api.Use(middleware.Authorization([]model.UserGroup{}))
	})

	wrapper := mvc.New(router)
	wrapper.Handle(
		new(controller.AdminController),
		applyRoutes(func(activator *mvc.ControllerActivator) {

			activator.Handle("GET", "/", "Index")
			activator.Handle("POST", "/login", "Login")
		}),
	)
}

func inituserApi(app *iris.Application) {

	wrapper := mvc.New(app.Party("/"))
	wrapper.Handle(
		new(controller.UserController),
		applyRoutes(func(activator *mvc.ControllerActivator) {

			activator.Handle("GET", "/", "Index")
			activator.Handle("GET", "/campaign/{uuid:string}/candidate/{uuid:string}", "Candidate")
		}),
	)
}
