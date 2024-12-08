package responsePresenter

import "app/internal/responseOutput"

type (
	GetTenantCommandGroups[Output_T any] struct {
		Message string `json:"message"`
		responseOutput.ResponseDataList[Output_T]
	}
)
