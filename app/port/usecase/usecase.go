package usecasePort

import (
	"app/internal/responseOutput"
	"context"
)

type (
	IMiddlewareUseCase interface {
		Execute(ctx context.Context) error
	}

	IUseCase[Request_Struc_T, Response_Struct_T any] interface {
		Execute(*Request_Struc_T) (*Response_Struct_T, error)
	}

	UseCase[Input_T, Output_T any] struct {
	}
)

func (this *UseCase[Input_T, Output_T]) GenerateOutput() *Output_T {

	return new(Output_T)
}

func (this *UseCase[Input_T, Output_T]) ErrorWithContext(
	input *Input_T, err error,
) *responseOutput.ErrorOutput[Input_T] {

	return responseOutput.NewErrorOutput(input, err)
}
