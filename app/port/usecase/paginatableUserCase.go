package usecasePort

import paginateServicePort "app/port/paginate"

type (
	PaginatableUseCase[
		Paginate_Cursor_T comparable, Input_T paginateServicePort.IPaginator[Paginate_Cursor_T], Output_T any,
	] struct {
		UseCase[Input_T, Output_T]
	}
)
