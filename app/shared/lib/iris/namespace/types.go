package namespace

import (
	"app/shared"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

type (
	IIrisNamespace interface {
		shared.INamespace
		iris.Party
		INamespaceContainer
	}
)

type (
	INamespaceContainer interface {
		GetContainer() *hero.Container
	}
)

type (
	INamespaceEvalutator interface {
		Namespace(path string, evaluators ...shared.INamespace) IIrisNamespace
	}
)
