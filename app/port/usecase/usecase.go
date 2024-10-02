package usecasePort

type (
	IUseCase[Request_Struc_T, Response_Struct_T any] interface {
		Execute(*Request_Struc_T) (*Response_Struct_T, error)
	}

	UseCase[Output_T any] struct {
	}
)

func (this *UseCase[Output_T]) GenerateOutput() *Output_T {

	return new(Output_T)
}
