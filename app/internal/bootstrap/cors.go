package bootstrap

import (
	libCommon "app/internal/lib/common"
	"fmt"

	"github.com/gofor-little/env"
)

var (
	allowed_origins []string
	scheme          string
)

func GetCORSSheme() string {

	if scheme != "" {

		return scheme
	}

	scheme = libCommon.Ternary(
		env.Get("HTTPS", "false") == "true",
		"https",
		"http",
	)

	return scheme
}

func GetAllowedOrigins() []string {

	if allowed_origins != nil {

		return allowed_origins
	}

	hostNames := GetDomainNames()

	ports := RetrieveAllowedCORSPorts()

	allowed_origins = make([]string, len(hostNames)+len(ports))

	// scheme := libCommon.Ternary(
	// 	env.Get("HTTPS", "false") == "true",
	// 	"https",
	// 	"http",
	// )

	scheme := GetCORSSheme()

	for i, hName := range hostNames {

		for j, aPort := range ports {

			aPort = libCommon.Ternary(
				aPort == "",
				aPort,
				fmt.Sprintf(`:%s`, aPort),
			)

			allowed_origins[i+j] = fmt.Sprintf(`%s://%s%s`, scheme, hName, aPort)
		}
	}

	return allowed_origins
}
