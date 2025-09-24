package assignmentServicePort

import (
	paginateServicePort "app/port/paginate"
	paginateUseCase "app/unitOfWork/genericUsecase/paginate"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IGetAssignmentPaginate[Cursor_T comparable] interface {
		// requestInput.IPaginationInput
		// requestInput.IMongoCursorPaginationInput
		paginateServicePort.IPaginator[Cursor_T]
		GetExpiredFilter() bool
	}

	GetAssignmentsFilter struct {
		PageNumber int
		Size       int
		Cursor     primitive.ObjectID
		Expired    bool
	}

	IGetAssignments[Data_T paginateServicePort.ICursorEntity[interface{}]] interface {
		Serve(
			TenantUUID uuid.UUID, filter paginateServicePort.IPaginator[interface{}], ctx context.Context,
		) (paginateUseCase.INavigator[Data_T, interface{}], error)
	}
)
