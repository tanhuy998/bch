package genericUseCase

import (
	opLog "app/unitOfWork/operationLog"
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
		UseCaseResultWrapper[Input_T, Output_T]
		opLog.OperationLogger
	}
)

func (this *UseCase[Input_T, Output_T]) GenerateOutput() *Output_T {

	return new(Output_T)
}

func (this *UseCase[Input_T, Output_T]) GenerateOutputFrom(input *Input_T) *Output_T {

	ret := new(Output_T)

	this.WrapOutput(input, &ret)

	return ret
}
