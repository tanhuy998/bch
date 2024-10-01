package authServicePort

import (
	"app/src/model"
)

type (
	IGetSingleCommandGroup interface {
		Serve(uuid string) (*model.CommandGroup, error)
		SearchByName(groupName string) (*model.CommandGroup, error)
		CheckCommandGroupExistence(groupName string) (bool, error)
	}
)
