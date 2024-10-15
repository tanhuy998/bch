package createAssignmentDomain

import (
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	"app/repository"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	MIN_DEADLINE_PERIOD_DAY               = 7
	MIN_DEALINE_DUR         time.Duration = time.Hour * 24 * MIN_DEADLINE_PERIOD_DAY
)

var (
	ERR_DUPLICATE_ASSIGNMENT = errors.New("createAssignmentService error: dublicate assignment")
	ERR_INVALID_DEADLINE     = fmt.Errorf(`assignment deadline must be at least %d days`, MIN_DEADLINE_PERIOD_DAY)
)

type (
	CreateAssignmentService struct {
		AssignmentRepo repository.IAssignment
	}
)

func (this *CreateAssignmentService) Serve(
	tenantUUID uuid.UUID, dataModel *model.Assignment, ctx context.Context,
) (*model.Assignment, error) {

	dataModel.CreatedAt = libCommon.PointerPrimitive(time.Now())

	switch {
	case !this.validateCreateAt(*dataModel.CreatedAt):
		return nil, fmt.Errorf("assignment issue date must be the same day")
	case !this.validateDeadLine(*dataModel.Deadline):
		return nil, ERR_INVALID_DEADLINE
	}

	findSimilarQuery := bson.D{
		{"tenantUUID", tenantUUID},
		{"title", dataModel.Title},
		{
			"createdAt", bson.D{
				{
					// check if there is no similar assignment in the same year
					"$gte", time.Date(dataModel.Deadline.Year(), 1, 0, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	switch existSimilar, err := this.AssignmentRepo.Find(findSimilarQuery, ctx); {
	case err != nil:
		return nil, err
	case existSimilar != nil:
		return nil, errors.Join(common.ERR_CONFLICT, fmt.Errorf("there are similar assignment in this year"))
	}

	dataModel.UUID = libCommon.PointerPrimitive(uuid.New())
	dataModel.CreatedAt = libCommon.PointerPrimitive(time.Now())
	dataModel.TenantUUID = &tenantUUID

	err := this.AssignmentRepo.Create(dataModel, context.TODO())

	if err != nil {

		return nil, errors.Join(common.ERR_INTERNAL, err)
	}

	return dataModel, nil
}

func (this *CreateAssignmentService) validateCreateAt(c time.Time) bool {

	return time.Until(c) < time.Hour*24
}
func (this *CreateAssignmentService) validateDeadLine(d time.Time) bool {
	fmt.Println(time.Until(d), MIN_DEALINE_DUR)
	return time.Until(d) >= MIN_DEALINE_DUR
}
