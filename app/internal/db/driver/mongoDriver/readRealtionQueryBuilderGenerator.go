package mongoDriver

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder/relationQueryBuilder"
	"app/internal/db/relation"
)

type (
	ReadRelationQueryBuilderGenerator struct {
	}
)

func (r *ReadRelationQueryBuilderGenerator) NewRelationQuery() relation.IClonableReadRelationQueryBuilder {
	ret := new(relationQueryBuilder.MongoRelationQueryBuilder)

	ret.Init()

	return ret
}
