package mongoDriver

// import (
// 	"app/internal/db/query"
// 	libCommon "app/internal/lib/common"
// 	libError "app/internal/lib/error"
// 	paginateServicePort "app/port/paginate"
// 	"context"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// const (
// 	SORT_ASC  = 1
// 	SORT_DESC = -1
// )

// type (
// 	paginate_err_msg struct {
// 		Operation string `json:"operation"`
// 		Status    string `json:"status,omitempty"`
// 		Message   string `json:"message"`
// 	}
// )

// type (
// 	PaginateExecutor[Entity_T any] struct {
// 		QueryExecutor[Entity_T]
// 	}
// )

// func (this *PaginateExecutor[Entity_T]) Paginate(ctx context.Context, fn query.PaginateInitFunc[interface{}]) (ret []Entity_T, err error) {

// 	// defer func() {

// 	// 	logLine := paginate_err_msg{
// 	// 		Operation: "db.Paginate",
// 	// 	}

// 	// 	switch {
// 	// 	case err != nil:
// 	// 		logLine.Status = "error"
// 	// 		logLine.Message = err.Error()
// 	// 	default:
// 	// 		logLine.Message = "success"
// 	// 	}

// 	// 	this.AccessLogger.PushTraceLogs(
// 	// 		ctx,
// 	// 		logLine,
// 	// 	)
// 	// }()

// 	var paginator paginateServicePort.IPaginator[interface{}]

// 	switch {
// 	case fn == nil:
// 		return nil, libError.NewInternal(
// 			fmt.Errorf("could not apply pagination query whose paginate init function is nil"),
// 		)
// 	}

// 	paginator = fn()

// 	if paginator == nil {

// 		return nil, libError.NewInternal(
// 			fmt.Errorf("could not apply pagination query, nil paginator given"),
// 		)
// 	}

// 	this.MongoAggregateQueryBuilder.PushStages(
// 		this.resolvePaginateOperator(paginator)...,
// 	)

// 	return this.exec(ctx)
// }

// func (this *PaginateExecutor[Entity_T]) resolvePaginateOperator(
// 	paginator paginateServicePort.IPaginator[interface{}],
// ) (ret []interface{}) {

// 	cursor_p, isCursor := (paginator).(paginateServicePort.ICursorPaginator[interface{}])
// 	nil_cursor_p, isNillableCursor := (paginator).(paginateServicePort.ICursorNillablePaginator[interface{}])

// 	var matchCursorQuery bool = isCursor && (cursor_p.GetCursor() != nil || isNillableCursor && cursor_p.GetCursor() != nil_cursor_p.CursorNilValue())

// 	pageNumber := paginator.GetPageNumber()

// 	if matchCursorQuery {

// 		return this.resovleCursorPaginateOperator(cursor_p)

// 	} else {

// 		ret = make([]interface{}, 2)

// 		ret[len(ret)-1] = bson.D{
// 			{"$limit", libCommon.Ternary(paginator.GetPageSize() < 1, 1, paginator.GetPageSize())},
// 		}

// 		ret[0] = bson.D{
// 			{"$skip", libCommon.Ternary(pageNumber <= 1, 0, pageNumber-1)},
// 		}
// 	}

// 	return
// }

// func (this *PaginateExecutor[Entity_T]) resovleCursorPaginateOperator(
// 	paginator paginateServicePort.ICursorPaginator[interface{}],
// ) (ret []interface{}) {

// 	isCursorPrevDir := paginator.GetCursorDirection() == paginateServicePort.CURSOR_DIRECTION_PREVIOUS

// 	if isCursorPrevDir {

// 		ret = make([]interface{}, 4)

// 		ret[1] = bson.D{
// 			{
// 				"$sort", bson.D{
// 					{"_id", SORT_DESC},
// 				},
// 			},
// 		}

// 		ret[len(ret)] = bson.D{
// 			{
// 				"$sort", bson.D{
// 					{"_id", SORT_ASC},
// 				},
// 			},
// 		}

// 	} else {

// 		ret = make([]interface{}, 3)

// 		ret[1] = bson.D{
// 			{
// 				"$sort", bson.D{
// 					{"_id", 1},
// 				},
// 			},
// 		}
// 	}

// 	ret[0] = bson.D{
// 		{
// 			"$match", bson.D{
// 				{
// 					"_id", bson.D{
// 						{libCommon.Ternary(isCursorPrevDir, "$lt", "$gt"), paginator.GetCursor()},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	ret[2] = bson.D{
// 		{
// 			"$limit", libCommon.Ternary(paginator.GetPageSize() < 1, 1, paginator.GetPageSize()),
// 		},
// 	}

// 	return
// }
