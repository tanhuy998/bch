package unary

import (
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/relation/internal/closure"
)

type (
	RelationByIntiator[Initiator_T relation.IDBRelationInitiator] struct {
		closure.ClosureGenerator[
			Initiator_T, Initiator_T,
		]
	}
)
