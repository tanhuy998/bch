package usecasePort

import (
	"app/internal/responseOutput"
	"app/unitOfWork"
	"app/valueObject/requestInput"
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
		unitOfWork.OperationLogger
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

/*
Intercept the responses of a usecase
*/
func (this *UseCase[Input_T, Output_T]) WrapResults(input **Input_T, output **Output_T, err *error) {

	if input == nil {

		return
	}

	in, ok := any(*input).(requestInput.IContextBringAlong)

	if !ok {

		return
	}

	this.wrapOutput(output, in.GetContext())
	this.wrapError(input, err)
}

func (this *UseCase[Input_T, Output_T]) wrapOutput(output **Output_T, ctx context.Context) {

	if ctx == nil {

		return
	}

	if output == nil {

		return
	}

	o, ok := any(*output).(responseOutput.IResponseContext)

	if !ok {

		return
	}

	o.SetContext(ctx)
}

func (this *UseCase[Input_T, Output_T]) wrapError(input **Input_T, err *error) {

	if err == nil {

		return
	}

	*err = this.ErrorWithContext(*input, *err)
}
