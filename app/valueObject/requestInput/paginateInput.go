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

func (this PaginateInput) ToGeneralPaginator() paginateServicePort.IPaginator[interface{}] {

	ret, _ := any(this).(paginateServicePort.IPaginator[interface{}])

	return ret
}

func (this PaginateInput) GetGeneralPaginator() paginateServicePort.IPaginator[interface{}] {

	return this.ToGeneralPaginator()
}
