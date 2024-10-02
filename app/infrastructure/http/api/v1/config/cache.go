package config

import (
	"app/internal/memoryCache"
	refreshTokenBlackListServicePort "app/port/refreshTokenBlackList"
)

func init() {

	memoryCache.NewTopic[string, refreshTokenBlackListServicePort.IRefreshTokenBlackListPayload](
		refreshTokenBlackListServicePort.REFRESH_TOKE_BLACK_LIST_TOPIC,
	)
}
