package paginateServicePort

type (
	IPaginateContext[Cursor_T comparable] interface {
		GetPaginator() IPaginator[Cursor_T]
	}
)
