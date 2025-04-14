package relation

type (
	ForeignManipulatorFunc func(foreign IReadRelationQueryBuilder)

	IRelationForeignUnwindable interface {
		UnwindLocal()
	}

	IRelationForeignNavigator interface {
		IDBRelationForeignManipulator
		IRelationForeignUnwindable
		SetLimit(uint64)
		GetLocalField() string
		GetForeignField() string
		GetAliasName() string
	}
)

type (
	IRelationForeignBeforeInitialization interface {
		BeforeInitialization()
	}

	IRelationForeignAfterInitialization interface {
		AfterInitialization()
	}
)
