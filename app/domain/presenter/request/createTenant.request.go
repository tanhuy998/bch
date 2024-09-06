package requestPresenter

type (
	CreateTenantInputData struct {
		Name            string `json:"name"`
		Description     string `json:"description"`
		TenantAgentUUID string `json:"tenantAgentUUID"`
	}

	CreateTenantRequest struct {
		Data CreateTenantInputData `json:"data"`
	}
)
