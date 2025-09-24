package api

import "github.com/kataras/iris/v12"

type (
	BranchInitializationFunc func(cur *APIBuilder)
)

type (
	APIBuilder struct {
		iris.Party
	}
)

func NewAPIBuilder(party iris.Party) *APIBuilder {

	return &APIBuilder{party}
}

func (this *APIBuilder) Branch(path string, fn BranchInitializationFunc) {

	child := this.Party.Party(path)

	fn(NewAPIBuilder(child))
}
