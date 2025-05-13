package responsePresenter

type (
	GetTenantAllGroups[Data_T any] struct {
		Message string   `json:"message"`
		Data    []Data_T `json:"data"`
	}
)
