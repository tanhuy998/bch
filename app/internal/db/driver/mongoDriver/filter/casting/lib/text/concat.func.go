package text

import (
	"app/internal/db/driver/mongoDriver/lib/mongoFunc"
	"fmt"
)

const (
	MONGO_FUNC_REGEX_CONCAT mongoFunc.MongoDBJSFuncPrototype = `
		function () {

			return new RegExp(targetField + %s)
		}
	`
)

type (
	MongoFuncRegexConcat struct {
		mongoFunc.MongoFunction
		//concate_elements []interface{}
	}
)

func NewMongoFuncRegexConcat(concatElements []interface{}) mongoFunc.IBson {

	params := make([]string, len(concatElements))

	for i := range concatElements {

		params[i] = fmt.Sprintf(`param%d`, i)
	}

	return mongoFunc.NewMongoFunction().
		Parameters(params...).
		SetArguments(concatElements...).
		Body(nil).
		Return(`new RegExp(targetField + %s)`).
		Build()
}

// func (this *MongoFuncRegexConcat) init() {

// 	this.resolveConcateElements()
// 	this.resolveBody()
// }

// func (this *MongoFuncRegexConcat) resolveConcateElements() {

// }

// func (this *MongoFuncRegexConcat) resolveBody() {

// 	//args_str := strings.Join(this.GetArgs(), ",")

// 	this.Return(
// 		`new RegExp(targetField + %s)`,
// 	)
// }

// func (this *MongoFuncRegexConcat) SetTargetField(fieldName string) {

// 	this.SetArgs(fieldName)
// }
