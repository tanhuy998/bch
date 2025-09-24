package repositoryAPI

import "context"

type (
	IStatisticAffectedCountableUnit interface {
		Count(ctx context.Context) (affectedDocsCount int64, err error)
	}
)
