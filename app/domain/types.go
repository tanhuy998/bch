package domain

import (
	paginateServicePort "app/port/paginate"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	PaginateCursorType = primitive.ObjectID

	IPaginator = paginateServicePort.IPaginator[PaginateCursorType]
)
