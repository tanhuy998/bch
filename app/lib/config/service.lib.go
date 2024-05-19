package libConfig

import (
	libCommon "app/lib/common"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func BindDependency[AbstractType any, ConcreteType any](
	container *hero.Container, concreteVal *ConcreteType,
) *hero.Dependency {

	abstractType := libCommon.Wrap[AbstractType]()

	checkInterfaceOrPanic(abstractType)

	var autowireField bool = false

	if concreteVal == nil {

		autowireField = true
		concreteVal = new(ConcreteType)
	}

	_checkImplementationOrPanic(abstractType, reflect.TypeOf(concreteVal))

	dep := container.Register(concreteVal)
	dep.DestType = abstractType
	dep.StructDependents = autowireField

	return dep
}

func BindAndMapDependencyToContext[AbstractType any, ConcreteType any](
	container *hero.Container, concreteVal *ConcreteType, contextKey string,
) *hero.Dependency {

	aType := libCommon.Wrap[AbstractType]()

	checkInterfaceOrPanic(aType)

	var autowireField bool = false

	if concreteVal == nil {

		autowireField = true
		concreteVal = new(ConcreteType)
	}

	_checkImplementationOrPanic(aType, reflect.TypeOf(concreteVal))

	mappedObj, _ := any(concreteVal).(AbstractType)

	dep := container.Register(func(ctx iris.Context) AbstractType {

		ctx.Values().Set(contextKey, concreteVal)

		return mappedObj
	})
	dep.StructDependents = autowireField

	return dep
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
