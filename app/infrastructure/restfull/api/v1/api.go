package v1

import (
	"app/infrastructure/restfull/api/v1/branch/assignment"
	"app/infrastructure/restfull/api/v1/branch/auth/authGenApi"
	commandgroup "app/infrastructure/restfull/api/v1/branch/auth/manipulation/commandGroup"
	"app/infrastructure/restfull/api/v1/branch/auth/manipulation/role"
	"app/infrastructure/restfull/api/v1/branch/auth/manipulation/user"
	authSignaturesApi "app/infrastructure/restfull/api/v1/branch/auth/signatures"
	"app/infrastructure/restfull/api/v1/branch/tenant"
	"app/shared/lib/iris/api"

	"github.com/kataras/iris/v12"
)

func Initialize(parent iris.Party) {

	builder := api.NewAPIBuilder(parent.Party("/v1"))

	builder.Branch("/auth",
		func(cur *api.APIBuilder) {

			cur.Branch("/gen", func(cur *api.APIBuilder) {

				authGenApi.API(cur.Party)
			})

			cur.Branch("/signatures",
				func(cur *api.APIBuilder) {

					authSignaturesApi.API(cur.Party)
				},
			)

			cur.Branch("/man",
				func(cur *api.APIBuilder) {

					cur.Branch("/users",
						func(cur *api.APIBuilder) {

							user.API(cur.Party)
						},
					)

					cur.Branch("/roles",
						func(cur *api.APIBuilder) {

							role.API(cur.Party)
						},
					)

					cur.Branch("/command",
						func(cur *api.APIBuilder) {

							cur.Branch("/groups", func(cur *api.APIBuilder) {

								commandgroup.API(cur.Party)
							})
						},
					)
				},
			)
		},
	)

	builder.Branch("tenants",
		func(cur *api.APIBuilder) {

			tenant.API(cur.Party)
		},
	)

	builder.Branch("/assigns",
		func(cur *api.APIBuilder) {

			assignment.API(cur.Party)
		},
	)
}
