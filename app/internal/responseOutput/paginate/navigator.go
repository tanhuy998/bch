package paginateOutput

import (
	libCommon "app/internal/lib/common"
	paginateUseCase "app/unitOfWork/genericUsecase/paginate"
)

type (
	PaginateNavigator[Data_T paginateUseCase.ICursorEntity[interface{}], Optional_Query_T any] struct {
		refNavigator paginateUseCase.INavigator[Data_T, interface{}]
		Previos      *NavigatorDir[Optional_Query_T] `json:"previous,omitempty"`
		Next         *NavigatorDir[Optional_Query_T] `json:"next,omitempty"`
	}
)

func (this *PaginateNavigator[Data_T, Optional_Query_T]) init(
	navigator paginateUseCase.INavigator[Data_T, interface{}],
) {

	this.refNavigator = navigator

	this.resovlePrev()

	if this.Previos == nil {

		return
	}

	this.resolveNext()
}

func (this *PaginateNavigator[Data_T, Optional_Query_T]) resolveNext() {

	var (
		dataLength = len(this.refNavigator.GetData())
		lastCursor = this.refNavigator.GetLastCursor()
		offset     = this.refNavigator.CurrentPageNumber()
	)

	switch {
	case dataLength == 0:
		return
	case dataLength < int(this.refNavigator.GetPageSize()):
		return
	default:
		this.Next = &NavigatorDir[Optional_Query_T]{
			Cursor:   lastCursor,
			PageSize: this.refNavigator.GetPageSize(),
			Offset:   libCommon.Ternary(offset < 0, 0, offset),
		}
		return
	}
}

func (this *PaginateNavigator[Data_T, Optional_Query_T]) resovlePrev() {

	var (
		dataLength  = len(this.refNavigator.GetData())
		firstCursor = this.refNavigator.GetFirstCursor()
		offset      = this.refNavigator.CurrentPageNumber()
	)

	switch {
	case dataLength == 0:
		return
	default:
		this.Next = &NavigatorDir[Optional_Query_T]{
			Cursor:   firstCursor,
			PageSize: this.refNavigator.GetPageSize(),
			Offset:   libCommon.Ternary(offset < 0, 0, offset),
		}
		return
	}
}

// func (this *PaginateNavigator[Data_T]) SetOptionalQuery(
// 	fn OptionalQuerySetFunc[Optional_Url_Query_T],
// ) {

// 	if fn == nil {

// 		return
// 	}

// 	switch {
// 	case this.Previos != nil:
// 		this.Previos.OptionalQuery = optionalQuery
// 		fallthrough
// 	case this.Next != nil:
// 		this.Next.OptionalQuery = optionalQuery
// 	}
// }
