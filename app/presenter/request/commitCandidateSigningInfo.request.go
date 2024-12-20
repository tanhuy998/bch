package requestPresenter

import (
	"app/internal/common"
	"app/model"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type CommitCandidateSigningInfoRequest struct {
	CandidateUUID *uuid.UUID                  `param:"candidateUUID" validate:"required"`
	CampaignUUID  *uuid.UUID                  `param:"campaignUUID" validate:"required"`
	Data          *model.CandidateSigningInfo `json:"data" validate:"required"`
}

/*
# IMPLEMENT IRequestBinder
*/
func (this *CommitCandidateSigningInfoRequest) Bind(ctx iris.Context) (err error) {

	defer func() {

		if err == nil {

			return
		}

		err = errors.Join(common.ERR_BAD_REQUEST, err)
	}()

	err = ctx.ReadURL(this)

	if err != nil {

		return err
	}

	err = ctx.ReadJSON(this)

	if err != nil {

		return err
	}

	err = this.validateYears()

	if err != nil {

		return err
	}

	return this.validateNames()
}

/*
# END IMPLEMENT IRequestBinder
*/

func (this *CommitCandidateSigningInfoRequest) validateYears() error {

	currentYear := time.Now().Year()

	data := this.Data

	if data == nil {

		return errors.New("no data")
	}

	dateOfBirth := data.CivilIndentity.DateOfBirth

	if currentYear-dateOfBirth.Year() < CANDIDATE_SIGNING_OLD {

		return errors.New("invalid year")
	}

	family := data.Family

	if dateOfBirth.Year()-family.Father.DateOfBirth.Year() < PARENT_THRESHOLD {

		return errors.New("invalid father date of birth")
	}

	if dateOfBirth.Year()-family.Mother.DateOfBirth.Year() < PARENT_THRESHOLD {

		return errors.New("invalid mother date of birth")
	}

	return nil
}

func (this *CommitCandidateSigningInfoRequest) validateNames() error {

	data := this.Data
	civilIdentity := &(data.CivilIndentity)

	if !isValidName(civilIdentity.Name) {

		return errorAlert("name", civilIdentity.Name)
	}

	// if !isValidName(civilIdentity.BirthPlace) {

	// 	return errorAlert("birth place")
	// }

	if !isValidName(civilIdentity.PlaceOfOrigin) {

		return errorAlert("place of origin", civilIdentity.PlaceOfOrigin)
	}

	if !isValidName(civilIdentity.Nationality) {

		return errorAlert("nationality", civilIdentity.Nationality)
	}

	family := &(data.Family)

	if !isValidName(family.Mother.Name) {

		return errorAlert("mother name", family.Mother.Name)
	}

	if !isValidName((family.Father.Name)) {

		return errorAlert("father name", family.Mother.Name)
	}

	if family.Anothers != nil {

		for _, m := range *(family.Anothers) {

			if !isValidName(m.Name) {

				return errorAlert("sibling name", family.Mother.Name)
			}
		}
	}

	return nil
}
