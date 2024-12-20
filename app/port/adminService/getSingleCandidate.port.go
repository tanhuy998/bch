package adminServicePort

import "app/model"

type (
	IGetSingleCandidate interface {
		Serve(uuid string) (*model.Candidate, error)
	}
)
