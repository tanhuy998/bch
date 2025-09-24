package text

import (
	"app/internal/db/driver/mongoDriver/filter/operator"
	"app/internal/db/query"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	text_default_operator_t struct {
		*operator.SecondaryOperator        //abstract_casting_expression_t
		default_text_encoding_regex_option query.RegexOption
	}
)

func (this *text_default_operator_t) BeginsWith(subStr interface{}) {

	regexPattern_str := `/^%s/`

	switch selfRef := this.AssertFieldIfSelfReference(subStr); {
	case selfRef != "":
		regexPattern_str = fmt.Sprintf(regexPattern_str, selfRef)
	default:
		regexPattern_str = fmt.Sprintf(regexPattern_str, subStr)
	}

	this._resolveRegexForCurrentField(
		regexPattern_str, nil,
	)
}

func (this *text_default_operator_t) EndsWith(subStr interface{}) {

	regexPattern_str := `/%s$/`

	switch selfRef := this.AssertFieldIfSelfReference(subStr); {
	case selfRef != "":
		regexPattern_str = fmt.Sprintf(regexPattern_str, selfRef)
	default:
		regexPattern_str = fmt.Sprintf(regexPattern_str, subStr)
	}

	this._resolveRegexForCurrentField(
		regexPattern_str, nil,
	)
}

func (this *text_default_operator_t) Not() query.IFilterTextOperator {

	this.PrimaryOperator.Negate()

	return this
}

func (this *text_default_operator_t) _resolveRegexForCurrentField(
	pattern string, option *string,
) {

	parent := this.OverrideExpression().AsUnaryExpression()

	parent.SetOperator("$regexMatch")
	parent.SetOperand(
		bson.D{
			{"input", fmt.Sprintf(`$%s`, this.Root().CurrentFieldName())},
			{"regex", pattern},
			{"options", option},
		},
	)
}

func (this *text_default_operator_t) Equal(str interface{}) {

	parent := this.OverrideExpression().AsBinaryExpression()

	switch innnerFieldName := this.AssertFieldIfSelfReference(str); {
	case innnerFieldName != "":
		parent.SetRightOperand(innnerFieldName)
	default:
		parent.SetRightOperand(str)
	}

	parent.SetOperator("$eq")
}

func (this *text_default_operator_t) MatchRegex(strPattern string, options ...query.RegexOption) {

	if strPattern == "" {

		panic(fmt.Sprintf("%s: string passed to method ByRegex() must not be empty", FILTER_TEXT_CASTING_PANIC_LABEL))
	}

	var opts = string(this.default_text_encoding_regex_option)

	for _, v := range options {

		opts += string(v)
	}

	this._resolveRegexForCurrentField(
		strPattern, regex_option(opts),
	)
}

func (this *text_default_operator_t) Contains(subStr interface{}) {

	regexPattern_str := `/%s/`

	switch selfRef := this.AssertFieldIfSelfReference(subStr); {
	case selfRef != "":
		regexPattern_str = fmt.Sprintf(regexPattern_str, selfRef)
	default:
		regexPattern_str = fmt.Sprintf(regexPattern_str, subStr)
	}

	this._resolveRegexForCurrentField(
		regexPattern_str, nil,
	)
}
