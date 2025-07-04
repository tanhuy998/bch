package mongoFunc

import (
	"app/internal/db/query"
	"fmt"
	"slices"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	__arguments__ struct {
		__parameters__
		/*
			A slice that store the arguments that is passed to the function.
			The arguments order is the parameters order.
		*/
		args []interface{}
		/*
			A map whose key and value correspond to parameter's name
			and argument index respectively.
		*/
		args_lookup_map map[string]int
	}
)

func (this *__arguments__) evaluateArgs() {

	if len(this.args) > 0 {

		return
	}

	this.args = make([]interface{}, 0)
}

func (this *__arguments__) SetArgs(args ...interface{}) {

	// this.args = args

	this.__setArgs(args)
}

func (this *__arguments__) __setArgs(args []interface{}) {

	copy(this.args, args)
}

func (this *__arguments__) SetArgument(paramName string, val interface{}) {

	this.evaluateArgs()

	if index, exists := this.args_lookup_map[paramName]; exists {

		this.args[index] = val
		return
	}

	this.AddArgument(val)
	this.args_lookup_map[paramName] = len(this.args) - 1
	this.__parameters__.AddParam(paramName)
}

func (this *__arguments__) DefineArgs(argsMap map[string]interface{}) {

	//this.__proto__.param_list = make([]string, len(argsMap))

	if len(argsMap) == 0 {

		return
	}

	// i := 0

	// lenArgs := len(this.args_lookup_map)
	// lenParams := len(this.__parameters__.param_list)

	// if lenArgs > lenParams {

	// }

	for paramName, val := range argsMap {

		// this.__parameters__.param_list[i] = paramName
		// this.args_lookup_map[paramName] = i

		this.SetArgument(paramName, val)
	}
}

func (this *__arguments__) AbandonArgument(paramName string) {

	this.SetArgument(paramName, bson.TypeUndefined)
}

func (this *__arguments__) SetArgumentByIndex(index int, val interface{}) {

	switch {
	case index+1 > len(this.args):
		this.__setArgsPadding(len(this.args)-1, index-1)
		fallthrough
	default:
		this.args[index] = val
	}
}

func (this *__arguments__) __setArgsPadding(offset int, n int) {

	switch {
	case offset <= 0 || n <= 0:
		return
	case offset == 0:
		additionSlice := make([]interface{}, n)
		this.args = slices.Concat(additionSlice, this.args)
		return
	case offset+1 > len(this.args):
		additionSlice := make([]interface{}, offset+1-len(this.args))
		this.args = slices.Concat(this.args, additionSlice)
		return
	default:
		ref := this.args

		this.args = slices.Concat(
			ref[:offset],
			make([]interface{}, n),
			ref[offset:],
		)
	}
}

func (this *__arguments__) AddArgument(val interface{}) {

	this.evaluateArgs()

	switch v := val.(type) {
	case query.IQuerySelfReference:
		this.args = append(this.args, fmt.Sprintf(`$%s`, v.GetQuerySelfReference()))
	default:
		this.args = append(this.args, val)
	}
}
