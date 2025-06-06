package expression

import (
	"app/internal/db/driver/mongoDriver/filter"
	"app/internal/db/query"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	ConditionExpressionInitializer struct {
		//logical_expression_generator
		ref *bson.D
	}
)

func NewConditionExpressionInitializer(ref *bson.D) *ConditionExpressionInitializer {

	return &ConditionExpressionInitializer{ref}
}

// func (this *ConditionExpressionInitializer) init() {

// 	b := (*bson.D)(this)

// 	if *b == nil {

// 		*b = make(bson.D, 1)
// 	}
// }

// func (this *ConditionExpressionInitializer) AsLogical(
// 	fn query.LogicalExpressionInitFunc,
// ) query.IDataConditionExpressionResult {

// 	if fn == nil {

// 		panic("condition expresstion logical init funciton must not be nil")
// 	}

// 	logicalExprGenerator := NewLogicalExpressionGenerator()

// 	logicalResult := fn(logicalExprGenerator)

// 	return logicalResult
// }

func (this *ConditionExpressionInitializer) Filter(
	fn query.FilterExpressInitFunc,
) query.IDataConditionExpressionResult {

	if fn == nil {

		panic("condtion expresstion filter init funciton must not be nil")
	}

	// conditionFilterGenerator := filter.NewConditionFilterGenerator()

	// fn(conditionFilterGenerator)

	// return conditionFilterGenerator

	filter := filter.NewConditionFilterResult(this.ref)

	fn(filter)

	return filter
}

func (this *ConditionExpressionInitializer) Logical() query.INegatableLogicalOperator {

	return NewLogicalExpressionGenerator(this.ref)
}
