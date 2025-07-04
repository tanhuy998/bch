package query

type (
	IFilterDateOperator interface {
		IGenericNegationOperator[IFilterDateOperator]
		Before() IFilterDateExtractionOperator
		After() IFilterDateExtractionOperator
		Equal() IFilterDateExtractionOperator
	}

	IFilterDateExtractionOperator interface {
		Date(date interface{})
		YearOf(date interface{})
		MonthOf(date interface{})
		DateOf(date interface{})
	}
)
