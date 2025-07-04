package mongoFunc

import (
	"go.mongodb.org/mongo-driver/bson"
)

type (
	MongoFunction struct {
		//__proto__
		// params       []string

		// body         string
		__prototype__
	}
)

func (this *MongoFunction) Resovle() {

}

// func (this *MongoFuncPrototype) resolveArgs() {

// 	this.params = make([]string, len(this.args))

// 	for i, arg := range this.args {

// 		this.params[i] = fmt.Sprintf(`param%d`, i)

// 		switch v := arg.(type) {
// 		case query.IQuerySelfReference:
// 			this.args[i] = fmt.Sprintf(`$%s`, v.GetQuerySelfReference())
// 		}
// 	}
// }

func (this *MongoFunction) init() {

}

func (this *MongoFunction) SetBody(body interface{}) {

	var body_str string

	switch v := body.(type) {
	case string:
		body_str = v
	case MongoFuncBody:
		body_str = v.String()
	default:
		panic("")
	}

	this.body = body_str
}

func (this MongoFunction) GetArgs() (args []interface{}) {

	copy(args, this.args)

	return
}

func (this MongoFunction) AsBson() bson.D {

	return bson.D{
		{"body", this.String()},
		{"args", this.GetArgs()},
		{"lang", "js"},
	}
}

func (this MongoFunction) ToBsonExpression() bson.D {

	return bson.D{
		{
			"$function", this.AsBson(),
		},
	}
}

func (this MongoFunction) MarshalBSON() ([]byte, error) {

	return bson.Marshal(this.AsBson())
}
