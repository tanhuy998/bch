package envService

import (
	"os"
	"path/filepath"

	"github.com/gofor-little/env"
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
