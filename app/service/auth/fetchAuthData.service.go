package authService

import (
	"app/repository"
	"errors"
)

var (
	ERR_USER_NOT_FOUND = errors.New("FetchAuthDataService: user not found")
)

type (
	//IFetchAuthdata = authServiceAdapter.IFetchAuthData

	FetchAuthDataService struct {
		UserRepo repository.IUser
	}
)

// func (this *FetchAuthDataService) Serve(userUUID uuid.UUID, ctx context.Context) (*valueObject.AuthData, error) {

// 	ret, err := repository.Aggregate[valueObject.AuthData](
// 		this.UserRepo.GetCollection(),
// 		mongo.Pipeline{
// 			bson.D{
// 				{
// 					"$match", bson.D{
// 						{"uuid", userUUID},
// 					},
// 				},
// 			},
// 			bson.D{
// 				{"$lookup",
// 					bson.D{
// 						{"from", "commandGroupUsers"},
// 						{"localField", "uuid"},
// 						{"foreignField", "userUUID"},
// 						{"as", "participatedCommandGroups"},
// 					},
// 				},
// 			},
// 			bson.D{
// 				{"$lookup",
// 					bson.D{
// 						{"from", "tenantAgents"},
// 						{"localField", "uuid"},
// 						{"foreignField", "userUUID"},
// 						{"as", "tenantAgentData"},
// 					},
// 				},
// 			},
// 			bson.D{
// 				{"$project",
// 					bson.D{
// 						{"uuid", 1},
// 						{"participatedCommandGroups", 1},
// 						{"tenantAgentData", 1},
// 					},
// 				},
// 			},
// 		},
// 		ctx,
// 	)

// 	if err != nil {

// 		return nil, err
// 	}

// 	if len(ret) == 0 {

// 		return nil, ERR_USER_NOT_FOUND
// 	}

// 	return ret[0], nil
// }
