package transform

import (
	"app/internal/db/relation"
)

type (
	RelationDataTransformer struct {
		QueryBuilderDataTransformer
		relation_navigator relation.IRelationForeignNavigator
	}
)

func NewRelationDataTransformer(
	navigator relation.IRelationForeignNavigator,
) *RelationDataTransformer {

	return &RelationDataTransformer{
		relation_navigator: navigator,
	}
}

func (this *RelationDataTransformer) Set(field string) relation.IRelationalDataTransformerSetterExpression {

	this.QueryBuilderDataTransformer.Set(field)

	return this
}

func (this *RelationDataTransformer) AsForeign() relation.IRelationForeignField {

	// use value instead of reference to inhance garabage collection
	return &foreign_field_expression{
		transformer: this,
	}
}
