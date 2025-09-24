package paginateUseCase

import (
	libCommon "app/internal/lib/common"
	paginateServicePort "app/port/paginate"
	opLog "app/unitOfWork/operationLog"
	"context"
)

const (
	PAGINATE_EXECUTOR_LOG_UNIT_STR = "paginate_executor"
)

type (
	log_unit struct {
		Op              string      `json:"operation"`
		Type            string      `json:"paginate_type"`
		Message         string      `json:"message"`
		CursorDirection string      `json:"cursor_direction,omitempty"`
		PageSize        uint64      `json:"page_size"`
		PageNumber      int64       `json:"page_number,omitempty"`
		Error           error       `json:"error,omitempty"`
		Cursor          interface{} `json:"cursor,omitempty"`
	}
)

type (
	logger_t struct {
		opLog.OperationLogger
	}
)

func (this *logger_t) logCursor(c interface{}, pageSize uint64, direction paginateServicePort.CursorDirection, err error, ctx context.Context) {

	l := log_unit{
		Op:              "paginate",
		Type:            "cursor",
		PageSize:        pageSize,
		Cursor:          c,
		Error:           err,
		CursorDirection: libCommon.Ternary(direction == paginateServicePort.CURSOR_DIRECTION_NEXT, "next", "previous"),
	}

	if err == nil {

		l.Message = "success"
	} else {

		l.Message = "failed"
	}

	// this.AccessLogger.PushTraceLogs(
	// 	ctx, l,
	// )

	this.Trace(PAGINATE_EXECUTOR_LOG_UNIT_STR).PushCustom(ctx, l)
}

func (this *logger_t) logOffset(pageNumber int64, pageSize uint64, err error, ctx context.Context) {

	l := log_unit{
		Op:         "paginate",
		Type:       "offset",
		PageSize:   pageSize,
		PageNumber: pageNumber,
		Error:      err,
	}

	if err == nil {

		l.Message = "success"
	} else {

		l.Message = "failed"
	}

	// this.AccessLogger.PushTraceLogs(
	// 	ctx, l,
	// )

	this.Trace(PAGINATE_EXECUTOR_LOG_UNIT_STR).PushCustom(ctx, l)
}
