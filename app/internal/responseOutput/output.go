package responseOutput

/*
this package define interfaces that helps controllers determine which status
the app response to the client.
*/

type (
	INoContentOutput interface {
		IsNotContent() bool
	}

	NoContent struct {
	}

	ErrorOutput[Input_T any] struct {
		error
		Input *Input_T
	}

	Result[Input_T, Output_T any] struct {
		Input  *Input_T
		Output *Output_T
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

func NewResult[Input_T, Output_T any](input *Input_T, output *Output_T) *Result[Input_T, Output_T] {

	return &Result[Input_T, Output_T]{
		Input:  input,
		Output: output,
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

func (this *NoContent) IsNotContent() bool {

	return true
}

type (
	ICreatedOutput interface {
		IsCreatedStatus() bool
	}

	CreatedResponse struct {
	}
)

func (this CreatedResponse) IsCreatedStatus() bool {

	return true
}
