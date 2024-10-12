package responseOutput

type (
	ErrorOutput[Input_T any] struct {
		error
		Input *Input_T
	}

	ErrorResult[Input_T, Output_T any] struct {
		Result[Input_T, Output_T]
		error
	}
)

func NewErrorOutput[Input_T any](input *Input_T, err error) *ErrorOutput[Input_T] {

	return &ErrorOutput[Input_T]{
		err,
		input,
	}
}

func NewErrorResult[Input_T, Output_T any](input *Input_T, output *Output_T, err error) *ErrorResult[Input_T, Output_T] {

	return &ErrorResult[Input_T, Output_T]{
		Result: Result[Input_T, Output_T]{
			Input:  input,
			Output: output,
		},
		error: err,
	}
}

func (this *ErrorOutput[Input_T]) Unwrap() error {

	return this.error
}
