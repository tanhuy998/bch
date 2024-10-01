package presenter

import (
	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IResponsePresenter interface {
		Bind(ctx iris.Context) error
	}

	IEmptyResponse any

	IPaginationResult interface {
		GetNavigation() *PaginationNavigation
		SetTotalCount(int64)
	}

	NavigationQuery struct {
		Cursor primitive.ObjectID `json:"p_pivot,omitempty"`
		Limit  *int               `json:"p_limit,omitempty"`
		IsPrev bool               `json:"p_prev,omitempty"`
	}

	PaginationNavigation struct {
		CurrentPage int              `json:"page"`
		Previous    *NavigationQuery `json:"previous,omitempty"`
		Next        *NavigationQuery `json:"next,omitempty"`
	}
)
