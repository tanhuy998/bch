package paginateUseCaseOption

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
	paginateUseCase "app/unitOfWork/genericUsecase/paginate"
)

func ByCursor[Cursor_T comparable](c Cursor_T) paginateUseCase.PaginationOption[Cursor_T] {

	return func(paginator paginateUseCase.IPaginatorInitializer[Cursor_T]) {

		paginator.SetCursor(c)
	}
}

func CursorDirection[Cursor_T comparable](dir paginateServicePort.CursorDirection) paginateUseCase.PaginationOption[Cursor_T] {

	return func(paginator paginateUseCase.IPaginatorInitializer[Cursor_T]) {

		paginator.SetCursorDirection(dir)
	}
}

func Offset[Cursor_T comparable](offset uint64) paginateUseCase.PaginationOption[Cursor_T] {

	return func(paginator paginateUseCase.IPaginatorInitializer[Cursor_T]) {

		paginator.SetOffset(offset)
	}
}

func Size[Cursor_T comparable](size uint64) paginateUseCase.PaginationOption[Cursor_T] {

	return func(paginator paginateUseCase.IPaginatorInitializer[Cursor_T]) {

		paginator.SetSize(size)
	}
}

func SelectFields[Cursor_T comparable](fields ...string) paginateUseCase.PaginationOption[Cursor_T] {

	return func(paginator paginateUseCase.IPaginatorInitializer[Cursor_T]) {

		paginator.Select(fields...)
	}
}

func ExcludeFields[Cursor_T comparable](fields ...string) paginateUseCase.PaginationOption[Cursor_T] {

	return func(paginator paginateUseCase.IPaginatorInitializer[Cursor_T]) {

		paginator.ExcludeField(fields...)
	}
}

func Filter[Cursor_T comparable](fn repositoryAPI.FilterFunc) paginateUseCase.PaginationOption[Cursor_T] {

	return func(paginator paginateUseCase.IPaginatorInitializer[Cursor_T]) {

		paginator.ApplyFilter(fn)
	}
}

func ByOffsetWhenNoCursor[Cursor_T comparable](offset uint64, size uint64) paginateUseCase.PaginationOption[Cursor_T] {

	return func(paginator paginateUseCase.IPaginatorInitializer[Cursor_T]) {

		paginator.SetOffset(offset)
		paginator.SetSize(size)
		paginator.CursorFirst()
	}
}

func CursorNilValue[Cursor_T comparable](val Cursor_T) paginateUseCase.PaginationOption[Cursor_T] {

	return func(paginator paginateUseCase.IPaginatorInitializer[Cursor_T]) {

		paginator.SetCursorNilValue(val)
	}
}
