package libService

import (
	libCommon "app/app/lib/common"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func BindDependency[AbstractType any, ConcreteType any](container *hero.Container, concreteVal *ConcreteType) {

	abstractType := libCommon.Wrap[AbstractType]()

	checkInterfaceOrPanic(abstractType)

	if concreteVal == nil {

		concreteVal = new(ConcreteType)
	}

	_checkImplementationOrPanic(abstractType, reflect.TypeOf(concreteVal))

	dep := container.Register(concreteVal)
	dep.DestType = abstractType
}

func BindAndMapDependencyToContext[AbstractType any, ConcreteType any](container *hero.Container, concreteVal *ConcreteType, contextKey string) {

	aType := libCommon.Wrap[AbstractType]()

	checkInterfaceOrPanic(aType)

	if concreteVal == nil {

		concreteVal = new(ConcreteType)
	}

	_checkImplementationOrPanic(aType, reflect.TypeOf(concreteVal))

	mappedObj, _ := any(concreteVal).(AbstractType)

	container.Register(func(ctx iris.Context) AbstractType {

		ctx.Values().Set(contextKey, concreteVal)

		return mappedObj
	})
}

func checkInterfaceOrPanic(t reflect.Type) {

	if t.Kind() != reflect.Interface {

		panic(
			fmt.Sprintf(
				"Could not use %s as abstract type which is not an interface",
				t.String(),
			),
		)
	}
}

func checkImplementationOrPanic[AbstractType any, ConcreteType any]() {

	_checkImplementationOrPanic(libCommon.Wrap[AbstractType](), libCommon.Wrap[ConcreteType]())
}

func _checkImplementationOrPanic(abstract reflect.Type, concrete reflect.Type) {

	if concrete.Implements(abstract) {

		return
	}

	panic(
		fmt.Sprintf(
			"Could not bind concrete type %s as interface %s",
			//reflect.TypeOf(concreteVal).String(),
			concrete.String(),
			abstract.String(),
		),
	)
}
