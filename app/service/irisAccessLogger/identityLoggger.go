package irisAccessLoggerService

import (
	libCommon "app/internal/lib/common"
	libIris "app/internal/lib/iris"
	"app/valueObject/log"
	"time"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

const (
	ANNONYMOUS_MSG = "annonymous"
)

type (
	identity_err_log_line struct {
		Type   string `json:"type"`
		Detail string `json:"detail"`
	}
)

type (
	identity_log_line struct {
		TenantUUID  *uuid.UUID `json:"tenantUUID"`
		UserUUID    *uuid.UUID `json:"userUUID"`
		TenantAgent bool       `json:"tenantAgent"`
		Expired     string     `json:"expired,omitempty"`
		ExpireAt    *time.Time `json:"expireAt,omitempty"`
	}
)

func assignIdentity(logObj *log.HTTPLogLine, ctx iris.Context) {

	switch at := libIris.GetAccessToken(ctx); {
	case at == nil:
		logObj.SetIdentity(ANNONYMOUS_MSG)
	default:
		logObj.SetIdentity(
			identity_log_line{
				TenantUUID:  libCommon.PointerPrimitive(at.GetTenantUUID()),
				UserUUID:    libCommon.PointerPrimitive(at.GetUserUUID()),
				TenantAgent: at.GetAuthData().IsTenantAgent(),
				Expired: libCommon.Ternary(
					!at.HasExpire(),
					"no_expire",
					libCommon.Ternary(at.Expired(), "true", "false"),
				),
				ExpireAt: at.GetExpireTime(),
			},
		)
	}
}
