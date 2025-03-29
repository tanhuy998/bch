package relation

type (
	ForeignManipulatorFunc func(foreign IReadRelationQueryBuilder)

	IRelationForeignNavigator interface {
		IDBRelationForeignManipulator
		SetLimit(uint64)
		GetLocalField() string
		GetForeignField() string
		GetAliasName() string
	}
)
