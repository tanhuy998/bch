package mongoRepository

import (
	"app/internal/db/driver/mongoDriver/condition/expression"
	"app/internal/db/driver/mongoDriver/filter"
	repositoryAPI "app/repository/api"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	err_condition_default = fmt.Errorf("(repository condition materializer error)")
)

type (
	filter_condition_materializer_t struct {
		f bson.D
	}
)

func (this *filter_condition_materializer_t) getFilterExpression() (bson.D, error) {

	// switch val := this.f.(type) {
	// case bson.D:
	// 	return val, nil
	// case nil:
	// 	return bson.D{}, nil
	// default:
	// 	fmt.Printf(`%t`, val)
	// 	return nil, libError.NewInternal(
	// 		err_condition_default,
	// 		fmt.Errorf("Invalid evaluation of repository condition expression (expression must be type of bson.D)"),
	// 	)
	// }

	switch this.f {
	case nil:
		return bson.D{}, nil
	default:
		return this.f, nil
	}
}

func (this *filter_condition_materializer_t) AsFilter(
	fn repositoryAPI.FilterFunc,
) {

	filter := filter.NewFilterGenerator()

	fn(filter)

	this.f = filter.AsBson()
}

func (this *filter_condition_materializer_t) AsConditionExpression(
	fn repositoryAPI.MatchFunc,
) {

	initializer := expression.NewConditionExpressionInitializer(&this.f)

	result := fn(initializer)

	if result == nil {

		return
	}

	result.ApplyConditionExpression()
}
