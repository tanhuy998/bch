package authServicePort

import (
	"app/model"
)

type (
	IGetSingleCommandGroup interface {
		Serve(uuid string) (*model.CommandGroup, error)
		SearchByName(groupName string) (*model.CommandGroup, error)
		CheckCommandGroupExistence(groupName string) (bool, error)
	}
)
