package internal

import (
	libCommon "app/internal/lib/common"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12/hero"
)

func CheckImplementationOrPanic[AbstractType any, ConcreteType any]() {

	CheckTypeImplementationOrPanic(libCommon.Wrap[AbstractType](), libCommon.Wrap[ConcreteType]())
}

func CheckTypeImplementationOrPanic(abstract reflect.Type, concrete reflect.Type) {

	if concrete.Implements(abstract) {

		return
	}

	panic(
		fmt.Sprintf(
			"Could not bind concrete type %s as interface %s",
			concrete.String(),
			abstract.String(),
		),
	)
}

func Register(container *hero.Container, abstractType reflect.Type, concreteObj interface{}) *hero.Dependency {

	CheckTypeImplementationOrPanic(abstractType, reflect.TypeOf(concreteObj))

	d := container.Register(concreteObj)

	d.DestType = abstractType
	d.Explicitly()

	return d
}

func CheckInterfaceOrPanic(t reflect.Type) {

	if t.Kind() != reflect.Interface {

		panic(
			fmt.Sprintf(
				"Could not use %s as abstract type which is not an interface",
				t.String(),
			),
		)
	}
}
