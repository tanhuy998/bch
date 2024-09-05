package bootstrap

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const (
	ENV_HOSTS = "HOSTS"
)

// var (
// 	cwd, err      = os.Getwd()
// 	isLoaded bool = false
// )

func init() {

	godotenv.Load()
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
