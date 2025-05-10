package binaryRelationTriat

import (
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/relation/internal/resolver"
)

type (
	abstract_one_to_n struct {
	}
)

func (this abstract_one_to_n) GetBridgeResolver() relation.IDBRelationResolver {

	// return resolver.Unwind{}

	return resolver.One_To_Many{}
}
