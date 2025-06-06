package filter

import "go.mongodb.org/mongo-driver/bson"

type (
	ConditionFilterResult struct {
		ConditionFilterGenerator
		refConditionExpression *bson.D
	}
)

func NewConditionFilterResult(refCondtionExpression *bson.D) *ConditionFilterResult {

	ret := new(ConditionFilterResult)

	ret.refConditionExpression = refCondtionExpression

	return ret
}

func (this *ConditionFilterResult) ApplyConditionExpression() {

	*this.refConditionExpression = (bson.D)(this.ConditionFilterGenerator.FilterGenerator)
}
