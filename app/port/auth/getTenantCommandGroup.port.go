package authServicePort

import (
	paginateServicePort "app/port/paginate"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// IGetTenantCommandGroup[Command_Group_Entity_T any] interface {
	// 	Serve(
	// 		tenantUUID uuid.UUID, page uint64, size uint64, cursor *primitive.ObjectID, isPrev bool, ctx context.Context,
	// 	) ([]Command_Group_Entity_T, error)
	// }

	IGetTenantCommandGroups[Entity_T any] interface {
		paginateServicePort.IPaginateService[Entity_T, primitive.ObjectID]
	}
)
