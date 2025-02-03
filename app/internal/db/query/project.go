package query

type (
	ISubQueryProjector interface {
		Select(fields ...string) ISubQueryBuilder
		ExcludeFields(fields ...string) ISubQueryBuilder
	}
)

type (
	// IRepositoryProjectableOperator[Model_T any] interface {
	// 	Find(ctx context.Context) ([]*Model_T, error)
	// 	FindOne(ctx context.Context) (*Model_T, error)
	// }

	IProjector[Model_T any] interface {
		//IDataRetrievalQuery[Model_T]
		Select(fields ...string) IQueryBuilder[Model_T]        //IDataRetrievalQuery[Model_T]        // IRepositoryProjectableOperator[Model_T]
		ExcludeFields(fields ...string) IQueryBuilder[Model_T] // IDataRetrievalQuery[Model_T] // IRepositoryProjectableOperator[Model_T]
	}
)
