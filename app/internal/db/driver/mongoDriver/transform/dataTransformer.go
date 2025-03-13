package transform

import "go.mongodb.org/mongo-driver/bson"

type (
	QueryBuilderDataTransformer struct {
		setter
	}
)

func NewDataTransformer() *QueryBuilderDataTransformer {

	return new(QueryBuilderDataTransformer)
}

func (this *QueryBuilderDataTransformer) GetQuery() []interface{} {

	ret := make([]interface{}, 0)

	if len(this.SetterMap) > 0 {

		ret = append(
			ret,
			bson.D{
				{"$set", this.SetterMap},
			},
		)
	}

	return ret
}
