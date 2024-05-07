package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gofor-little/env"
)

const (
	ENV_HOSTS = "HOSTS"
)

var (
	cwd, err      = os.Getwd()
	isLoaded bool = false
)

func InitEnv() error {

	if isLoaded {

		return nil
	}

	if err != nil {

		return err
	}

	return env.Load(filepath.Join(cwd, ".env"))
}

func RetrieveCORSHosts() []string {

	err := InitEnv()

	if err != nil {

		panic(err)
	}

	hostString := env.Get(ENV_HOSTS, "*")

	return strings.Split(hostString, ",")
}
