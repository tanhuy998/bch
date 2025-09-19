package tls

import (
	"app/main/internal/dependencies/log"
	"os"
	"path"
)

func GetSSLCert() string {

	__dir, err := os.Getwd()

	if err != nil {

		panic(err)
	}

	log.Main().Println("Read TLS certififate.")

	d, err := os.ReadFile(path.Join(__dir, "cert.pem"))

	if err != nil {

		panic(err)
	}
	log.Main().Println("TLS certificate accquired")
	return string(d)
}

func GetSSLKey() string {

	__dir, err := os.Getwd()

	if err != nil {

		panic(err)
	}

	log.Main().Println("Read TLS key.")

	d, err := os.ReadFile(path.Join(__dir, "key.pem"))

	if err != nil {

		panic(err)
	}

	log.Main().Println("TLS key accquired")

	return string(d)
}
