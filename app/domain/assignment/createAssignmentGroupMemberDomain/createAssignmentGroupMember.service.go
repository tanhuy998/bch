package createAssignmentGroupMemberDomain

import (
	"app/domain"
	"app/internal/common"
	libCommon "app/internal/lib/common"
	"app/model"
	authServicePort "app/port/auth"
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
		GetUnAssignedCommandGroupUsersService authServicePort.IGetAssignmentGroupUnAssignedCommandGroupUsers
		AssignmentGroupRepo                   repository.IAssignmentGroup
		AssignmentGroupMemberRepo             repository.IAssignmentGroupMember
		CommandGroupUserRepo                  repository.ICommandGroupUser
	}
)

func (this *CreateAssignmentGroupMemberService) Serve(
	tenantUUID uuid.UUID, assignmentGroupUUID uuid.UUID, data []*model.AssignmentGroupMember, ctx context.Context,
) ([]*model.AssignmentGroupMember, error) {

	if len(data) == 0 {

		return nil, errors.Join(
			common.ERR_BAD_REQUEST, fmt.Errorf("(CreateAssignmentGroupMemberService error) empty command group user list given"),
		)
	}

	if !this.InDomainContext(ctx) {

		switch {
		case tenantUUID == uuid.Nil:
			return nil, common.ERR_UNAUTHORIZED
		case assignmentGroupUUID == uuid.Nil:
			return nil, common.ERR_UNAUTHORIZED
		}

	}

	existingAssignmentGroup, err := this.AssignmentGroupRepo.FindOneByUUID(assignmentGroupUUID, ctx)

	switch {
	case err != nil:
		return nil, err
	case existingAssignmentGroup == nil:
		return nil, errors.Join(common.ERR_NOT_FOUND, fmt.Errorf("assignment group not found"))
	case *existingAssignmentGroup.TenantUUID != tenantUUID:
		return nil, errors.Join(common.ERR_FORBIDEN, fmt.Errorf("assignment group not in tenant"))
	}

	if !this.InDomainContext(ctx) {

		if err := this.validateCommandGroupUsers(tenantUUID, data, ctx); err != nil {

			return nil, err
		}
	}

	data, err = this.lookupUnAssigned(
		tenantUUID, *existingAssignmentGroup.UUID, data, ctx,
	)

	if err != nil {

		return nil, err
	}

	if len(data) == 0 {

		return nil, nil
	}

	err = this.AssignmentGroupMemberRepo.CreateMany(data, ctx)

	return data, err
}

func (this *CreateAssignmentGroupMemberService) validateCommandGroupUsers(
	tenantUUID uuid.UUID, uList []*model.AssignmentGroupMember, ctx context.Context,
) error {

	criterias := make(bson.A, len(uList))

	for i, v := range uList {

		criterias[i] = v.CommandGroupUserUUID
	}

	res, err := this.CommandGroupUserRepo.FindMany(
		bson.D{
			{
				"uuid", bson.D{
					{"$in", criterias},
				},
			},
			{"tenantUUID", tenantUUID},
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

func (this *CreateAssignmentGroupMemberService) lookupUnAssigned(
	tenantUUID, AssignmentGroupUUID uuid.UUID, data []*model.AssignmentGroupMember, ctx context.Context,
) ([]*model.AssignmentGroupMember, error) {

	commandGroupUserUUIDList := make([]uuid.UUID, len(data))

	for i, v := range data {

		commandGroupUserUUIDList[i] = *v.CommandGroupUserUUID
	}

	unAssignedList, err := this.GetUnAssignedCommandGroupUsersService.LookupUnAssigned(
		commandGroupUserUUIDList, tenantUUID, AssignmentGroupUUID, ctx,
	)

	if err != nil {

		return nil, err
	}

	m := make(map[uuid.UUID]*model.AssignmentGroupMember)

	for _, v := range data {

		m[*v.CommandGroupUserUUID] = v
	}

	ret := make([]*model.AssignmentGroupMember, len(unAssignedList))

	for i, v := range unAssignedList {

		if v, ok := m[*v.UUID]; ok {

			v.UUID = libCommon.PointerPrimitive(uuid.New())
			v.TenantUUID = &tenantUUID
			ret[i] = v
		}
	}

	return ret, nil
}
