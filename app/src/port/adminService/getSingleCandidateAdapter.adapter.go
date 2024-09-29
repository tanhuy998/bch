package adminServiceAdapter

import (
	"app/domain/model"
)

type (
	IGetSingleCandidate interface {
		Serve(uuid string) (*model.Candidate, error)
	}
)
