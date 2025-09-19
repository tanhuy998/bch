package controller

import "app/infrastructure/restfull/common"

type (
	TenantController struct {
		common.CrudAPILauncher[*read, *create, *update, *delete]
	}
)
