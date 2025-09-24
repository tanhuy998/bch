package repositoryAPI

type (
	IStatisticRepsitory interface {
		Statistic(fn FilterFunc) IStatisticUnit
	}

	IStatisticUnit interface {
		IStatisticAffectedCountableUnit
	}

	ISelfStatisticable interface {
		SelfStatistic() IStatisticAffectedCountableUnit
	}
)
