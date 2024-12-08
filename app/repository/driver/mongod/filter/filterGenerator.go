package mongoRepositoryFilter

import (
	repositoryAPI "app/repository/api"
)

type (
	MongoRepositoryFilterGenerator []interface{}
)

func (this *MongoRepositoryFilterGenerator) reset() {

}

func (this *MongoRepositoryFilterGenerator) init() {

	if *this == nil {

		*this = make([]interface{}, 0)
	}
}

func (this *MongoRepositoryFilterGenerator) Add(exprs ...interface{}) repositoryAPI.IFilterGenerator {

	this.init()

	*this = append(*this, exprs...)

	return this
}

func (this *MongoRepositoryFilterGenerator) Get() []interface{} {

	return *this
}

func (this *MongoRepositoryFilterGenerator) Field(name string) repositoryAPI.IFilterExpressionOperator {

	return &mongo_filter_expr{
		ref: *this,
		lhs: name,
	}
}
