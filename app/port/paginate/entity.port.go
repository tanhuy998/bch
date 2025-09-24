package paginateServicePort

type (
	ICursorEntity[Cursor_T comparable] interface {
		GetCursor() Cursor_T
	}
)
