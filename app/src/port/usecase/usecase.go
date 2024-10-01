package usecasePort

type (
	IUseCase[Request_Struc_T, Response_Struct_T, Result_T any] interface {
		Execute(*Request_Struc_T, *Response_Struct_T) (Result_T, error)
	}
)
