package bootstrap

import (
	libCommon "app/internal/lib/common"
	"log"
	"os"
	"strconv"
)

var (
	at_test_flag  bool
	rt_test_flag  bool
	gen_test_flag bool
)

func IsTestingAccessToken() bool {

	return at_test_flag
}

func IsTestingRefreshToken() bool {

	return rt_test_flag
}

func IsTestingGenToken() bool {

	return gen_test_flag
}

func parseTestLoginFlags() {

	test_env := os.Getenv(ENV_TEST_LOGIN)

	flags, err := strconv.Atoi(test_env)

	if err != nil {

		log.Default().Fatal("error occur when reading login testing flags", err.Error())
	}

	if flags > 7 {

		log.Default().Fatal("invalid login test value")
	}

	err = os.Setenv(ENV_TEST, "true")

	if err != nil {

		panic(err)
	}

	at_test_flag = libCommon.Ternary(flags&0x1 > 0, true, false)
	rt_test_flag = libCommon.Ternary(flags&0x2 > 0, true, false)
	gen_test_flag = libCommon.Ternary(flags&0x4 > 0, true, false)
}
