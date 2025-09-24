package comparision

import (
	"app/internal/db/driver/mongoDriver/filter/operator"
	"fmt"
)

type (
	/*
		abstract_comparison_t implements expression.IBinaryExpression
	*/
	abstract_comparision_t struct {
		operator.SecondaryOperator
	}
)

func (this *abstract_comparision_t) prepare() {

	this.SetLeftOperand(
		fmt.Sprintf(`$%s`, this.Root().CurrentFieldName()),
	)
}
