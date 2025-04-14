package requestInput

import (
	paginateServicePort "app/port/paginate"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	PaginateInput struct {
		RangePaginateInput
		MongoCursorPaginateInput
	}
)

func (this PaginateInput) GetPaginator() paginateServicePort.IPaginator[primitive.ObjectID] {

	return this
}
