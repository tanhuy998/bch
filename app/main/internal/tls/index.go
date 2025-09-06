package tls

import (
	"os"
	"path"
)

// var (
// 	// host_names           []string = make([]string, 0)
// 	// cors_allowed_methods []string = []string{
// 	// 	"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH",
// 	// }
// 	// cors_allowed_headers []string = []string{
// 	// 	"Authorization", "X-Forwarded-For", "X-Forwarded-Proto", "X-Real-Ip",
// 	// }
// 	server_ssl_cert string
// 	server_ssl_key  string
// )

// func init() {

// 	ReadSSlCert()
// }

// func ReadSSlCert() {

// 	__dir, err := os.Getwd()

// 	if err != nil {

// 		panic(err)
// 	}

// 	d, err := os.ReadFile(path.Join(__dir, "cert.pem"))

// 	if err != nil {

// 		panic(err)
// 	}

// 	server_ssl_cert = string(d)

// 	d, err = os.ReadFile(path.Join(__dir, "key.pem"))

// 	if err != nil {

// 		panic(err)
// 	}

// 	server_ssl_key = string(d)
// }

func GetSSLCert() string {

	__dir, err := os.Getwd()

	if err != nil {

		panic(err)
	}

	d, err := os.ReadFile(path.Join(__dir, "cert.pem"))

	if err != nil {

		panic(err)
	}

	return string(d) //server_ssl_cert
}

func GetSSLKey() string {

	__dir, err := os.Getwd()

	if err != nil {

		panic(err)
	}

	d, err := os.ReadFile(path.Join(__dir, "key.pem"))

	if err != nil {

		panic(err)
	}

	return string(d) // server_ssl_key
}
