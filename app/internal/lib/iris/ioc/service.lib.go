package irisIoc

import (
	libCommon "app/internal/lib/common"
	"app/internal/lib/iris/ioc/internal"
	iocOption "app/internal/lib/iris/ioc/option"
	"reflect"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

var (
	concrete_pool map[reflect.Type]interface{} = map[reflect.Type]interface{}{}
)

// func RegisterObject(container *hero.Container, obj any) {

// 	dep := container.Register(obj)

// 	dep.StructDependents = true
// 	dep.Explicitly()
// }

func RegisterDependency[ConcreteType any](
	container *hero.Container, concreateObj *ConcreteType, options ...iocOption.DependencyOptionFunc,
) {

	var autowired bool = false

	switch {
	case container == nil:
		panic("nil container passed to libConfig.BindAs()")
	// case len(abstractTypes) == 0:
	// 	panic("empty abstract type passed to libConfig.BindAs()")
	case concreateObj == nil:
		autowired = true
		concreateObj = resolve_concrete_instance[ConcreteType]()
	}

	// for _, abstract := range abstractTypes {

	// 	dep := _bindDependency(container, abstract, concreateObj)

	// 	for _, optionFn := range options {
	// 		optionFn(dep)
	// 	}
	// }

	dep := container.Register(concreateObj)
	dep.StructDependents = autowired
	dep.Explicitly()

	for _, fn := range options {

		fn(container, dep)
	}
}

func BindDependency[AbstractType, ConcreteType any](
	container *hero.Container, concreteVal *ConcreteType,
) *hero.Dependency {

	abstractType := libCommon.Wrap[AbstractType]()

	return _bindDependency(container, abstractType, concreteVal)
}

func _bindDependency[ConcreteType any](container *hero.Container, abstractType reflect.Type, concreteVal *ConcreteType) *hero.Dependency {

	internal.CheckInterfaceOrPanic(abstractType)

	var autowireField bool = false

	if concreteVal == nil {

		autowireField = true
		concreteVal = resolve_concrete_instance[ConcreteType]()
	}

	internal.CheckTypeImplementationOrPanic(abstractType, reflect.TypeOf(concreteVal))

	// dep := container.Register(concreteVal)
	// dep.DestType = abstractType

	dep := internal.Register(container, abstractType, concreteVal)

	dep.StructDependents = autowireField
	//dep.Explicitly()

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

	internal.CheckInterfaceOrPanic(aType)

	var autowireField bool = false

	if concreteVal == nil {

		autowireField = true
		concreteVal = new(ConcreteType)
	}

	internal.CheckTypeImplementationOrPanic(aType, reflect.TypeOf(concreteVal))

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

// func checkImplementationOrPanic[AbstractType any, ConcreteType any]() {

// 	_checkImplementationOrPanic(libCommon.Wrap[AbstractType](), libCommon.Wrap[ConcreteType]())
// }

// func _checkImplementationOrPanic(abstract reflect.Type, concrete reflect.Type) {

// 	if concrete.Implements(abstract) {

// 		return
// 	}

// 	panic(
// 		fmt.Sprintf(
// 			"Could not bind concrete type %s as interface %s",
// 			//reflect.TypeOf(concreteVal).String(),
// 			concrete.String(),
// 			abstract.String(),
// 		),
// 	)
// }
