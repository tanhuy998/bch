package operator

import (
	"app/internal/db/driver/mongoDriver/expression"

	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type (
	IFilterExpression interface {
		//expression.IBinaryExpression[interface{}, interface{}]
		// AssertField(fieldName string, assertedType bsontype.Type)
		// Negate()
		CurrentFieldName() string
	}
)

type (
	IFilterExpressionPivot interface {
		IRootExpressionPivot
		IFilterExpression
		//Negate()
	}
)

type (
	IExpressionEvaluator interface {
		//IIterator
		AsBinaryExpression() expression.IBinaryExpression[interface{}, interface{}]
		AsUnaryExpression() expression.IUnaryExpression[interface{}]
	}
)

type (
	IRootExpressionPivot interface {
		AssertField(fieldName string, assertedType bsontype.Type)
	}
)

type (
	IPrimaryOperator interface {
		SetRoot(p IFilterExpressionPivot)
		Pivot() IFilterExpressionPivot
		Root() IFilterExpression
		AssertFieldIfSelfReference(
			fieldName interface{},
		) string
		OverrideExpression() IExpressionEvaluator
		SetAssertedType(t bsontype.Type)
		GetBasedExpression() expression.IBinaryExpression[interface{}, interface{}]
	}
)
