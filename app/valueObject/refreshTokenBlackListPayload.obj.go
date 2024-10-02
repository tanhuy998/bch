package valueObject

import "github.com/google/uuid"

type (
	RefreshTokenBlackListPayload struct {
		UserUUID *uuid.UUID
	}
)

func (this *RefreshTokenBlackListPayload) GetUserUUID() uuid.UUID {

	return *this.UserUUID
}
