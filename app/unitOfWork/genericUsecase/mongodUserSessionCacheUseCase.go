package genericUseCase

import (
	"app/internal/bootstrap"
	"app/internal/generalToken"
	cacheListServicePort "app/port/cacheList"
	generalTokenServicePort "app/port/generalToken"
	"app/repository"
	opLog "app/unitOfWork/operationLog"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	RefreshTokenBlackList = cacheListServicePort.ICacheList[string, bootstrap.RefreshTokenBlackListCacheValue]
	GeneralTokenWhiteList = cacheListServicePort.ICacheList[generalToken.GeneralTokenID, bootstrap.GeneralTokenWhiteListCacheValue]

	AuthToken interface {
		GetUserUUID() uuid.UUID
		GetExpireTime() *time.Time
		GetTokenID() generalToken.GeneralTokenID
	}
)

type (
	UserSessionCacheUseCase struct {
		GeneralTokenWhiteList
		RefreshTokenBlackList
	}

	MongoUserSessionCacheUseCase[Output_T any] struct {
		UserSessionCacheUseCase
		opLog.OperationLogger
		UserSessionRepo repository.IUserSession
		MongoClient     *mongo.Client
	}
)

func (this *MongoUserSessionCacheUseCase[Output_T]) ModifyUserSession(
	ctx context.Context, fn func(context.Context) (*Output_T, error),
) (*Output_T, error) {

	mongoSession, err := this.MongoClient.StartSession()

	if err != nil {

		return nil, err
	}

	defer mongoSession.EndSession(ctx)

	ret, err := mongoSession.WithTransaction(
		ctx,
		func(ctx mongo.SessionContext) (interface{}, error) {

			return fn(ctx)
		},
	)

	if err != nil {

		return nil, err
	}

	if output, ok := ret.(*Output_T); ok {

		return output, nil
	}

	return nil, nil
}

func (this *MongoUserSessionCacheUseCase[Output_T]) RemoveUserSession(
	ctx context.Context, userUUID uuid.UUID, //beforeRemoveFns ...func() error,
) (err error) {

	userSessions, err := this.UserSessionRepo.FindMany(
		bson.D{
			{"userUUID", userUUID},
		},
		ctx,
	)

	if err != nil {

		return err
	}

	this.AccessLogger.PushTraceLogs(ctx, userSessions)

	defer func() {

		if err != nil {

			return
		}

		for _, v := range userSessions {
			// Delete caches of current user sessions
			// ctx of this funciton is a transaction context, therefore fetched data from
			// db have not committed until the whole transaction committed
			fmt.Println("existing session", *v.SessionID)
			_, err = this.GeneralTokenWhiteList.Delete(*v.SessionID, ctx)

			if err != nil {

				return
			}
		}
	}()

	// for _, fn := range beforeRemoveFns {

	// 	err = fn()

	// 	if err != nil {

	// 		return
	// 	}
	// }

	return nil
}

func (this *MongoUserSessionCacheUseCase[Output_T]) HasUserSession(
	generalToken generalTokenServicePort.IGeneralToken, ctx context.Context,
) (bool, error) {

	switch generalTokenInWhiteList, err := this.GeneralTokenWhiteList.Has(generalToken.GetTokenID(), ctx); {
	case err != nil:
		return false, err
	case generalTokenInWhiteList:
		return true, nil
	}

	switch existingUserSession, err := this.UserSessionRepo.Find(
		bson.D{
			{"sessionID", generalToken.GetTokenID()},
		},
		ctx,
	); {
	case err != nil:
		return false, err
	case existingUserSession != nil:
		return true, nil
	default:
		return false, nil
	}
}
