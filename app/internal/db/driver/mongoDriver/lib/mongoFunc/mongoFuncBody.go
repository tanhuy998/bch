package mongoFunc

import "strings"

type (
	MongoFuncStatement = string
)

type (
	MongoFuncBody []MongoFuncStatement
)

func (this MongoFuncBody) String() string {

	return strings.Join(this, ";")
}
