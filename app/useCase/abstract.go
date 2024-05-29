package usecase

import (
	"app/domain/model"
	responsePresenter "app/domain/presenter/response"
	"app/internal/common"
	"app/repository"
	"encoding/json"
	"errors"

	"github.com/kataras/iris/v12/mvc"
)

type (
	IResponseWrapper interface {
	}

	ResponseWrapper[T any] struct {
		Target T
	}

	ActionResultUseCase struct {
	}

	ModelInterfaceForPagination interface {
		model.IModel
	}
)

func (this *ActionResultUseCase) MarshallOutput(resultContent interface{}, response *mvc.Response) error {

	res := NewResponse()

	return MarshalResponseContent(resultContent, res)
}

func (this *ActionResultUseCase) NewActionResponse() *mvc.Response {

	return NewResponse()
}

func NewResponse() *mvc.Response {

	return &mvc.Response{
		ContentType: "application/json",
	}
}

func MarshalResponseContent(context interface{}, res *mvc.Response) error {

	data, err := json.Marshal(context)

	if err != nil {

		return err
	}

	res.Content = data

	return nil
}

func resolveNext[Model_T ModelInterfaceForPagination](
	output responsePresenter.IPaginationResult,
	dataPack *repository.PaginationPack[Model_T],
	pageNumber common.PaginationPage,
) error {

	lastIndex := len(dataPack.Data) - 1

	if lastIndex <= 0 {

		return nil
	}

	lastElement, ok := any(dataPack.Data[lastIndex]).(Model_T)

	if !ok {

		return errors.New("")
	}

	output.GetNavigation().Next = lastElement.GetObjectID().Hex()

	return nil
}

func resolvePrev[Model_T ModelInterfaceForPagination](
	output responsePresenter.IPaginationResult,
	dataPack *repository.PaginationPack[Model_T],
	pageNumber common.PaginationPage,
) error {

	// firstIndex := 0

	// firstElement, ok := any(dataPack.Data[firstIndex]).(Model_T)

	// if !ok {

	// 	return errors.New("")
	// }

	return nil
}

func preparePaginationNavigation[Model_T ModelInterfaceForPagination](
	output responsePresenter.IPaginationResult,
	dataPack *repository.PaginationPack[Model_T],
	pageNumber common.PaginationPage,
) error {

	err := resolveNext[Model_T](output, dataPack, pageNumber)

	if err != nil {

		return err
	}

	return resolvePrev[Model_T](output, dataPack, pageNumber)
}
