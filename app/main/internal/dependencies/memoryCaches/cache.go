package memoryCaches

import (
	"app/internal/generalToken"
	"app/internal/memoryCache"
	"errors"
	"fmt"
	"os"
)

const (
	REFRESH_TOKE_BLACK_LIST_TOPIC = "refresh-token-black-list"
	GENERAL_TOKE_WHITE_LIST_TOPIC = "general-token-white-list"
	ENV_CACHE_LOG                 = "CACHE_LOG"
)

type (
	GeneralTokenWhiteListCacheValue = struct{}
	RefreshTokenBlackListCacheValue = struct{}

	GeneralTokenWhiteListCacheClient = memoryCache.CacheClient[generalToken.GeneralTokenID, GeneralTokenWhiteListCacheValue]
	RefreshTokenBlackListCacheClient = memoryCache.CacheClient[string, RefreshTokenBlackListCacheValue]
)

func init() {

	err := memoryCache.NewTopic[generalToken.GeneralTokenID, GeneralTokenWhiteListCacheValue](GENERAL_TOKE_WHITE_LIST_TOPIC)

	if err != nil && !errors.Is(err, memoryCache.ERR_TOPIC_EXIST) {

		panic(errors.Join(fmt.Errorf("error while initializing General token cache topic"), err))
	}

	err = memoryCache.NewTopic[string, RefreshTokenBlackListCacheValue](REFRESH_TOKE_BLACK_LIST_TOPIC)

	if err != nil && !errors.Is(err, memoryCache.ERR_TOPIC_EXIST) {

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

func CacheLog() bool {

	return os.Getenv(ENV_CACHE_LOG) == "true"
}
