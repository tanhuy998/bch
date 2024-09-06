package tenantAgentServiceAdapter

import "app/domain/model"

type (
	IGetSingleTenantAgentServiceAdapter interface {
		Serve(uuid string) (*model.TenantAgent, error)
		SearchByUsername(username string) (*model.TenantAgent, error)
		CheckUsernameExistence(username string) (bool, error)
	}
)
