package paginateUseCase

import (
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
	opLog "app/unitOfWork/operationLog"
	"context"
)

type (
	paginator_t[Entity_T any, Cursor_T comparable] struct {
		repositoryAPI.IPaginationUnit[Entity_T]
		log           *opLog.LogRecorder
		Cursor        *Cursor_T // interface{}
		CursorNilVal  *Cursor_T //interface{}
		Offset        int64
		Size          uint64
		IsPrev        bool
		isCursorFirst bool
	}
)

func NewPaginator[Entity_T any, Cursor_T comparable](
	repo repositoryAPI.IPaginateClonableRepository[Entity_T],
) *paginator_t[Entity_T, Cursor_T] {

	// if repo == nil {

	// 	panic("could not initialized new paginator, nil repository given")
	// }

	ret := &paginator_t[Entity_T, Cursor_T]{
		IPaginationUnit: repo.Clone(),
	}

	ret.log = opLog.NewLogContext(context.TODO())

	return ret
}

func (this *paginator_t[Entity_T, Cursor_T]) getLogRecorder() *opLog.LogRecorder {

	return this.log
}

func (this *paginator_t[Entity_T, Cursor_T]) SetCursor(c Cursor_T) {

	this.Cursor = &c
}

func (this *paginator_t[Entity_T, Cursor_T]) SetCursorNilValue(v Cursor_T) {

	this.CursorNilVal = &v
}

func (this *paginator_t[Entity_T, Cursor_T]) Free() {

	this.Cursor = nil
}

func (this *paginator_t[Entity_T, Cursor_T]) HasCursor() bool {

	return this.Cursor != nil && (this.CursorNilVal == nil || *this.Cursor != *this.CursorNilVal)
}

func (this *paginator_t[Entity_T, Cursor_T]) SetOffset(offset int64) {

	this.Offset = offset
}

func (this *paginator_t[Entity_T, Cursor_T]) SetSize(size uint64) {

	this.Size = size
}

func (this *paginator_t[Entity_T, Cursor_T]) ApplyFilter(fn repositoryAPI.FilterFunc) {

	this.IPaginationUnit.Filter(fn)
}

func (this *paginator_t[Entity_T, Cursor_T]) Select(fields ...string) {

	this.IPaginationUnit.Select(fields...)
}

func (this *paginator_t[Entity_T, Cursor_T]) ExcludeField(fields ...string) {

	this.IPaginationUnit.ExcludeFields(fields...)
}

func (this *paginator_t[Entity_T, Cursor_T]) CursorFirst() {

	this.isCursorFirst = true
}

func (this *paginator_t[Entity_T, Cursor_T]) SetCursorDirection(dir paginateServicePort.CursorDirection) {

	if dir == paginateServicePort.CURSOR_DIRECTION_PREVIOUS {

		this.IsPrev = true
	}
}

func (this paginator_t[Entity_T, Cursor_T]) GetLogRecorder() *opLog.LogRecorder {

	return this.log
}
