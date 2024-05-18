package controller

import adminService "app/app/service/admin"

type CandidateController struct {
	AddnewCandidateOperation adminService.IAddNewCandidate
	DeleteCandidateOperation adminService.IDeleteCandidate
	ModifyCandidateOperation adminService.IModifyCandidate
}

func (this *CandidateController) GetCandidate() {

}

func (this *CandidateController) PostCandidate() {

}

func (this *CandidateController) UpdateCandidate() {

}

func (this *CandidateController) DeleteCandidate() {

}

func (this *CandidateController) GetCandidateByPage() {

}

func (this *CandidateController) SearchByInformation() {

}
