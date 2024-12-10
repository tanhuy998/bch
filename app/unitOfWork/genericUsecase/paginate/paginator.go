package paginateUseCase

import repositoryAPI "app/repository/api"

type (
	IPaginationFilter repositoryAPI.FilterFunc

	ICursorPaginator interface {
		SetCursor(c interface{})
		CursorFirst()
	}

	IOffsetPaginator interface {
		SetOffset(offset uint64)
		SetSize(size uint64)
	}

	IPagaintorProjection interface {
		Select(fields ...string)
		ExcludeField(fields ...string)
	}

	IPaginatorFilter interface {
		Filter(fn repositoryAPI.FilterFunc)
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
		Skip          uint64
		Size          uint64
		Cursor        interface{}
		IsPrev        bool
		isCursorFirst bool
	}
)

func NewPaginator[Entity_T any](
	repo repositoryAPI.IPaginationRepository[Entity_T],
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

	this.Skip = offset
}

func (this *paginator[Entity_T]) SetSize(size uint64) {

	this.Skip = size
}

func (this *paginator[Entity_T]) Filter(fn repositoryAPI.FilterFunc) {

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
