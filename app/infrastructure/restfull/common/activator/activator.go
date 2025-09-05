package activator

import (
	"github.com/kataras/iris/v12/mvc"
)

type (
	IActivator interface {
		mvc.BeforeActivation
	}
)

type (
	IControllerActivator interface {
		BeforeActivation(activator mvc.BeforeActivation)
	}
)

type (
	ActivateController struct {
		activator mvc.BeforeActivation
	}
)

func (this *ActivateController) BeforeActivation(activator mvc.BeforeActivation) {

	this.activator = activator
}

func (this *ActivateController) Activator() IActivator {

	return this.activator
}
