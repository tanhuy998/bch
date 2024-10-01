package config

import (
	"app/src/internal/memoryCache"
	refreshTokenBlackListServicePort "app/src/port/refreshTokenBlackList"
)

func init() {

	memoryCache.NewTopic[string, refreshTokenBlackListServicePort.IRefreshTokenBlackListPayload](
		refreshTokenBlackListServicePort.REFRESH_TOKE_BLACK_LIST_TOPIC,
	)
}
