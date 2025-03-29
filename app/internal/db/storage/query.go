package storage

type (
	IQueryEndingPhase interface {
		Done()
	}

	IArbitraryQuery interface {
		IQueryEndingPhase
		GetArbitraryQuery() interface{}
	}
)
