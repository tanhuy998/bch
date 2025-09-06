package http

var (
	cors_allowed_methods []string = []string{
		"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH",
	}
	cors_allowed_headers []string = []string{
		"Authorization", "X-Forwarded-For", "X-Forwarded-Proto", "X-Real-Ip",
	}
)

func CORS_Allowed_Methods() []string {

	return cors_allowed_methods
}

func CORS_Allowed_Headers() []string {

	return cors_allowed_headers
}
