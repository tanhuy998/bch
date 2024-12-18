package paginateUseCase

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
)

type (
	IPaginationFilter repositoryAPI.FilterFunc

	ICursorPaginator[Cursor_T comparable] interface {
		SetCursor(c Cursor_T)
		SetCursorNilValue(v Cursor_T)
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

	IPaginatorInitializer[Cusor_T comparable] interface {
		ICursorPaginator[Cusor_T]
		IOffsetPaginator
		IPaginatorFilter
		IPagaintorProjection
	}
)

type (
	paginator[Entity_T any, Cursor_T comparable] struct {
		repositoryAPI.IPaginationRepository[Entity_T]
		Cursor        *Cursor_T // interface{}
		CursorNilVal  *Cursor_T //interface{}
		Offset        uint64
		Size          uint64
		IsPrev        bool
		isCursorFirst bool
	}
)

func NewPaginator[Entity_T any, Cursor_T comparable](
	repo repositoryAPI.IPaginateClonableRepository[Entity_T],
) *paginator[Entity_T, Cursor_T] {

	if repo == nil {

		panic("could not initialized new paginator, nil repository given")
	}

	return &paginator[Entity_T, Cursor_T]{
		IPaginationRepository: repo.Clone(),
	}
}

func (this *paginator[Entity_T, Cursor_T]) SetCursor(c Cursor_T) {

	this.Cursor = &c
}

func (this *paginator[Entity_T, Cursor_T]) SetCursorNilValue(v Cursor_T) {

	this.CursorNilVal = &v
}

func (this *paginator[Entity_T, Cursor_T]) Free() {

	this.Cursor = nil
}

func (this *paginator[Entity_T, Cursor_T]) HasCursor() bool {

	return this.Cursor != nil && (this.CursorNilVal == nil || *this.Cursor != *this.CursorNilVal)
}

func (this *paginator[Entity_T, Cursor_T]) SetOffset(offset uint64) {

	this.Offset = offset
}

func (this *paginator[Entity_T, Cursor_T]) SetSize(size uint64) {

	this.Size = size
}

func (this *paginator[Entity_T, Cursor_T]) ApplyFilter(fn repositoryAPI.FilterFunc) {

	this.IPaginationRepository.Filter(fn)
}

func (this *paginator[Entity_T, Cursor_T]) Select(fields ...string) {

	this.IPaginationRepository.Select(fields...)
}

func (this *paginator[Entity_T, Cursor_T]) ExcludeField(fields ...string) {

	this.IPaginationRepository.ExcludeFields(fields...)
}

func (this *paginator[Entity_T, Cursor_T]) CursorFirst() {

	this.isCursorFirst = true
}

func (this *paginator[Entity_T, Cursor_T]) SetCursorDirection(dir paginateServicePort.CursorDirection) {

	if dir == paginateServicePort.CURSOR_DIRECTION_PREVIOUS {

		this.IsPrev = true
	}
}
