package bootstrap

import (
	"app/internal/generalToken"
	"app/internal/memoryCache"
	"errors"
	"fmt"
)

const (
	REFRESH_TOKE_BLACK_LIST_TOPIC = "refresh-token-black-list"
	GENERAL_TOKE_WHITE_LIST_TOPIC = "general-token-white-list"
)

type (
	GeneralTokenWhiteListCacheValue = struct{}
	RefreshTokenBlackListCacheValue = struct{}

	GeneralTokenWhiteListCacheClient = memoryCache.CacheClient[generalToken.GeneralTokenID, GeneralTokenWhiteListCacheValue]
	RefreshTokenBlackListCacheClient = memoryCache.CacheClient[string, RefreshTokenBlackListCacheValue]
)

func init() {

	err := memoryCache.NewTopic[generalToken.GeneralTokenID, GeneralTokenWhiteListCacheValue](GENERAL_TOKE_WHITE_LIST_TOPIC)

	if err != nil {

		panic(errors.Join(fmt.Errorf("error while initializing General token cache topic"), err))
	}

	err = memoryCache.NewTopic[string, RefreshTokenBlackListCacheValue](REFRESH_TOKE_BLACK_LIST_TOPIC)

	if err != nil {

		panic(errors.Join(fmt.Errorf("error while initializing Refresh token cache topic"), err))
	}

}

func NewGeneralTokenWhiteListClient() (*memoryCache.CacheClient[generalToken.GeneralTokenID, GeneralTokenWhiteListCacheValue], error) {

	return memoryCache.NewClient[generalToken.GeneralTokenID, GeneralTokenWhiteListCacheValue](
		GENERAL_TOKE_WHITE_LIST_TOPIC,
	)
}

func NewRefreshTokenBlackListClient() (*memoryCache.CacheClient[string, RefreshTokenBlackListCacheValue], error) {

	return memoryCache.NewClient[string, RefreshTokenBlackListCacheValue](
		REFRESH_TOKE_BLACK_LIST_TOPIC,
	)
}
