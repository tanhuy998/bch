package main

import (
	v1 "app/infrastructure/http"
	"app/infrastructure/http/api/v1/config"
	"app/internal/bootstrap"
	"app/internal/db"
	"log"

	"os"
	"path"

	"github.com/gofor-little/env"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kataras/iris/v12"
)

const (
	ENV_HOSTS = "HOSTS"
)

var (
	host_names           []string = make([]string, 0)
	cors_allowed_methods []string = []string{
		"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH",
	}
	cors_allowed_headers []string = []string{
		"Authorization", "X-Forwarded-For", "X-Forwarded-Proto", "X-Real-Ip",
	}
	server_ssl_cert string
	server_ssl_key  string
)

type (
	cache_log_t struct {
		socketio.Conn
	}
)

func (this *cache_log_t) Write(b []byte) (n int, err error) {

	if this.Conn == nil {

		return
	}

	this.Conn.Emit("log", b)

	return
}

func init() {

	//godotenv.Load()

	db.Init()
	readSSlCert()

	//host_names = bootstrap.RetrieveCORSHosts()
	//config.InitializeAuthEncryptionData()

	err := db.CheckDBConnection()

	if err != nil {

		panic(err)
	}

	db := db.GetDB()

	config.InitDomainIndexes(db)
	host_names = append(host_names, bootstrap.GetDomainNames()...)
}

func main() {

	var app *iris.Application = iris.New()

	defer config.ConfigureLogger(app).Close()

	// app.UseRouter(
	// 	cors.New(cors.Options{
	// 		AllowedOrigins:   bootstrap.GetAllowedOrigins(),
	// 		AllowedMethods:   cors_allowed_methods,
	// 		AllowedHeaders:   cors_allowed_headers,
	// 		AllowCredentials: true,
	// 	}),
	// )

	v1.Initialize(app)

	log.Fatal(
		app.Listen(
			env.Get("HTTP_PORT", ":80"),
			iris.WithoutBodyConsumptionOnUnmarshal,
			iris.WithOptimizations,
		),
	)
}

func initWebsocket() {

	// socketServer := socketio.NewServer(nil)

	// socketServer.OnConnect("/monitor/cache", func(c socketio.Conn) error {

	// 	fmt.Printf("new cache monitor %s", c.ID())
	// 	memoryCache.AddLogListener(c.ID(), &cache_log_t{c})

	// 	return nil
	// })

	// socketServer.OnDisconnect("/monitor/cache", func(c socketio.Conn, s string) {

	// 	fmt.Printf("cache monitor %s exit", c.ID())
	// 	memoryCache.RemoveListener(c.ID())
	// })

	// func() {

	// 	srv := &http.Server{
	// 		Addr:    ":4000",
	// 		Handler: socketServer,
	// 	}

	// 	go log.Fatal(srv.ListenAndServe())

	// }()
	//http.Handle("/monitor/cache", socketServer)
}

func readSSlCert() {

	__dir, err := os.Getwd()

	if err != nil {

		panic(err)
	}

	d, err := os.ReadFile(path.Join(__dir, "cert.pem"))

	if err != nil {

		panic(err)
	}

	server_ssl_cert = string(d)

	d, err = os.ReadFile(path.Join(__dir, "key.pem"))

	if err != nil {

		panic(err)
	}

	server_ssl_key = string(d)
}
