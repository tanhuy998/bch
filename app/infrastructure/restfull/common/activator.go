package common

import (
	"github.com/kataras/iris/v12/mvc"
)

type (
	activate_controller_t struct {
		activator mvc.BeforeActivation
	}
)

func (this *activate_controller_t) BeforeActivation(activator mvc.BeforeActivation) {

	this.activator = activator
}
