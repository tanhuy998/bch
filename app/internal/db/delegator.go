package db

import (
	"app/internal/db/query"
	"app/model"
	"reflect"
)

type (
	IDBDelegator[Model_T any] interface {
		query.IGenericQueryBuilder[Model_T]
		//Clone() IDBDelegator[Model_T]
	}
)

type (
	DBDelegator[Entity_T any] struct {
		default_projection map[string]int
	}
)

func test(q IDBDelegator[model.User]) {

	// q.Join(
	// 	"asdsa",
	// 	func(queryBuilder query.IJoinField) {

	// 		queryBuilder.On("localField", "foreignField").
	// 			As("alias").
	// 			Join(
	// 				"asdsa",
	// 				func(query query.IJoinField) {

	// 				},
	// 			).Join(
	// 			"asdssad",
	// 			func(query query.IJoinField) {

	// 			},
	// 		)
	// 	},
	// ).Filter(
	// 	func(filter query.IFilterExpression) {

	// 	},
	// ).Transform(
	// 	func(tran query.IDataTransformer) {
	// 		tran.Set("new").Value("sad")
	// 		tran.Set("new").Ref("name.last")
	// 	},
	// )

	// q.Join(
	// 	"asdasd",
	// 	func(queryBuilder query.IJoinField) {

	// 		queryBuilder.On("id", "id").
	// 		As("asd").
	// 		Join("asdasd", func(query query.IJoinField) {

	// 			query.On()
	// 		}).
	// 	}
	// )
	// // .Filter(
	// // 	func(filter query.IFilterGenerator) {

	// // 	},
	// // ).
}

func (this *DBDelegator[Entity_T]) Init() {

	this.initDefaultProjection()
}

func (this *DBDelegator[Entity_T]) initDefaultProjection() {

	if len(this.default_projection) > 0 {

		return

	} else {

		this.default_projection = make(map[string]int)
	}

	ref := reflect.TypeFor[Entity_T]()

	for i := 0; i < ref.NumField(); i++ {

		field := ref.Field(i)

		switch field.Type.Kind() {
		case reflect.Interface, reflect.Func:
			continue
		}

		fieldName := field.Name

		this.default_projection[fieldName] = 1
	}
}
