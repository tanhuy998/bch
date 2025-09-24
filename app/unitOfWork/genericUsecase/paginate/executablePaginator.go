package paginateUseCase

type (
	executable_paginator[Entity_T any, Cursor_T comparable] struct {
		paginator_t[Entity_T, Cursor_T]
	}
)

func (this executable_paginator[Entity_T, Cursor_T]) HasCursor() bool {

	return this.paginator_t.HasCursor()
}

func (this executable_paginator[Entity_T, Cursor_T]) GetCursor() *Cursor_T {

	return this.paginator_t.Cursor
}

func (this executable_paginator[Entity_T, Cursor_T]) GetPageSize() uint64 {

	return this.paginator_t.Size
}

func (this executable_paginator[Entity_T, Cursor_T]) GetOffset() int64 {

	return this.paginator_t.Offset
}

func (this executable_paginator[Entity_T, Cursor_T]) IsPrevDirection() bool {

	return this.paginator_t.IsPrev
}

func (this executable_paginator[Entity_T, Cursor_T]) Free() {

	this.paginator_t.Free()
}

func (this *executable_paginator[Entity_T, Cursor_T]) get_paginator_initializer() IAbstractOptionableOffsetPaginatorInitializer {

	return &this.paginator_t
}
