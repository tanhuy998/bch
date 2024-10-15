package requestInput

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	MongoCursorPaginateInput struct {
		ObjectID primitive.ObjectID `url:"p_cursor"`
		IsPrev   bool               `url:"p_prev"`
	}
)

func (this *MongoCursorPaginateInput) GetCursor() primitive.ObjectID /**tv(Cursor_Type)*/ {

	return this.ObjectID
}

func (this *MongoCursorPaginateInput) IsPrevious() bool {

	return this.IsPrev
}

func (this *MongoCursorPaginateInput) HasCursor() bool {

	return this.ObjectID != primitive.NilObjectID
}
