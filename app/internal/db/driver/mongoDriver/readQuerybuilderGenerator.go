package mongoDriver

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder"
	"app/internal/db/query"
)

type (
	ReadQueryBuilderGenerator struct {
	}
)

func (this *ReadQueryBuilderGenerator) NewQueryBuilder() query.IClonableQueryBuilder {

	return new(mongoQueryBuilder.MongoAggregateQueryBuilder)
}
