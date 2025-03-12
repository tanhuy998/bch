package domain

import (
	paginateServicePort "app/port/paginate"
	usecasePort "app/port/usecase"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	PaginateCursorType = primitive.ObjectID

	PaginatableUseCase[
		Input_T paginateServicePort.IPaginator[PaginateCursorType], Output_T any,
	] struct {
		usecasePort.PaginatableUseCase[PaginateCursorType, Input_T, Output_T]
	}

	Paginator = paginateServicePort.IPaginator[PaginateCursorType]

	IPaginator = paginateServicePort.IPaginator[PaginateCursorType]
)
