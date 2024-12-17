package paginateUseCase

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
)

type (
	IPaginationFilter repositoryAPI.FilterFunc

	ICursorPaginator interface {
		SetCursor(c interface{})
		SetCursorDirection(paginateServicePort.CursorDirection)
		CursorFirst()
	}

	IOffsetPaginator interface {
		SetOffset(offset uint64)
		SetSize(size uint64)
	}

	IPagaintorProjection = paginateServicePort.IPaginateProjector

	IPaginatorFilter interface {
		ApplyFilter(fn repositoryAPI.FilterFunc)
	}

	IPaginatorInitializer interface {
		ICursorPaginator
		IOffsetPaginator
		IPaginatorFilter
		IPagaintorProjection
	}
)

type (
	paginator[Entity_T any] struct {
		repositoryAPI.IPaginationRepository[Entity_T]
		Offset        uint64
		Size          uint64
		Cursor        interface{}
		CursorNilVal  interface{}
		IsPrev        bool
		isCursorFirst bool
	}
)

func NewPaginator[Entity_T any](
	repo repositoryAPI.IPaginateClonableRepository[Entity_T],
) *paginator[Entity_T] {

	if repo == nil {

		panic("could not initialized new paginator, nil repository given")
	}

	return &paginator[Entity_T]{
		IPaginationRepository: repo.Clone(),
	}
}

func (this *paginator[Entity_T]) SetCursor(c interface{}) {

	this.Cursor = c
}

func (this *paginator[Entity_T]) SetOffset(offset uint64) {

	this.Offset = offset
}

func (this *paginator[Entity_T]) SetSize(size uint64) {

	this.Size = size
}

func (this *paginator[Entity_T]) ApplyFilter(fn repositoryAPI.FilterFunc) {

	this.IPaginationRepository.Filter(fn)
}

func (this *paginator[Entity_T]) Select(fields ...string) {

	this.IPaginationRepository.Select(fields...)
}

func (this *paginator[Entity_T]) ExcludeField(fields ...string) {

	this.IPaginationRepository.ExcludeFields(fields...)
}

func (this *paginator[Entity_T]) CursorFirst() {

	this.isCursorFirst = true
}

func (this *paginator[Entity_T]) SetCursorDirection(dir paginateServicePort.CursorDirection) {

	if dir == paginateServicePort.CURSOR_DIRECTION_PREVIOUS {

		this.IsPrev = true
	}
}
