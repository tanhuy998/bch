package authService

import (
	"app/src/model"
	"app/src/repository"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ICheckCommandGroupUserRole interface {
		//Serve(groupUUID string, userUUID string, roleUUIDs []string) error
		Compare(groupUUID uuid.UUID, userUUID uuid.UUID, roleUUID []uuid.UUID) (unAssignedRoles []uuid.UUID, err error)
	}

	CheckCommandGroupUserRoleService struct {
		CommandGroupUserRepo repository.ICommandGroupUser
	}
)

// /*
// check if a command group user has list of role
// */
// func (this *CheckCommandGroupUserRoleService) Serve(
// 	groupUUID_str string,
// 	userUUID_str string,
// 	roleUUID_str_list []string,
// ) error {

// 	groupUUID, err := uuid.Parse(groupUUID_str)

// 	if err != nil {

// 		return err
// 	}

// 	userUUID, err := uuid.Parse(userUUID_str)

// 	if err != nil {

// 		return err
// 	}

// 	var conditions bson.D

// 	conditions, err = this._prepareRolesConditions(roleUUID_str_list)

// 	if err != nil {

// 		return err
// 	}

// 	res, err := repository.Aggregate[model.CommandGroupUserRole](
// 		this.CommandGroupUserRepo.GetCollection(),
// 		mongo.Pipeline{
// 			bson.D{
// 				{
// 					"$match", bson.D{
// 						{"groupUUID", groupUUID},
// 						{"userUUID", userUUID},
// 					},
// 				},
// 			},
// 			bson.D{
// 				{"$lookup",
// 					bson.D{
// 						{"from", "commandGroupUserRoles"},
// 						{"localField", "uuid"},
// 						{"foreignField", "commandGroupUserUUID"},
// 						{"as", "roles"},
// 					},
// 				},
// 			},
// 			bson.D{{"$unwind", bson.D{{"path", "$roles"}}}},
// 			bson.D{
// 				{"$set",
// 					bson.D{
// 						{"roleUUID", "$roles.roleUUID"},
// 						{"userUUID", "$role.userUUID"},
// 					},
// 				},
// 			},
// 			bson.D{
// 				{
// 					"$match", bson.D{
// 						{"$or", conditions},
// 					},
// 				},
// 			},
// 			bson.D{
// 				{"$project",
// 					bson.D{
// 						{"_id", 0},
// 						{"uuid", 0},
// 						{"roles", 0},
// 					},
// 				},
// 			},
// 		},
// 		context.TODO(),
// 	)

// 	if err != nil {

// 		return err
// 	}

// 	if len(res) != len(roleUUID_str_list) {

// 		return ERR_ROLE_EXIST
// 	}

// 	return nil
// }

func (this *CheckCommandGroupUserRoleService) _prepareRolesConditions(
	roleUUIDList []uuid.UUID,
) ([]bson.D, error) {

	if len(roleUUIDList) == 0 {

		return nil, errors.New("no input")
	}

	var ret []bson.D = make([]bson.D, len(roleUUIDList))

	for i, v := range roleUUIDList {

		ret[i] = bson.D{{"roleUUID", v}}
	}

	return ret, nil
}

func (this *CheckCommandGroupUserRoleService) Compare(
	groupUUID uuid.UUID,
	userUUID uuid.UUID,
	roleUUIDList []uuid.UUID,
) (unGrantedRoles []uuid.UUID, err error) {

	var conditions []bson.D

	conditions, err = this._prepareRolesConditions(roleUUIDList)

	if err != nil {

		return unGrantedRoles, err
	}

	res, err := repository.Aggregate[model.CommandGroupUserRole](
		this.CommandGroupUserRepo.GetCollection(),
		mongo.Pipeline{
			bson.D{
				{
					"$match", bson.D{
						{"commandGroupUUID", groupUUID},
						{"userUUID", userUUID},
					},
				},
			},
			bson.D{
				{"$lookup",
					bson.D{
						{"from", "commandGroupUserRoles"},
						{"localField", "uuid"},
						{"foreignField", "commandGroupUserUUID"},
						{"as", "roles"},
					},
				},
			},
			bson.D{{"$unwind", bson.D{{"path", "$roles"}}}},
			bson.D{
				{"$set",
					bson.D{
						{"roleUUID", "$roles.roleUUID"},
						{"userUUID", "$role.userUUID"},
					},
				},
			},
			bson.D{
				{
					"$match", bson.D{
						{"$or", conditions},
					},
				},
			},
			bson.D{
				{"$project",
					bson.D{
						{"_id", 0},
						{"uuid", 0},
						{"roles", 0},
					},
				},
			},
		},
		context.TODO(),
	)

	if err != nil {

		return nil, err
	}

	if len(res) >= len(roleUUIDList) {
		/*
			length of the fetched result always less than or equal to
			the input list
		*/
		return nil, nil
	}

	diff := this.differentiate(roleUUIDList, res)

	fmt.Println(diff)

	return diff, nil
}

/*
private method, just have been called after fetch command group roles of a user
to get the unassigned roles
*/
func (this *CheckCommandGroupUserRoleService) differentiate(
	input []uuid.UUID,
	fetchedRoles []*model.CommandGroupUserRole,
) (unAssignedRoles []uuid.UUID) {

	tempInput := make([]uuid.UUID, 0)
	tempInput = append(tempInput, input...)

	// var r_map map[uuid.UUID]*model.CommandGroupUserRole = make(map[uuid.UUID]*model.CommandGroupUserRole)

	// for _, v := range fetchedRoles {

	// 	r_map[v.RoleUUID] = v
	// }

	// for i, v := range ret {

	// 	if _, ok := r_map[v]; ok {

	// 		ret = slices.Delete(ret, i, i)
	// 	}
	// }

	var visitedInputDistinctValues map[uuid.UUID]int = make(map[uuid.UUID]int, 0)

	for _, v := range tempInput {

		if times, ok := visitedInputDistinctValues[v]; ok {

			visitedInputDistinctValues[v] = times + 1
			continue
		}

		visitedInputDistinctValues[v] = 1
	}

	//fmt.Println(visitedInputDistinctValues)

	for _, v := range fetchedRoles {

		roleUUID := v.RoleUUID

		times, visited := visitedInputDistinctValues[*roleUUID]

		if !visited {

			continue
		}

		timesLeft := times - 1

		if timesLeft > 1 {

			visitedInputDistinctValues[*roleUUID] = timesLeft

		} else {

			delete(visitedInputDistinctValues, *roleUUID)
		}
	}
	//fmt.Println(visitedInputDistinctValues)
	var ret []uuid.UUID = make([]uuid.UUID, 0)

	for roleUUID := range visitedInputDistinctValues {

		ret = append(ret, roleUUID)
	}

	return ret
}
