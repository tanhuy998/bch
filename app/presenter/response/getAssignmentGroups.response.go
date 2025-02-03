package responsePresenter

type (
	GetAssignmentGroups[Res_Data_T any] struct {
		Message string       `json:"message"`
		Data    []Res_Data_T `json:"data"`
	}
)
