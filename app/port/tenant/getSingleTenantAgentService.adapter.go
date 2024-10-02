package tenantServicePort

import "app/model"

type (
	IGetSingleTenantAgent interface {
		Serve(uuid string) (*model.TenantAgent, error)
		SearchByUsername(username string) (*model.TenantAgent, error)
		CheckUsernameExistence(username string) (bool, error)
	}
)
