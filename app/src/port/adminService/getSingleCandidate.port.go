package adminServicePort

import "app/src/model"

type (
	IGetSingleCandidate interface {
		Serve(uuid string) (*model.Candidate, error)
	}
)
