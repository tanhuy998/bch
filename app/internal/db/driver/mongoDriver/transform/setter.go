package transform

import (
	"app/internal/db/query"
	"fmt"
)

type (
	setter struct {
		cur       string
		SetterMap map[string]interface{}
	}
)

// func (this *data_transformer)

func (this *setter) Set(field string) query.IDataTransformSetter {

	this.cur = field

	return this
}

func (this *setter) Value(val interface{}) {

	field := this.cur

	if field == "" {

		return
	}

	this.SetterMap[field] = val
}

func (this *setter) Ref(field string) {

	newField := this.cur

	if this.cur == "" {

		return
	}

	this.SetterMap[newField] = fmt.Sprintf("$%s", field)
}
