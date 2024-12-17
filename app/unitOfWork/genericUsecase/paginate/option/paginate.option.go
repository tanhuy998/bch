package paginateUseCaseOption

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
	paginateUseCase "app/unitOfWork/genericUsecase/paginate"
)

func ByCursor(c interface{}) paginateUseCase.PaginationOption {

	return func(paginator paginateUseCase.IPaginatorInitializer) {

		paginator.SetCursor(c)
	}
}

func CursorDirection(dir paginateServicePort.CursorDirection) paginateUseCase.PaginationOption {

	return func(paginator paginateUseCase.IPaginatorInitializer) {

		paginator.SetCursorDirection(dir)
	}
}

func Offset(offset uint64) paginateUseCase.PaginationOption {

	return func(paginator paginateUseCase.IPaginatorInitializer) {

		paginator.SetOffset(offset)
	}
}

func Size(size uint64) paginateUseCase.PaginationOption {

	return func(paginator paginateUseCase.IPaginatorInitializer) {

		paginator.SetSize(size)
	}
}

func SelectFields(fields ...string) paginateUseCase.PaginationOption {

	return func(paginator paginateUseCase.IPaginatorInitializer) {

		paginator.Select(fields...)
	}
}

func ExcludeFields(fields ...string) paginateUseCase.PaginationOption {

	return func(paginator paginateUseCase.IPaginatorInitializer) {

		paginator.ExcludeField(fields...)
	}
}

func Filter(fn repositoryAPI.FilterFunc) paginateUseCase.PaginationOption {

	return func(paginator paginateUseCase.IPaginatorInitializer) {

		paginator.ApplyFilter(fn)
	}
}

func ByOffsetWhenNoCursor(offset uint64, size uint64) paginateUseCase.PaginationOption {

	return func(paginator paginateUseCase.IPaginatorInitializer) {

		paginator.SetOffset(offset)
		paginator.SetSize(size)
		paginator.CursorFirst()
	}
}
