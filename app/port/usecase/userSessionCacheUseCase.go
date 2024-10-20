package usecasePort

import (
	"app/internal/bootstrap"
	"app/internal/cacheList"
	"app/internal/generalToken"
	"app/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	RefreshTokenBlackList = cacheList.CacheListManipulator[string, bootstrap.RefreshTokenBlackListCacheValue]
	GeneralTokenWhiteList = cacheList.CacheListManipulator[generalToken.GeneralTokenID, bootstrap.GeneralTokenWhiteListCacheValue]
)

type (
	UserSessionCacheUseCase struct {
		GeneralTokenWhiteList
		RefreshTokenBlackList
		UserSessionRepo repository.IUserSession
		MongoClient     *mongo.Client
	}
)
