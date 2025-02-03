package transform

import "go.mongodb.org/mongo-driver/bson"

type (
	data_transformer struct {
		setter
	}
)

func NewDataTransformer() *data_transformer {

	return new(data_transformer)
}

func (this *data_transformer) GetQuery() []interface{} {

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
