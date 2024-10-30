package requestPresenter

import (
	"app/valueObject/requestInput"
)

type (
	GetTenantUsers struct {
		requestInput.ContextInput
		requestInput.AuthorityInput
		requestInput.TenantMappingInput
		requestInput.RangePaginateInput
		requestInput.MongoCursorPaginateInput
	}
)
