package repositoryAPI

type (
	IJoinFilterMethod interface {
		Filter(filter FilterFunc) IJoinInitializer
	}

	IJoinField interface {
		On(lhsField string, rhsField string)
	}

	IJoinInitializer interface {
		Between(lhs string, rhs string) IJoinField
	}

	ReposioryJoinFunc = func(initializer IJoinInitializer)
)
