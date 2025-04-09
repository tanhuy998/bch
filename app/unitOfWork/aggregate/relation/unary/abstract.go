package unary

import (
	"app/unitOfWork/aggregate/relation/internal/api"
	"app/unitOfWork/aggregate/relation/internal/closure"
)

type (
	UnaryRelation[
		Initializer_T api.IInitializer,
		Resolver_T api.IResolver,
	] struct {
		Initiator closure.ClosureGenerator[
			Resolver_T, Initializer_T,
		]
	}
)
