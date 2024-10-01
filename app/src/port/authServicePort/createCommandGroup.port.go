package authServicePort

import (
	"app/src/model"
	crudResourcePort "app/src/port/crudResource"
)

type (
	ICreateCommandGroup interface {
		crudResourcePort.ICreateResource[model.CommandGroup]
		Serve(groupName string) error
	}
)
