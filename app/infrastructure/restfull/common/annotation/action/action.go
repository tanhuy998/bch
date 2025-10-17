package action

import (
	"app/infrastructure/restfull/common/annotation/input"
	"app/infrastructure/restfull/common/endpoint"
	"app/infrastructure/restfull/common/endpoint/annotation"
	"app/infrastructure/restfull/common/endpoint/annotationScope"
	"app/infrastructure/restfull/common/result"
	usecasePort "app/port/usecase"
	"fmt"

	"github.com/kataras/iris/v12/mvc"
)

type (
	UseCaseFor[Input_T any, Output_T any] struct {
		annotation.Annotation
		annotationScope.Method
		input          input.Bind[Input_T]
		result.Handler `ioc:"autowired"`
		Usecase        usecasePort.IUseCase[Input_T, Output_T] `ioc:"autowired"`
	}
)

func (a UseCaseFor[Input_T, Output_T]) Singleton() {}

func (a UseCaseFor[Input_T, Output_T]) Once() {}

func (this UseCaseFor[Input_T, Output_T]) Apply(e endpoint.IEndpoint, asset interface{}) {
	// BuildAction() removes temporary empty route
	// BuildAction() is once operation, if invoke multiple time or invoke it when
	// the action was built in the endpoint construction method, it will throws panic.
	// BuildAction() must be invoked before input binding
	// When input binding is applied before BuildAction,
	// all Binded dependencies before BuildAction are flushed.
	e.BuildAction(
		func(input *Input_T) (mvc.Result, error) {
			fmt.Println()
			output, err := this.Usecase.Execute(input)

			switch u := this.Usecase.(type) {
			case result.IDisposer:
				return this.ResultAndDispose(
					output, err, u,
				)
			default:
				return this.ResultOf(
					input, err,
				)
			}
		},
	)

	this.input.Apply(e, asset)
}

func (u UseCaseFor[Input_T, Output_T]) Panic() {}

func (this UseCaseFor[Input_T, Output_T]) Accumulate(asset interface{}) interface{} {

	return this.input.Accumulate(asset)
}

func (this UseCaseFor[Input_T, Output_T]) GetAccumulatorKey() interface{} {

	return this.input.GetAccumulatorKey()
}
