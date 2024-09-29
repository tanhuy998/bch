package authServiceAdapter

import (
	crudResourcePort "app/adapter/crudResorce"
	"app/domain/model"
)

type (
	ICreateCommandGroup interface {
		crudResourcePort.ICreateResource[model.CommandGroup]
		Serve(groupName string) error
	}
)
