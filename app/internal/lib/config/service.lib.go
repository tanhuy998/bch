package libConfig

import (
	libCommon "app/internal/lib/common"
	"fmt"
	"reflect"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

var (
	concrete_pool map[reflect.Type]interface{} = map[reflect.Type]interface{}{}
)

func RegisterObject(container *hero.Container, obj any) {

	dep := container.Register(obj)

	dep.StructDependents = true
	dep.Explicitly()
}

func BindDependency[AbstractType, ConcreteType any](
	container *hero.Container, concreteVal *ConcreteType,
) *hero.Dependency {

	abstractType := libCommon.Wrap[AbstractType]()

	checkInterfaceOrPanic(abstractType)

	var autowireField bool = false

	if concreteVal == nil {

		autowireField = true
		concreteVal = resolve_concrete_instance[ConcreteType]()
	}

	_checkImplementationOrPanic(abstractType, reflect.TypeOf(concreteVal))

	dep := container.Register(concreteVal)
	dep.DestType = abstractType
	dep.StructDependents = autowireField
	dep.Explicitly()

	return dep
}

func OverrideDependency[AbstractType, ConcreteType any](
	container *hero.Container, concreateVal *ConcreteType,
) {

	t := libCommon.Wrap[AbstractType]()
	targetIndex := -1

	for i, dep := range container.Dependencies {

		if dep.DestType == t {

			targetIndex = i
			break
		}
	}

	if targetIndex >= 0 {

		list := container.Dependencies
		container.Dependencies = make([]*hero.Dependency, len(list)-1)
		copy(container.Dependencies, append(list[:targetIndex], list[targetIndex+1:]...))
	}

	BindDependency[AbstractType](container, concreateVal)
}

func resolve_concrete_instance[Concrete_T any]() *Concrete_T {

	wrapType := libCommon.Wrap[Concrete_T]()

	v, ok := concrete_pool[wrapType]

	if !ok {

		ret := new(Concrete_T)

		concrete_pool[wrapType] = ret

		// dep := container.Register(ret)
		// dep.StructDependents = true
		// dep.Explicitly()

		return ret
	}

	if ret, ok := v.(*Concrete_T); ok {

		return ret
	}

	panic("there are issue when resolving concrete instance in service container concrete bool")
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

	container.Handler(func(dep AbstractType) AbstractType {

		return dep
	})

	dep.StructDependents = autowireField
	dep.DestType = aType
	//dep.Explicitly()

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
