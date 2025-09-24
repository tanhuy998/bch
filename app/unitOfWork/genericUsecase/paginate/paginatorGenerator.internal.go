package paginateUseCase

import (
	"app/internal/db/query"
	paginateServicePort "app/port/paginate"
	repositoryAPI "app/repository/api"
	"app/unitOfWork/genericUsecase/paginate/condition"
	"strconv"

	"github.com/google/uuid"
)

const (
	PAGINATOR_GENERATOR_LOG_UNIT_STR = "paginator_generator"
)

type (
	internal_executable_paginator_generator_t[
		Repository_T repositoryAPI.IPaginateClonableRepository[Entity_T],
		Entity_T any,
		Cursor_T comparable,
	] struct {
		logger_t
		Repository Repository_T
	}
)

func (this internal_executable_paginator_generator_t[Repository_T, Entity_T, Cursor_T]) generateExecutablePaginatorByDefault(
	tenantUUID uuid.UUID, options []PaginationOption[Cursor_T],
) *executable_paginator[Entity_T, Cursor_T] /* IExecutablePaginator[Entity_T, Cursor_T] */ {

	ret := &executable_paginator[Entity_T, Cursor_T]{
		paginator_t: *NewPaginator[Entity_T, Cursor_T](this.Repository),
	}

	defer this.Debug(PAGINATOR_GENERATOR_LOG_UNIT_STR).Messure("generate_default_exe_paginator", "", ret.log)(nil)

	paginator := &ret.paginator_t

	for _, fn := range options {

		fn(paginator)
	}

	paginator.IPaginationUnit.Filter(
		func(filter repositoryAPI.IFilterGenerator) {

			filter.Field("tenantUUID").Equal(tenantUUID)
		},
	)

	return ret
}

func (this internal_executable_paginator_generator_t[Repository_T, Entity_T, Cursor_T]) generateExecutablePaginatorByCustomPaginator(
	tenantUUID uuid.UUID, requestedCustomPaginator paginateServicePort.IPaginator[Cursor_T],
) /* *executable_paginator[Entity_T, Cursor_T] */ IExecutablePaginator[Entity_T] {

	switch v := this.tryGeneralExecutablePaginatorAsGeneralPaginator(tenantUUID, requestedCustomPaginator); {
	case v != nil:

		return v
	}

	ret := &executable_paginator[Entity_T, Cursor_T]{
		// paginator: paginator[Entity_T, Cursor_T]{
		// 	IPaginationUnit: this.Repository.Clone(),
		// },
		paginator_t: *NewPaginator[Entity_T, Cursor_T](this.Repository),
	}

	debug := this.Debug(PAGINATOR_GENERATOR_LOG_UNIT_STR)

	// _p := paginator[Entity_T, Cursor_T]{
	// 	IPaginationRepository: this.Repository.Clone(),
	// }

	_p := &ret.paginator_t

	//_p := ret.get_paginator_initializer()

	defer debug.Messure("generate_default_custom_exe_paginator", "", _p.getLogRecorder())(nil)

	// _p.SetOffset(customPaginator.GetPageNumber())
	// debug.Push("set_paginate_offset", "", _p.getLogRecorder())

	// _p.SetSize(customPaginator.GetPageSize())
	// debug.Push("set_paginate_page_size", "", _p.getLogRecorder())

	this.initOffsetPagination(requestedCustomPaginator, ret)

	this.resolveQueryFilter(tenantUUID, requestedCustomPaginator, _p.IPaginationUnit)

	debug.Push("resolve_query_condition", "", ret.log)

	// _p.IPaginationUnit.Filter(
	// 	func(filter repositoryAPI.IFilterGenerator) {

	// 		if v, ok := (customPaginator).(IFilterablePaginator); ok {

	// 			v.ApplyPaginateFilter(filter)
	// 		}

	// 		filter.Field("tenantUUID").Equal(tenantUUID)
	// 	},
	// )

	if v, ok := (requestedCustomPaginator).(paginateServicePort.IProjectionPaginator); ok {

		v.ApplyPaginateProjection(_p)

		debug.Push("apply_query_projection", "", ret.log)
	}

	if v, ok := (requestedCustomPaginator).(paginateServicePort.ICursorNillablePaginator[Cursor_T]); ok {

		if v.CursorNilValue() != nil {

			_p.SetCursorNilValue(*v.CursorNilValue())

			debug.Push("paginate_cursor_nil_value_acknowledgement", "", ret.log)
		}
	}

	// switch v := customPaginator.(type) {
	// case paginateServicePort.ICursorPaginator[Cursor_T]:
	// 	if v.GetCursor() != nil {

	// 		_p.SetCursor(*v.GetCursor())
	// 		_p.SetCursorDirection(v.GetCursorDirection())
	// 		_p.CursorFirst()

	// 		debug.Push("apply_cursor_pagination", "", ret.log)
	// 	}
	// case paginateServicePort.IGeneralPaginator:

	// 	generalPaginator := v.GetGeneralPaginator()

	// 	if generalPaginator.GetCursor() != nil {

	// 		_p.SetCursor(generalPaginator.GetCursor())
	// 		_p.SetCursorDirection(v.GetCursorDirection())
	// 		_p.CursorFirst()

	// 		debug.Push("apply_cursor_pagination", "", ret.log)
	// 	}

	// }

	if v, ok := (requestedCustomPaginator).(paginateServicePort.ICursorPaginator[Cursor_T]); ok {

		if v.GetCursor() != nil {

			_p.SetCursor(*v.GetCursor())
			_p.SetCursorDirection(v.GetCursorDirection())
			_p.CursorFirst()

			debug.Push("apply_cursor_pagination", "", ret.log)
		}

	}

	return ret
}

