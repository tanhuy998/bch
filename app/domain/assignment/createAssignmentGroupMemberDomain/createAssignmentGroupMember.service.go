package createAssignmentGroupMemberDomain

import (
	"app/domain"
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	"app/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	CreateAssignmentGroupMemberService struct {
		domain.ContextualDomainService[domain_context]
		AssignmentGroupRepo       repository.AssignmentGroupRepository
		AssignmentGroupMemberRepo repository.IAssignmentGroupMember
		CommandGroupUserRepo      repository.ICommandGroupUser
	}
)

func (this *CreateAssignmentGroupMemberService) Serve(
	tenantUUID uuid.UUID, assignmentGroupUUID uuid.UUID, commandGroupUserUUIDList []*model.AssignmentGroupMember, ctx context.Context,
) error {

	if !this.InDomainContext(ctx) {

		switch {
		case tenantUUID == uuid.Nil:
			return common.ERR_UNAUTHORIZED
		case assignmentGroupUUID == uuid.Nil:
			return common.ERR_UNAUTHORIZED
		case len(commandGroupUserUUIDList) == 0:
			return errors.Join(
				common.ERR_BAD_REQUEST, fmt.Errorf("(CreateAssignmentGroupMemberService error) empty command group user list given"),
			)
		}

		switch existingAssignmentGroup, err := this.AssignmentGroupRepo.FindOneByUUID(assignmentGroupUUID, ctx); {
		case err != nil:
			return err
		case existingAssignmentGroup == nil:
			return errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("command group not found"))
		case *existingAssignmentGroup.TenantUUID != tenantUUID:
			return errors.Join(common.ERR_FORBIDEN, fmt.Errorf("command group not in tenant"))
		}

		if err := this.validateCommandGroupUsers(tenantUUID, commandGroupUserUUIDList, ctx); err != nil {

			return err
		}
	}

	for _, v := range commandGroupUserUUIDList {

		v.TenantUUID = &tenantUUID
		v.UUID = libCommon.PointerPrimitive(uuid.New())
	}

	err := this.AssignmentGroupMemberRepo.CreateMany(commandGroupUserUUIDList, ctx)

	return err
}

func (this *CreateAssignmentGroupMemberService) validateCommandGroupUsers(
	tenantUUID uuid.UUID, uList []*model.AssignmentGroupMember, ctx context.Context,
) error {

	criterias := make(bson.A, len(uList))

	for i, v := range uList {

		criterias[i] = bson.D{
			{"uuid", v.CommandGroupUserUUID},
			{"tenantUUID", tenantUUID},
		}
	}

	res, err := this.CommandGroupUserRepo.FindMany(
		bson.D{
			{"$or", criterias},
		},
		ctx,
	)

	if err != nil {

		return err
	}

	if len(res) != len(uList) {

		return errors.Join(
			common.ERR_FORBIDEN, fmt.Errorf("(tenant agent error) any of requested users is not in the tenant"),
		)
	}

	return nil
}
