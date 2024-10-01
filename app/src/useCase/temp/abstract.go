package usecase

type (
	IUseCase[Request_T, Response_T, ActionResult_I any] interface {
		Execute(input *Request_T, output *Response_T) (ActionResult_I, error)
	}
)
