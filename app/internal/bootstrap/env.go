package bootstrap

import (
	"os"
	"strings"

	"github.com/gofor-little/env"
	"github.com/joho/godotenv"
)

const (
	ENV_HOSTS              = "HOSTS"
	ENV_ALLOWED_CORS_PORTS = "ALLOWED_CORS_PORTS"
	ENV_PROJECT_STAGE      = "PROJECT_STAGE"
	ENV_APP_NAME           = "APP_NAME"
	ENV_HTTPS              = "HTTPS"
)

var (
	host_names            []string
	host_names_dictionary map[string]bool = make(map[string]bool)
	allowed_cors_ports    []string
)

// var (
// 	cwd, err      = os.Getwd()
// 	isLoaded bool = false
// )

func GetDomainNames() []string {

	return host_names
}

func GetAppName() string {

	s := os.Getenv(ENV_APP_NAME)

	if s == "" {

		return "bch"
	}

	return s
}

func HasHostName(name string) bool {

	if _, ok := host_names_dictionary[name]; !ok {

		return false
	}

	return true
}

func init() {

	defer ignorePanicWhenUnitTesting()

	err := godotenv.Load()

	if err != nil {

		panic("error while loading env: " + err.Error())
	}

	host_names = RetrieveCORSHosts()

	for _, val := range host_names {

		host_names_dictionary[val] = true
	}

	initializeAuthEncryptionData()

}

func IsHTTPS() bool {

	return env.Get(ENV_HTTPS, "false") == "true	"
}

// func InitEnv() error {

// 	if isLoaded {

// 		return nil
// 	}

// 	if err != nil {

// 		return err
// 	}

// 	return env.Load(filepath.Join(cwd, ".env"))
// }

func RetrieveCORSHosts() []string {

	// err := InitEnv()

	// if err != nil {

	// 	panic(err)
	// }

	hostString := os.Getenv(ENV_HOSTS) // env.Get(ENV_HOSTS, "*")

	return strings.Split(hostString, ",")
}

func RetrieveAllowedCORSPorts() []string {

	str := env.Get(ENV_ALLOWED_CORS_PORTS, "")

	return strings.Split(str, ",")
}

func ignorePanicWhenUnitTesting() {

	r := recover()

	switch {
	case r == nil:
	case os.Getenv(ENV_PROJECT_STAGE) == "unit-testing":
	default:
		panic(r)
	}
}
