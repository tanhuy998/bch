package unitOfWork

import (
	genericUseCase "app/unitOfWork/genericUsecase"
	opLog "app/unitOfWork/operationLog"
	"app/valueObject/requestInput"
)

type (
	OperationLogger  = opLog.OperationLogger
	IOperationLogger = opLog.IOperationLogger

	UseCaseResultWrapper[Input_T, Output_T any] struct {
		genericUseCase.UseCaseResultWrapper[Input_T, Output_T]
	}
	GenericUseCase[Input_T, Output_T any] struct {
		genericUseCase.UseCase[Input_T, Output_T]
	}
	MongoUserSessionCacheUseCase[Output_T any] struct {
		genericUseCase.MongoUserSessionCacheUseCase[Output_T]
	}
	MongodAuthDomainUseCase[Input_T requestInput.ITenantDomainInput] struct {
		genericUseCase.MongodAuthDomainUseCase[Input_T]
	}

	TenantDomainUseCase[Input_T requestInput.ITenantDomainInput, Output_T any] struct {
		genericUseCase.TenantDomainUseCase[Input_T, Output_T]
	}
)
