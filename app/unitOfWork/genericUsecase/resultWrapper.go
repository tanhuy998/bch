package genericUseCase

import (
	"app/internal/responseOutput"
	"app/valueObject/requestInput"
)

type (
	UseCaseResultWrapper[Input_T, Output_T any] struct {
	}
)

/*
Intercept the responses of a ResultsWrapper
*/
func (this *UseCaseResultWrapper[Input_T, Output_T]) WrapResults(input *Input_T, output **Output_T, err *error) {

	this.WrapOutput(input, output)
	this.WrapError(input, err)
}

func (this *UseCaseResultWrapper[Input_T, Output_T]) WrapOutput(input *Input_T, output **Output_T) {

	if output == nil {

		return
	}

	in, ok := any(input).(requestInput.IContextBringAlong)

	if !ok {

		return
	}

	o, ok := any(*output).(responseOutput.IResponseContext)

	if !ok {

		return
	}

	if o.GetContext() != nil {

		return
	}

	o.SetContext(in.GetContext())
}

func (this *UseCaseResultWrapper[Input_T, Output_T]) WrapError(input *Input_T, err *error) {

	if err == nil || *err == nil {

		return
	}

	if _, ok := (*err).(*responseOutput.ErrorOutput[Input_T]); ok {
		// the error is wrapped
		return
	}

	*err = this.ErrorWithContext(input, *err)
}

func (this *UseCaseResultWrapper[Input_T, Output_T]) ErrorWithContext(
	input *Input_T, err error,
) *responseOutput.ErrorOutput[Input_T] {

	return responseOutput.NewErrorOutput(input, err)
}
