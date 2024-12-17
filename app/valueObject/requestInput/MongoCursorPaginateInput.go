package requestInput

import (
	libCommon "app/internal/lib/common"
	paginateServicePort "app/port/paginate"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	MongoCursorPaginateInput struct {
		ObjectID primitive.ObjectID `url:"p_cursor"`
		IsPrev   bool               `url:"p_prev"`
	}
)

func (this MongoCursorPaginateInput) GetCursor() primitive.ObjectID /**tv(Cursor_Type)*/ {

	return this.ObjectID
}

func (this MongoCursorPaginateInput) IsPrevious() bool {

	return this.IsPrev
}

func (this MongoCursorPaginateInput) HasCursor() bool {

	return this.ObjectID != primitive.NilObjectID
}

func (this MongoCursorPaginateInput) CursorNilValue() primitive.ObjectID {

	return primitive.NilObjectID
}

func (this MongoCursorPaginateInput) GetCursorDirection() paginateServicePort.CursorDirection {

	return libCommon.Ternary(this.IsPrev, paginateServicePort.CURSOR_DIRECTION_PREVIOUS, paginateServicePort.CURSOR_DIRECTION_NEXT)
}
