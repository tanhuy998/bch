package createAssignmentDomain

import (
	libCommon "app/src/internal/lib/common"
	"app/src/model"
	assignmentServicePort "app/src/port/assignment"
	"app/src/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	MIN_DEALINE time.Duration = time.Hour * 24 * 7
)

var (
	ERR_DUPLICATE_ASSIGNMENT = errors.New("createAssignmentService error: dublicate assignment")
	ERR_INVALID_DEADLINE     = fmt.Errorf(`assignment deadline must be at least %d days`, MIN_DEALINE)
)

type (
	CreateAssignmentService struct {
		GetSingleAssignmentService assignmentServicePort.IGetSingleAssignnment
		AssignmentRepo             repository.IAssignment
	}
)

func (this *CreateAssignmentService) Serve(data *model.Assignment, ctx context.Context) (*model.Assignment, error) {

	// existing, err := this.GetSingleAssignmentService.Search(
	// 	&model.Assignment{
	// 		TenantUUID: data.TenantUUID,

	// 	},
	// 	ctx,
	// )

	if !this.validateCreateAt(*data.CreatedAt) {

		return nil, errors.New("assignment issue date must be the same day")
	}

	if !this.validateDeadLine(*data.Deadline) {

		return nil, ERR_INVALID_DEADLINE
	}

	existSimilar, err := this.AssignmentRepo.Find(
		bson.D{
			{"tenantUUID", data.TenantUUID},
			{
				"createdAt", bson.D{
					{
						// check if there is no similar assignment in the same year
						"$gte", time.Date(data.Deadline.Year(), 1, 0, 0, 0, 0, 0, time.UTC),
					},
				},
			},
		},
		ctx,
	)

	if err != nil {

		return nil, err
	}

	if existSimilar != nil {

		return nil, ERR_DUPLICATE_ASSIGNMENT
	}

	data.UUID = libCommon.PointerPrimitive(uuid.New())

	err = this.AssignmentRepo.Create(data, context.TODO())

	if err != nil {

		return nil, err
	}

	return data, nil
}

func (this *CreateAssignmentService) validateCreateAt(c time.Time) bool {

	return time.Until(c) < time.Hour*24
}
func (this *CreateAssignmentService) validateDeadLine(d time.Time) bool {

	return time.Until(d) >= MIN_DEALINE
}
