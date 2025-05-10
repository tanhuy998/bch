package determiner

import "app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder"

type (
	left_join struct {
		dispatcher *queryBuilder.JoinOperationDispatcher
	}
)

func (this left_join) ApplyJoinOperation() {

	this.dispatcher.AsLeftJoin()
}
