package paginateUseCase

import (
	libCommon "app/internal/lib/common"
	paginateServicePort "app/port/paginate"
)

type (
	paginate_navigator_t[Data_T any /*ICursorEntity[Cursor_T]*/, Cursor_T comparable] struct {
		data                 []Data_T
		current_page_number  int64
		request_data_size    int64
		query_total_count    int64
		is_cursor_pagination bool
	}
)

func (this *paginate_navigator_t[Data_T, Cursor_T]) GetData() []Data_T {

	return this.data
}

func (this *paginate_navigator_t[Data_T, Cursor_T]) QueryDocumentCount() int64 {

	return this.query_total_count
}

func (this *paginate_navigator_t[Data_T, Cursor_T]) getCursorofIndex(n int) *Cursor_T {

	var targetElement Data_T

	length := len(this.data)

	switch {
	case length == 0:
		return nil
	case n > length:
		targetElement = this.data[length-1]
		// return libCommon.PointerPrimitive(
		// 	this.data[length-1].GetCursor(),
		// )
	default:
		targetElement = this.data[n]
		// return libCommon.PointerPrimitive(
		// 	this.data[n].GetCursor(),
		// )
	}

	switch v := any(&targetElement).(type) {
	case paginateServicePort.ICursorEntity[Cursor_T]:
		return libCommon.PointerPrimitive(v.GetCursor())
	}

	return nil
}

func (this *paginate_navigator_t[Data_T, Cursor_T]) GetFirstCursor() *Cursor_T {

	switch {
	case len(this.data) == 0:
		return nil
	default:
		return this.getCursorofIndex(0)
	}
}

func (this *paginate_navigator_t[Data_T, Cursor_T]) GetLastCursor() *Cursor_T {

	switch {
	case len(this.data) == 0:
		return nil
	default:
		return this.getCursorofIndex(len(this.data) - 1)
	}

}

func (this *paginate_navigator_t[Data_T, Cursor_T]) CurrentPageNumber() int64 {

	return this.current_page_number
}

func (this *paginate_navigator_t[Data_T, Cursor_T]) GetPageSize() uint64 {
	return uint64(len(this.data))
}

func (this *paginate_navigator_t[Data_T, Cursor_T]) GetRequestDataSize() uint64 {

	return uint64(this.request_data_size)
}
