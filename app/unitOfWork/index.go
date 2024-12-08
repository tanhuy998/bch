package unitOfWork

import (
	"app/repository"
	genericUseCase "app/unitOfWork/genericUsecase"
	opLog "app/unitOfWork/operationLog"
	"app/valueObject/requestInput"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	OperationLogger  = opLog.OperationLogger
	IOperationLogger = opLog.IOperationLogger

	PaginateUseCase[Entity_T any, Cursor_T comparable, Repository_T repository.IPaginateRepository[bson.E, Entity_T, Cursor_T]] struct {
		genericUseCase.PaginateUseCase[Entity_T, Cursor_T, Repository_T]
	}

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
