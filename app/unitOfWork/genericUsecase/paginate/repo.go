package paginateUseCase

import (
	libCommon "app/internal/lib/common"
	repositoryAPI "app/repository/api"
)

type (
	ICursorPaginateException interface {
		IOffsetPaginate
		Or() IOffsetPaginate
	}

	IOffsetPaginate interface {
		Offset(page uint64, size uint64) IPaginate
	}

	IPaginate interface {
		Paginate()
	}
)

type (
	repo[Entity_T any] struct {
		Repo    repositoryAPI.IPaginationRepository[Entity_T]
		page    uint64
		size    uint64
		cursor  interface{}
		is_prev bool
	}
)

func (this *repo[Entity_T]) Cursor(isPrev bool, cursorVal interface{}) ICursorPaginateException {

	this.is_prev = isPrev
	this.cursor = cursorVal

	return this
}

func (this *repo[Entity_T]) Offset(page uint64, size uint64) IPaginate {

	this.page = page
	this.size = libCommon.Ternary(size == 0, 1, size)

	return this
}

func (this *repo[Entity_T]) Or() IOffsetPaginate {

	return this
}

func (this *repo[Entity_T]) Paginate() {

}
