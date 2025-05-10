package determiner

import "app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder"

type (
	inner_join struct {
		dispatcher *queryBuilder.JoinOperationDispatcher
	}
)

func (this inner_join) ApplyJoinOperation() {

	this.dispatcher.AsInnerJoin()
}
