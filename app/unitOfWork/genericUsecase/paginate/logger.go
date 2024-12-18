package paginateUseCase

import (
	libCommon "app/internal/lib/common"
	paginateServicePort "app/port/paginate"
	opLog "app/unitOfWork/operationLog"
	"context"
)

type (
	log_unit struct {
		Op              string      `json:"operation"`
		Type            string      `json:"paginate_type"`
		Message         string      `json:"message"`
		CursorDirection string      `json:"cursor_direction,omitempty"`
		PageSize        uint64      `json:"page_size"`
		PageNumber      uint64      `json:"page_number,omitempty"`
		Error           error       `json:"error,omitempty"`
		Cursor          interface{} `json:"cursor,omitempty"`
	}

	logger struct {
		opLog.OperationLogger
	}
)

func (this *logger) logCursor(c interface{}, pageSize uint64, direction paginateServicePort.CursorDirection, err error, ctx context.Context) {

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

	this.AccessLogger.PushTraceLogs(
		ctx, l,
	)
}

func (this *logger) logOffset(pageNumber uint64, pageSize uint64, err error, ctx context.Context) {

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

	this.AccessLogger.PushTraceLogs(
		ctx, l,
	)
}
