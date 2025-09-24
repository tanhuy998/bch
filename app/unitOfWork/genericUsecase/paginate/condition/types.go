package condition

import (
	"app/internal/db/query"
	repositoryAPI "app/repository/api"
)

type (
	IQueryCondtionMaterializer interface {
		AsFilter(fn repositoryAPI.FilterFunc) query.IDataConditionExpressionResult
		AsConditionExpression(fn repositoryAPI.MatchFunc) query.IDataConditionExpressionResult
	}

	SetQueryCondtionFunc func(condition IQueryCondtionMaterializer) IConditionsOperation

	IQueryConditionPaginator interface {
		//SetQueryCondition(fn SetQueryCondtionFunc)
		MatchCondition(
			condition query.IGeneralDataConditionExpression, //IQueryCondtionMaterializer,
		) query.IDataConditionExpressionResult
	}

	IConditionsOperation interface {
		ApplyQueryCondition()
	}
)
