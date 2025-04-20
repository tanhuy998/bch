package binaryRelationTriat

import (
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/relation/internal/pivot"
	"app/unitOfWork/aggregate/relation/internal/resolver"
)

type (
	OneToOneBy[
		Bridge_Entity_T any,
		Foreign_Entity_T any,
	] struct {
		abstract_one_to_n
		pivot.Bridge[Bridge_Entity_T]
		pivot.Foreign[Foreign_Entity_T]
	}
)

func (this OneToOneBy[Bridge_Entity_T, Foreign_Entity_T]) GetForeignResolver() relation.IDBRelationResolver {

	return resolver.One_To_One{}
}
