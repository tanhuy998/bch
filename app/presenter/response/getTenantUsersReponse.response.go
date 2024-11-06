package responsePresenter

import (
	"app/internal/responseOutput"
)

type (
	GetTenantUsers[Data_T any] struct {
		//Data []model.User `json:"data"`
		responseOutput.ResponseDataList[Data_T]
	}
)
