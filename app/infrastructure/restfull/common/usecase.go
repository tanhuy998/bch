package common

import usecasePort "app/port/usecase"

type (
	IDisposer interface {
		Dispose(obj ...interface{})
	}

	IDisposableUsecase[Input_T any, Output_T any] interface {
		usecasePort.IUseCase[Input_T, Output_T]
		IDisposer
	}
)

type (
	DisposableUseCase[
		Usecase_T IDisposableUsecase[Input_T, Output_T],
		Input_T, Output_T any,
	] struct {
		U Usecase_T
	}
)

func (this *DisposableUseCase[Usecase_T, Input_T, Output_T]) Execute(
	input *Input_T,
) (*Output_T, error) {

	ret, err := this.U.Execute(input)

	return ret, err
}
