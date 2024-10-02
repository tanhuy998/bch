package authServicePort

import (
	"app/model"
	crudResourcePort "app/port/crudResource"
)

type (
	ICreateCommandGroup interface {
		crudResourcePort.ICreateResource[model.CommandGroup]
		Serve(groupName string) error
	}
)
