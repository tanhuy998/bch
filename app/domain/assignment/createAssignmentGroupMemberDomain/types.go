package createAssignmentGroupMemberDomain

import (
	"context"
	"fmt"
)

var (
	err_already_assigned = fmt.Errorf("the given command group users already assigned")
)

type (
	query_result struct {
	}

	domain_context struct {
		context.Context
	}
)
