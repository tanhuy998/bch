package namespace

// import (
// 	"app/shared"

// 	"github.com/kataras/iris/v12"
// 	"github.com/kataras/iris/v12/hero"
// )

// type (
// 	IrisNamespacePrototype struct {
// 		*iris.Application
// 		parent iris.Party
// 	}
// )

// func New(parent iris.Party) *iris.Application {

// 	ret := iris.New()

// 	if parent == nil {

// 		return ret
// 	}

// 	//container := ret.ConfigureContainer().Container

// 	ret.ConfigureContainer().Container = parent.ConfigureContainer().Container.Clone()

// 	return ret
// }

// func (this *IrisNamespacePrototype) Init() {

// 	if this.Application != nil {

// 		return
// 	}

// 	this.Application = iris.New()

// 	this._initializeContainer()
// }

// func (this *IrisNamespacePrototype) _initializeContainer() {

// 	if this.parent == nil {

// 		return
// 	}
// }

// func (this *IrisNamespacePrototype) RegisterDependencies(
// 	bindings ...DependenyBindingFunc,
// ) {

// 	for _, fn := range bindings {

// 		switch fn {
// 		case nil:
// 			continue
// 		default:
// 			fn(this.ConfigureContainer().Container)
// 		}
// 	}
// }

// func (this *IrisNamespacePrototype) GetContainer() *hero.Container {

// 	return this.Application.ConfigureContainer().Container
// }

// func (this *IrisNamespacePrototype) Namespace(path string, evaluators ...shared.INamespace) IIrisNamespace {

// }

// func (this *IrisNamespacePrototype) Next() {

// }
