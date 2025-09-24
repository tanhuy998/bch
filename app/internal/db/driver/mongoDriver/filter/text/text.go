package text

import (
	"app/internal/db/driver/mongoDriver/filter/operator"
	"app/internal/db/query"
	libCommon "app/internal/lib/common"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	FILTER_TEXT_CASTING_PANIC_LABEL = "text_casting_filter_t"
)

type (
	filter_expression_text_casting_t struct {
		text_default_operator_t
	}
)

func regex_option(option string) *string {

	return libCommon.PointerPrimitive(option)
}

func NewFilterFieldTextCasting(
	operator *operator.SecondaryOperator,
) *filter_expression_text_casting_t {

	switch {
	case operator == nil:
		panic("operator of filter_expression_text_casting_t must not be nil")
	}

	ret := &filter_expression_text_casting_t{}

	ret.default_text_encoding_regex_option = query.RegexOption("u")

	ret.SecondaryOperator = operator

	ret.SetAssertedType(bson.TypeString)

	ret.Init()

	return ret
}

func (this *filter_expression_text_casting_t) ConcatWith(vals ...interface{}) query.IFilterTextOperator {

	concatList := make([]interface{}, len(vals)+1)

	concatList[0] = fmt.Sprintf(`$%s`, this.Root().CurrentFieldName())

	for i, inputValue := range vals {

		if i == 0 {
			continue
		}

		switch v := this.AssertFieldIfSelfReference(inputValue); {
		case v != "":
			concatList[i+1] = v
		default:
			concatList[i+1] = vals
		}
	}

	parent := this.OverrideExpression().AsBinaryExpression()

	parent.SetLeftOperand(
		bson.D{
			{"$concat", concatList},
		},
	)

	return this
}
