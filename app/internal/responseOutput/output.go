package responseOutput

/*
this package define interfaces that helps controllers determine which status
the app response to the client.
*/

type (
	INoContentOutput interface {
		IsNotContent() bool
	}

	ICreatedOutput interface {
		IsCreatedStatus() bool
	}

	NoContent struct {
	}

	Result[Input_T, Output_T any] struct {
		Input  *Input_T
		Output *Output_T
	}

	CreatedResponse struct {
	}
)

func NewResult[Input_T, Output_T any](input *Input_T, output *Output_T) *Result[Input_T, Output_T] {

	return &Result[Input_T, Output_T]{
		Input:  input,
		Output: output,
	}
}

func (this *NoContent) IsNotContent() bool {

	return true
}

type ()

func (this CreatedResponse) IsCreatedStatus() bool {

	return true
}
