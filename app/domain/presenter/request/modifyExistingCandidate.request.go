package requestPresenter

import (
	"app/domain/model"
)

type UpdatedCandidate struct {
	Name     *string `json:"name,omitempty" bson:"name,omitempty" validate:"required_without_all,alphaunicodw"`
	IDNumber *string `json:"idNumber,omitempty" bson:"idNumber,omitempty" validate:"required_without_all,number,len=12"`
	Address  *string `json:"address,omitempty" bson:"address,omitempty" validate:"required_without_all,alphanumunicode"`
	Phone    *string `json:"phone,omitempty" bson:"phone,omitempty" validate:"number"`
}

type ModifyExistingCandidateRequest struct {
	UUID string           `paramm:"uuid" validate:"required,uuid_rfc4122"`
	Data *model.Candidate `json:"data" validate:"required"`
}
