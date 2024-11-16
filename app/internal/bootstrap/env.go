package bootstrap

import (
	"app/internal/cli"
	"os"
	"strings"

	"github.com/gofor-little/env"
	"github.com/joho/godotenv"
)

const (
	ENV_HOSTS              = "HOSTS"
	ENV_ALLOWED_CORS_PORTS = "ALLOWED_CORS_PORTS"
	ENV_PROJECT_STAGE      = "PROJECT_STAGE"
	ENV_HTTP_PORT          = "HTTP_PORT"
	ENV_HTTPS              = "HTTPS"
)

const (
	ENV_SSL_CERT_DIR             = "SSL_CERT_DIR"
	ENV_SSL_KEY_DIR              = "SSL_KEY_DIR"
	ENV_AUTH_JWT_PUBLIC_KEY_DIR  = "AUTH_JWT_PUBLIC_KEY_DIR"
	ENV_AUTH_JWT_PRIVATE_KEY_DIR = "AUTH_JWT_PRIVATE_KEY_DIR"
	ENV_HMAC_SECRET              = "HMAC_SECRET"
)

const (
	ENV_APP_ID     = "APP_ID"
	ENV_API_KEY    = "API_KEY"
	ENV_APP_SECRET = "APP_SECRET"
	ENV_APP_NAME   = "APP_NAME"
)

const (
	ENV_TRACE_LOG         = "TRACE_LOG"
	ENV_CACHE_LOG         = "CACHE_LOG"
	ENV_OP_TRACE_DURATION = "OP_TRACE_DURATION"
)

const (
	ENV_MONGOD_CONN_STR   = "MONGOD_CONN_STR"
	ENV_MONGOD_CREDENTIAL = "MONGOD_CREDENTIAL"
	ENV_MONGOD_DB_NAME    = "MONGOD_DB_NAME"
)

const (
	ENV_AUTH_INTERNAL = "AUTH_INTERNAL"
	ENV_AUTH_HEADER   = "AUTH_HEADER"
)

const (
	ENV_TEST       = "TEST"
	ENV_TEST_LOGIN = "TEST_LOGIN"
)

var (
	host_names            []string
	host_names_dictionary map[string]bool = make(map[string]bool)
	allowed_cors_ports    []string
)

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

func init() {

	readCLIFlags()
	parseTestLoginFlags()
}

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

func IsHTTPS() bool {

	return env.Get(ENV_HTTPS, "false") == "true	"
}

func RetrieveCORSHosts() []string {

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

func readCLIFlags() {

	if cli.TraceLog() {

		os.Setenv(ENV_TRACE_LOG, "true")
	}
}
