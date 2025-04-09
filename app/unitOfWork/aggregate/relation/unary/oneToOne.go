package unary

import (
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/relation/internal/closure"
	"app/unitOfWork/aggregate/relation/internal/resolver"
)

type (
	OneToOneRelationBy[Initializer_T relation.IDBRelationInitializer] struct {
		closure.ClosureGenerator[
			resolver.One_To_One, Initializer_T,
		]
	}
)
