package usecase

import (
	"app/domain/model"
	requestPresenter "app/domain/presenter/request"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	adminService "app/service/admin"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IGetCampaignCandidateList interface {
		Execute(
			*requestPresenter.GetCampaignCandidateListRequest,
			*responsePresenter.GetCampaignCandidateListResponse,
		) (mvc.Result, error)
	}

	GetCampaignCandidateListUseCase struct {
		GetCampaignCandidateListService adminService.IGetCampaignCandidateList
	}
)

func (this *GetCampaignCandidateListUseCase) Execute(
	input *requestPresenter.GetCampaignCandidateListRequest,
	output *responsePresenter.GetCampaignCandidateListResponse,
) (mvc.Result, error) {

	var (
		candidatePivot_id string = input.PivotID
		err               error
	)

	dataPack, err := this.GetCampaignCandidateListService.Serve(
		input.CampaignUUID,
		candidatePivot_id,
		input.PageSizeLimit,
		input.IsPrev,
	)

	if err != nil {

		return nil, err
	}

	res := NewResponse()
	output.Message = "success"
	output.Data = dataPack.Data

	if input.ExposeHeader {

		output.Header.Campaign = dataPack.Header
	}

	pageNumber := common.PAGINATION_FIRST_PAGE

	err = preparePaginationNavigation[model.Candidate](output, dataPack.PaginationPack, pageNumber)

	if err != nil {

		return nil, err
	}

	err = MarshalResponseContent(output, res)

	if err != nil {

		return nil, err
	}

	return res, nil
}
