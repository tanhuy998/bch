package requestPresenter

import (
	"app/model"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type AddCandidateRequest struct {
	CampaignUUID   *uuid.UUID       `param:"campaignUUID" validate:"required"`
	InputCandidate *model.Candidate `json:"data" validate:"required"`
}

/*
# IMPLEMENT IRequestBinder
*/
func (this *AddCandidateRequest) Bind(ctx iris.Context) error {

	err := ctx.ReadURL(this)

	if err != nil {

		return err
	}

	err = ctx.ReadJSON(this)

	if err != nil {

		return err
	}

	candidate := this.InputCandidate

	err = validateCandidateDateOfBirth(*candidate.DateOfBirth)

	if err != nil {

		return err
	}

	err = validateFormalName(*candidate.Name)

	if err != nil {

		return err
	}

	return nil
}

/*
# END IMPLEMENT IRequestBinder
*/
