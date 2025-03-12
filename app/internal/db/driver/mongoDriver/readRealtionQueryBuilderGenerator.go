package mongoDriver

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder"
	"app/internal/db/relation"
)

type (
	ReadRelationQueryBuilderGenerator struct {
	}
)

func (r *ReadRelationQueryBuilderGenerator) NewRelationQuery() relation.IClonableReadRelationQueryBuilder {
	return new(mongoQueryBuilder.MongoRelationQueryBuilder)
}
