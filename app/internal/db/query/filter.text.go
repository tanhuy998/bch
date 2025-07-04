package query

type (
	RegexOption string

	TextOption string

	IFilterTextOperator interface {
		IGenericNegationOperator[IFilterTextOperator]
		Equal(str interface{})
		//Like(subStr string)
		Contains(subStr interface{})
		BeginsWith(subStr interface{})
		EndsWith(subStr interface{})
	}
)
