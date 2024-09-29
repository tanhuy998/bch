package config

import (
	refreshTokenBlackListServicePort "app/adapter/refreshTokenBlackList"
	memoryCache "app/mermoryCache"
)

func init() {

	memoryCache.NewTopic[string, refreshTokenBlackListServicePort.IRefreshTokenBlackListPayload](
		refreshTokenBlackListServicePort.REFRESH_TOKE_BLACK_LIST_TOPIC,
	)
}
