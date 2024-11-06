package responseOutput

type (
	ResponseData[Data_T any] struct {
		Data Data_T `json:"data"`
	}

	ResponseDataList[Data_T any] struct {
		Data []Data_T `json:"data"`
	}
)