func (this internal_executable_paginator_generator_t[Repository_T, Entity_T, Cursor_T]) tryGeneralExecutablePaginatorAsGeneralPaginator(
	tenantUUID uuid.UUID, requestedCustomPaginator paginateServicePort.IPaginator[Cursor_T],
) IExecutablePaginator[Entity_T] {

	switch _p := requestedCustomPaginator.(type) {
	case paginateServicePort.IGeneralPaginator:

		ret := &general_paginator_t[Entity_T]{
			paginator_t:     *NewPaginator[Entity_T, interface{}](this.Repository),
			customPaginator: _p,
		}

		this.initOffsetPagination(requestedCustomPaginator, ret)

		this.Debug(PAGINATOR_GENERATOR_LOG_UNIT_STR).Push(
			"generate_exec_paginator_by_general_custom_paginator", "", ret.GetLogRecorder(),
		)

		return ret
	default:
		return nil
	}
}

func (this internal_executable_paginator_generator_t[Repository_T, Entity_T, Cursor_T]) initOffsetPagination(
	requestedCustomOffsetPaginator paginateServicePort.IOffsetPaginator, paginator IExecutablePaginator[Entity_T],
) {

	_p := paginator.get_paginator_initializer()

	debug := this.Debug(PAGINATOR_GENERATOR_LOG_UNIT_STR)

	_p.SetOffset(requestedCustomOffsetPaginator.GetPageNumber())
	debug.Push(
		"set_paginate_offset", strconv.FormatUint(uint64(requestedCustomOffsetPaginator.GetPageNumber()), 10), _p.getLogRecorder(),
	)

	_p.SetSize(requestedCustomOffsetPaginator.GetPageSize())
	debug.Push(
		"set_paginate_page_size", strconv.FormatUint(requestedCustomOffsetPaginator.GetPageSize(), 10), _p.getLogRecorder(),
	)
}

func (this internal_executable_paginator_generator_t[Repository_T, Entity_T, Cursor_T]) resolveQueryFilter(
	tenantUUID uuid.UUID, customPaginator paginateServicePort.IPaginator[Cursor_T], paginationUnit repositoryAPI.IPaginationUnit[Entity_T],
) {

	conditionUnit := paginationUnit.QueryConditionUnit()

	conditionUnit.AsConditionExpression(
		func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

			return expression.Logical().And(
				func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

					return expression.Filter(
						func(filter query.IDataConditionFilterExpression) {

							filter.Field("tenantUUID").Equal(tenantUUID)
						},
					)
				},
				func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

					switch v := customPaginator.(type) {
					case condition.IQueryConditionPaginator:
						return v.MatchCondition(expression)
					default:
						return nil
					}
				},
			)
		},
	)
}
