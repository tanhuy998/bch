package requestPresenter

import (
	"app/domain/model"
	"app/internal/common"
	"time"

	"github.com/kataras/iris/v12"
)

const (
	CANDIDATE_SIGNING_OLD = 17
)

type CommitCandidateSigningInfoRequest struct {
	CandidateUUID string                      `param:"candidateUUID" validate:"required"`
	CampaignUUID  string                      `param:"campaignUUID" validate:"required"`
	Data          *model.CandidateSigningInfo `json:"data" validate:"required"`
}

/*
# IMPLEMENT IRequestBinder
*/
func (this *CommitCandidateSigningInfoRequest) Bind(ctx iris.Context) error {

	err := ctx.ReadURL(this)

	if err != nil {

		return err
	}

	err = ctx.ReadJSON(this)

	if err != nil {

		return err
	}

	currentYear := time.Now().Year()

	data := this.Data

	dateOfBirth := data.CivilIndentity.DateOfBirth

	if dateOfBirth == nil || currentYear-dateOfBirth.Year() < CANDIDATE_SIGNING_OLD {

		return common.ERR_INVALID_HTTP_INPUT
	}

	return nil
}

/*
# END IMPLEMENT IRequestBinder
*/
