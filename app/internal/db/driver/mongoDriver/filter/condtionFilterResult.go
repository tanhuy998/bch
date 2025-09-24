package filter

import (
	"go.mongodb.org/mongo-driver/bson"
)

type (
	ConditionFilterResult struct {
		ConditionFilterGenerator
		//filter_t
		refConditionExpression *bson.D
	}
)

func NewConditionFilterResult(refCondtionExpression *bson.D) *ConditionFilterResult {

	ret := new(ConditionFilterResult)

	ret.refConditionExpression = refCondtionExpression

	//*refCondtionExpression = ret.filter_t.AsBson()

	return ret
}

func (this *ConditionFilterResult) ApplyConditionExpression() {

	//*this.refConditionExpression = (bson.D)(this.ConditionFilterGenerator.FilterGenerator)
	*this.refConditionExpression = this.ConditionFilterGenerator.filter_t.AsBson()
}
