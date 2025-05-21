package paginateOutput

import (
	libCommon "app/internal/lib/common"
	paginateUseCase "app/unitOfWork/genericUsecase/paginate"
	"reflect"
)

type (
	OptionableQueryPaginateData[Data_T paginateUseCase.ICursorEntity[interface{}], Optional_Query_T comparable] struct {
		Data                []Data_T                                    `json:"data"`
		CurrentPage         int64                                       `json:"currentPage"`
		PageSize            uint64                                      `json:"pageSize"`
		TotalCount          int64                                       `json:"totalCount"`
		Navigator           PaginateNavigator[Data_T, Optional_Query_T] `json:"navigator,omitempty"`
		optionalQuery       Optional_Query_T                            `json:"-"`
		ignoreOptionalQuery bool                                        `json:"-"`
	}
)

func (this *OptionableQueryPaginateData[Data_T, Optional_Query_T]) init(
	navigator paginateUseCase.INavigator[Data_T, interface{}],
) {

	this.CurrentPage = navigator.CurrentPageNumber()
	this.CurrentPage = libCommon.Ternary(this.CurrentPage == 0, 1, this.CurrentPage)

	this.Data = navigator.GetData()
	this.PageSize = navigator.GetPageSize()
	this.TotalCount = navigator.QueryDocumentCount()

	this.Navigator.init(navigator)

	this.linkOptionQueryWithNavigators()
}

func (this *OptionableQueryPaginateData[Data_T, Optional_Query_T]) linkOptionQueryWithNavigators() {

	if this.ignoreOptionalQuery {

		this.unlinkOptionQuery()
		return
	}

	var (
		prev = this.Navigator.Previos
		next = this.Navigator.Next
	)

	switch {
	case prev != nil:
		prev.OptionalQuery = &this.optionalQuery
		fallthrough
	case next != nil:
		next.OptionalQuery = &this.optionalQuery
	}
}

func (this *OptionableQueryPaginateData[Data_T, Optional_Query_T]) unlinkOptionQuery() {

	var (
		prev = this.Navigator.Previos
		next = this.Navigator.Next
	)

	switch {
	case prev != nil:
		prev.OptionalQuery = nil
		fallthrough
	case next != nil:
		next.OptionalQuery = nil
	}
}

func (this *OptionableQueryPaginateData[Data_T, Optional_Query_T]) ensureOptionalQueryTypeIsStruct() {

	reflection := reflect.TypeFor[Optional_Query_T]()

	switch {
	case reflection == reflect.TypeFor[ignore_optional_query_t]():
		this.ignoreOptionalQuery = true
		return
	case reflection.Kind() != reflect.Struct:
		panic("Optional_Query_T passed as type parameter to PaginateData must be struct type")
	}
}

func (this *OptionableQueryPaginateData[Data_T, Optional_Query_T]) SetData(
	navigator paginateUseCase.INavigator[Data_T, interface{}],
) {

	this.ensureOptionalQueryTypeIsStruct()
	this.init(navigator)
}

func (this *OptionableQueryPaginateData[Data_T, Optional_Query_T]) GetOptionQuery() *Optional_Query_T {

	return &this.optionalQuery
}
