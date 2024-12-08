package responseOutput

import (
	"app/valueObject/requestInput"
)

const (
	DEFAULT_PAGINATE_SIZE = 10
)

type (
	IResponseData[Data_t any] interface {
		SetData(Data_t)
		GetData() Data_t
	}

	ResponseData[Data_T any] struct {
		Data Data_T `json:"data"`
	}

	ResponseDataList[Data_T any] struct {
		Data                 []Data_T `json:"data"`
		*PaginationNavigator `json:"navigation,omitempty"`
		paginateSize         uint64 `json:"-"`
	}

	// NavigationQuery struct {
	// 	Cursor primitive.ObjectID `json:"p_pivot,omitempty"`
	// 	Limit  *int               `json:"p_limit,omitempty"`
	// 	IsPrev bool               `json:"p_prev,omitempty"`
	// }

	// PaginationNavigation struct {
	// 	CurrentPage int              `json:"page"`
	// 	Previous    *NavigationQuery `json:"previous,omitempty"`
	// 	Next        *NavigationQuery `json:"next,omitempty"`
	// }

	PaginationNav struct {
		requestInput.RangePaginateInput
		requestInput.MongoCursorPaginateInput
	}

	PaginationNavigator struct {
		Prev *PaginationNav `json:"previous,omitempty"`
		Next *PaginationNav `json:"next,omitempty"`
	}
)

func (this *ResponseDataList[Data_T]) SetData(d []Data_T) {

	this.Data = d
}

func (this *ResponseDataList[Data_T]) GetData() []Data_T {

	return this.Data
}

func (this *ResponseDataList[Data_T]) ResolvePaginateNavigator(pageSize uint64) {

	if pageSize == 0 {

		panic("pagination size must be at least 1")
	}

	this.paginateSize = pageSize

	this.ResolveNavNext()
	this.ResolveNavPrev()
}

func (this *ResponseDataList[Data_T]) ResolveNavNext() {

	if len(this.Data) == 0 {

		this.Next = nil
		return
	}

	if this.Next == nil {

		this.Next = new(PaginationNav)
	}

	lastIndex := len(this.Data) - 1

	if lastIndex <= 0 {

		return
	}

	lastElement := this.Data[lastIndex]

	if v, ok := any(lastElement).(*requestInput.MongoCursorPaginateInput); ok {

		this.Next.ObjectID = v.GetCursor()
	}

	this.Next.PageSize = uint64(len(this.Data))
	this.Next.IsPrev = false
}

func (this *ResponseDataList[Data_T]) ResolveNavPrev() {

	if len(this.Data) == 0 {

		this.Prev = nil
		return
	}

	if this.Prev == nil {

		this.Prev = new(PaginationNav)
	}

	firstElement := this.Data[0]

	if v, ok := any(firstElement).(*requestInput.MongoCursorPaginateInput); ok {

		this.Next.ObjectID = v.GetCursor()
	}

	this.Prev.PageSize = uint64(len(this.Data))
	this.Next.IsPrev = false
}
