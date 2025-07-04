package mongoFunc

import "go.mongodb.org/mongo-driver/bson"

type (
	IFunctionBuild interface {
		Build() IBson
	}

	IFunctionParameter interface {
		Parameters(names ...string) IFunctionPreBody
	}

	IFunctionPreBody interface {
		SetArguments(args ...interface{}) IFunctionBody
	}

	IFunctionBody interface {
		Body(interface{}) IFunctionReturn
	}

	IFunctionReturn interface {
		IFunctionBuild
		Return(string) IFunctionBuild
	}

	IFuncitonBuilder interface {
		IFunctionParameter
		IFunctionPreBody
		IFunctionBody
		IFunctionReturn
	}

	IBson interface {
		bson.Marshaler
	}
)

type (
	mongo_func_builder struct {
		MongoFunction
	}
)

func (this *mongo_func_builder) Parameters(names ...string) IFunctionPreBody {

	this.MongoFunction.__parameters__.param_list = names

	return this
}

func (this *mongo_func_builder) SetArguments(args ...interface{}) IFunctionBody {

	this.MongoFunction.__setArgs(args)

	return this
}

func (this *mongo_func_builder) Body(body interface{}) IFunctionReturn {

	this.SetBody(body)

	return this
}

func (this *mongo_func_builder) Return(statement string) IFunctionBuild {

	this.MongoFunction.Return(statement)

	return this
}

func (this *mongo_func_builder) Build() IBson {

	return &this.MongoFunction
}

func NewMongoFunction() IFunctionParameter {

	return new(mongo_func_builder)
}
