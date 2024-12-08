package requestPresenter

import "app/valueObject/requestInput"

type (
	GetTenantCommandGroups struct {
		requestInput.ContextInput
		requestInput.AuthorityInput
		requestInput.TenantMappingInput
		requestInput.RangePaginateInput
		requestInput.MongoCursorPaginateInput
	}
)
