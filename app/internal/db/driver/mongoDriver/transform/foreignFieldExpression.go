package transform

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	foreign_field_expression struct {
		transformer *RelationDataTransformer
	}
)

func (this *foreign_field_expression) FirstElement() {

	this.transformer.Value(
		bson.D{
			{
				"$arrayElemAt", bson.A{
					//"$commandGroup", 0,
					fmt.Sprintf("$%s", this.transformer.cur), 0,
				},
			},
		},
	)
}

func (this *foreign_field_expression) At(index uint64) {

	this.transformer.Value(
		bson.D{
			{
				"$arrayElemAt", bson.A{
					//"$commandGroup", 0,
					fmt.Sprintf("$%s", this.transformer.cur), index,
				},
			},
		},
	)
}

func (this *foreign_field_expression) Count() {

	this.transformer.Value(
		bson.D{
			{"$count", this.transformer.cur},
		},
	)
}
