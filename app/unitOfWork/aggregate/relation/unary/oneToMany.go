package unary

import (
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/relation/internal/closure"
	"app/unitOfWork/aggregate/relation/internal/resolver"
)

type (
	OneToManyRelationBy[Initializer_T relation.IDBRelationInitializer] struct {
		closure.ClosureGenerator[
			resolver.One_To_Many, Initializer_T,
		]
	}
)
