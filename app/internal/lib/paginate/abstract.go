package paginate

// import (
// 	"app/model"
// 	"app/repository"
// 	"encoding/json"

// 	"github.com/kataras/iris/v12/mvc"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type PaginationPage int

// const (
// 	PAGINATION_FIRST_PAGE PaginationPage = 0
// 	PAGINATION_LAST_PAGE  PaginationPage = 0xffffffff
// )

// type (
// 	// IResponsePresenter interface {
// 	// 	Bind(ctx iris.Context) error
// 	// }

// 	IResponseWrapper interface {
// 	}

// 	ResponseWrapper[T any] struct {
// 		Target T
// 	}

// 	ActionResultUseCase struct {
// 	}

// 	/*
// 		Interface that is is used for pagination purposes.
// 		Because pagination functions just accept some specific
// 		intefaces as generic type.
// 		Therefore, it's implementations must implement this interface
// 		as value receiver in order to not catching compile
// 		time error when passing the model type to a any
// 		pagination functions.
// 	*/
// 	ModelInterfaceForPagination interface {
// 		model.IModel
// 	}
// )

// type (
// 	IEmptyResponse any

// 	IPaginationResult interface {
// 		GetNavigation() *PaginationNavigation
// 		SetTotalCount(int64)
// 	}

// 	NavigationQuery struct {
// 		Cursor primitive.ObjectID `json:"p_pivot,omitempty"`
// 		Limit  *int               `json:"p_limit,omitempty"`
// 		IsPrev bool               `json:"p_prev,omitempty"`
// 	}

// 	PaginationNavigation struct {
// 		CurrentPage int              `json:"page"`
// 		Previous    *NavigationQuery `json:"previous,omitempty"`
// 		Next        *NavigationQuery `json:"next,omitempty"`
// 	}
// )

// func (this *ActionResultUseCase) MarshallOutput(resultContent interface{}, response *mvc.Response) error {

// 	res := NewResponse()

// 	return MarshalResponseContent(resultContent, res)
// }

// func (this *ActionResultUseCase) NewActionResponse() *mvc.Response {

// 	return NewResponse()
// }

// func NewResponse() *mvc.Response {

// 	return &mvc.Response{
// 		ContentType: "application/json",
// 	}
// }

// func MarshalResponseContent(context interface{}, res *mvc.Response) error {

// 	data, err := json.Marshal(context)

// 	if err != nil {

// 		return err
// 	}

// 	res.Content = data

// 	return nil
// }

// func resolveNext[Model_T ModelInterfaceForPagination](
// 	output IPaginationResult,
// 	dataPack *repository.PaginationPack[Model_T],
// 	pageNumber PaginationPage,
// ) error {

// 	if len(dataPack.Data) == 0 {

// 		return nil
// 	}

// 	lastIndex := len(dataPack.Data) - 1

// 	if lastIndex <= 0 {

// 		return nil
// 	}
// 	/*
// 		implementation state is checked at compile time,
// 		no need any assertions at runtime
// 	*/
// 	lastElement := dataPack.Data[lastIndex]

// 	output.GetNavigation().Next = &NavigationQuery{
// 		Cursor: (*lastElement).GetObjectID(),
// 	}

// 	return nil
// }

// func resolvePrev[Model_T ModelInterfaceForPagination](
// 	output IPaginationResult,
// 	dataPack *repository.PaginationPack[Model_T],
// 	pageNumber PaginationPage,
// ) error {

// 	if len(dataPack.Data) == 0 {

// 		return nil
// 	}

// 	/*
// 		implementation state is checked at compile time,
// 		no need any assertions at runtime
// 	*/
// 	firstElement := dataPack.Data[0]

// 	output.GetNavigation().Previous = &NavigationQuery{
// 		Cursor: (*firstElement).GetObjectID(),
// 		IsPrev: true,
// 	}

// 	return nil
// }

// /*
// preparePaginationNavigation expects the input generic type implements
// ModelInterfaceForPagination and whose methods must be implemeted as value
// receiver
// */
// func preparePaginationNavigation[Model_T ModelInterfaceForPagination](
// 	output IPaginationResult,
// 	dataPack *repository.PaginationPack[Model_T],
// 	pageNumber PaginationPage,
// ) error {

// 	output.SetTotalCount(dataPack.Count)

// 	err := resolveNext[Model_T](output, dataPack, pageNumber)

// 	if err != nil {

// 		return err
// 	}

// 	return resolvePrev[Model_T](output, dataPack, pageNumber)
// }
