package checkAssignmentParticipantDomain

import (
	"app/internal/common"
	"app/internal/db/query"
	assignmentServicePort "app/port/assignment"
	"errors"
	"fmt"
)

type (
	CheckAssignmentPariticipantService struct {
		CheckAssignmentParticipantAggregate
	}
)

func (this *CheckAssignmentPariticipantService) Serve(ctx assignmentServicePort.ICheckAssignmentParticipantContext) error {

	ret, err := this.CheckAssignmentParticipantAggregate.MergeRelations(
		ctx,
	).Read(
		func(queryBuilder query.IQueryBuilder) {

			queryBuilder.Filter(
				func(filter query.IFilterExpression) {

					filter.Field("$assignmentGroupMember").Not().Equal([]interface{}{})
				},
			)
		},
	).First(ctx)

	if err != nil {

		return err
	}

	if ret == nil {
		return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("the current user cannot access the resource"))
	}

	return nil
}
