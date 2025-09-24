package date

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	self_ref_time_comparator_t struct {
		current_field string
		ref_field     string
		aspect        string
		m             byte
	}
)

func (this *self_ref_time_comparator_t) includeDay() {

	this.includeMonth()
	this.m |= 4
}

func (this *self_ref_time_comparator_t) includeMonth() {

	this.includeYear()
	this.m |= 2
}

func (this *self_ref_time_comparator_t) includeYear() {

	this.m |= 1
}

func (this *self_ref_time_comparator_t) MarshalBSON() ([]byte, error) {

	var ret bson.A

	ref_currentField := fmt.Sprintf("$%s", this.current_field)

	switch {
	case this.m&4 != 0:
		ret = bson.A{
			bson.D{
				{
					"$expr", bson.D{
						{
							"$eq", [2]bson.D{
								bson.D{{"$year", ref_currentField}}, bson.D{{"$year", this.ref_field}},
							},
						},
					},
				},
				{
					"$expr", bson.D{
						{
							"$eq", [2]bson.D{
								bson.D{{"$month", ref_currentField}}, bson.D{{"$month", this.ref_field}},
							},
						},
					},
				},
				{
					"$expr", bson.D{
						{
							this.aspect, [2]bson.D{
								bson.D{{"$dayOfMonth", ref_currentField}}, bson.D{{"$dayOfMonth", this.ref_field}},
							},
						},
					},
				},
			},
		}
	case this.m&2 != 0:
		ret = bson.A{
			bson.D{
				{
					"$expr", bson.D{
						{
							"$eq", [2]bson.D{
								bson.D{{"$year", ref_currentField}}, bson.D{{"$year", this.ref_field}},
							},
						},
					},
				},
				{
					"$expr", bson.D{
						{
							this.aspect, [2]bson.D{
								bson.D{{"$month", ref_currentField}}, bson.D{{"$month", this.ref_field}},
							},
						},
					},
				},
			},
		}
	case this.m&1 != 0:
		ret = bson.A{
			bson.D{
				{
					"$expr", bson.D{
						{
							this.aspect, [2]bson.D{
								bson.D{{"$year", this.current_field}}, bson.D{{"$year", this.ref_field}},
							},
						},
					},
				},
			},
		}
	default:
		ret = make(primitive.A, 1)
	}

	return bson.Marshal(ret)
}
