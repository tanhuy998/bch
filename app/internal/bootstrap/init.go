package bootstrap

import (
	"app/internal/internal/cmd"
	libCommon "app/internal/lib/common"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("bootstrap init")
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

func init() {

	cmd.ToggleDebugMode(
		libCommon.Ternary(
			os.Getenv(ENV_DEBUG_LOG) == "true",
			true, false,
		),
	)
}

func Boot() {

}
