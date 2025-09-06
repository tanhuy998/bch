package ioc

import (
	"app/internal/generalToken"
	irisIoc "app/internal/lib/iris/ioc"
	"app/main/internal/dependencies/memoryCaches"

	cacheListServicePort "app/port/cacheList"

	cacheListService "app/service/cacheList"

	"github.com/kataras/iris/v12/hero"
)

func RegisterCaches(container *hero.Container) {

	refreshTokenBlackListCacheClient, err := memoryCaches.NewRefreshTokenBlackListClient()

	if err != nil {

		panic("error while inittiating refresh token blacklist cache client: " + err.Error())
	}

	container.Register(refreshTokenBlackListCacheClient)

	generalTokenWhiteListCacheClient, err := memoryCaches.NewGeneralTokenWhiteListClient()

	if err != nil {

		panic("error while inittiating general token whitelist cache client: " + err.Error())
	}

	container.Register(generalTokenWhiteListCacheClient)

	irisIoc.BindDependency[
		cacheListServicePort.ICacheList[string, memoryCaches.RefreshTokenBlackListCacheValue],
		cacheListService.CacheListManipulator[string, memoryCaches.RefreshTokenBlackListCacheValue],
	](
		container,
		cacheListService.NewCacheListManipulator[string, memoryCaches.RefreshTokenBlackListCacheValue]("refresh_token_black_list"),
	)

	irisIoc.BindDependency[
		cacheListServicePort.ICacheList[generalToken.GeneralTokenID, memoryCaches.GeneralTokenWhiteListCacheValue],
		cacheListService.CacheListManipulator[generalToken.GeneralTokenID, memoryCaches.GeneralTokenWhiteListCacheValue],
	](
		container,
		cacheListService.NewCacheListManipulator[generalToken.GeneralTokenID, memoryCaches.GeneralTokenWhiteListCacheValue]("general_token_white_list"),
	)
}
