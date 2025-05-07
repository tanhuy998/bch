package relations

import "app/unitOfWork/aggregate"

type (
	DomainContext interface {
		aggregate.IDomainContext
		aggregate.IAssignmentGroupDomain
	}
)
