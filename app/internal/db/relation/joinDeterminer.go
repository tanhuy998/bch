package relation

type (
	IJoinOperator interface {
		ApplyJoinOperation()
	}

	IJoinDeterminerInitializer interface {
		AsInnerJoin() IJoinOperator
		AsLeftJoin() IJoinOperator
	}

	IDBRelationInitiatorJoinOperationDeterminer interface {
		DetermineJoinOperation(initializer IJoinDeterminerInitializer) IJoinOperator
	}
)
