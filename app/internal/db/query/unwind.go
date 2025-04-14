package query

type (
	IUnwindJoinedData interface {
		UnwindForeign() IQueryBuilder
	}

	IJoinUnwindableQueryBuilder interface {
		IQueryBuilder
		IUnwindJoinedData
	}
)
