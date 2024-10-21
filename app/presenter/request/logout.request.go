package requestPresenter

import "app/valueObject/requestInput"

type (
	Logout struct {
		requestInput.ContextInput
		requestInput.TenantMappingInput
	}
)
