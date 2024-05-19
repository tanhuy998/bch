package controller

import (
	"app/domain/model"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type CandidateController struct {
	AddnewCandidateOperation adminService.IAddNewCandidate
	DeleteCandidateOperation adminService.IDeleteCandidate
	ModifyCandidateOperation adminService.IModifyCandidate
}

func (this *CandidateController) GetCandidate() {

}

func (this *CandidateController) PostCandidate(inputCampaignUUID string, inputCandidate *model.Candidate) mvc.Response {

	err := this.AddnewCandidateOperation.Execute(inputCampaignUUID, inputCandidate)

	if err != nil {

		return BadRequest(err)
	}

	return Ok()
}

func (this *CandidateController) UpdateCandidate(inputUUID string, inputModel *model.Candidate) mvc.Response {

	err := this.ModifyCandidateOperation.Execute(inputUUID, inputModel)

	if err != nil {

		return BadRequest(err)
	}

	return Ok()
}

func (this *CandidateController) DeleteCandidate(inputUUID string) mvc.Response {

	err := this.DeleteCandidateOperation.Execute(inputUUID)

	if err != nil {

		return BadRequest(err)
	}

	return Ok()
}

func (this *CandidateController) GetCandidateByPage() {

}

func (this *CandidateController) SearchByInformation() {

}
