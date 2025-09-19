package constraint

import (
	"app/infrastructure/restfull/common/middleware/hook"
)

type (
	TenantAgent struct{}
)

func (TenantAgent) Retrieve() []hook.AuthorityConstraint {

	return []hook.AuthorityConstraint{
		hook.AuthRequireTenantAgent,
	}
}
