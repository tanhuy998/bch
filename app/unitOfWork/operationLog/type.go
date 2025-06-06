package opLog

import "context"

type (
	IInternalGeneralLogger interface {
		pushArbitrary(level string, ctx context.Context, lines []interface{})
		messure(
			level string, op string, msg string, ctx context.Context,
		) func(err error)

		pushIfError(
			level string, err error, op string, msg string, ctx context.Context,
		)

		push(
			level string, op string, msg string, ctx context.Context,
		)

		pushCond(
			level string, op string, msgIfNoErr string, ctx context.Context,
		) func(err error, msgIfErr string)

		pushCondWithMessurement(
			level string, op string, ctx context.Context,
		) func(msgIfNoErr string, err error, msgIfErr string)

		pushError(
			level string, op string, err error, defaultMsg string, ctx context.Context,
		)
	}

	IGeneralLogger interface {
		TryMergeLogs(ctx context.Context) bool

		PushCustom(level string, ctx context.Context, lines ...interface{})
		Messure(
			level string, op string, msg string, ctx context.Context,
		) func(err error)

		PushIfError(
			level string, err error, op string, msg string, ctx context.Context,
		)

		Push(
			level string, op string, msg string, ctx context.Context,
		)

		PushCond(
			level string, op string, msgIfNoErr string, ctx context.Context,
		) func(err error, msgIfErr string)

		PushCondWithMessurement(
			level string, op string, ctx context.Context,
		) func(msgIfNoErr string, err error, msgIfErr string)

		PushError(
			level string, op string, err error, defaultMsg string, ctx context.Context,
		)
	}

	ILogUseCase interface {
		PushCustom(ctx context.Context, lines ...interface{})
		Messure(
			op string, msg string, ctx context.Context,
		) func(err error)

		PushIfError(
			err error, op string, msg string, ctx context.Context,
		)

		Push(
			op string, msg string, ctx context.Context,
		)

		PushCond(
			op string, msgIfNoErr string, ctx context.Context,
		) func(err error, msgIfErr string)

		PushCondWithMessurement(
			op string, ctx context.Context,
		) func(msgIfNoErr string, err error, msgIfErr string)

		PushError(
			op string, err error, defaultMsg string, ctx context.Context,
		)
	}
)
