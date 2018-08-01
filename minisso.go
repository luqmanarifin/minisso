package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/bukalapak/packen/metric"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/luqmanarifin/minisso/database"
	"github.com/luqmanarifin/minisso/service"
)

//Healthz - health check
func Healthz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"ok\"}"))
}

// Metric is used to control the flow of GET /metrics endpoint
func Metric(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	metric.Handler(w, r)
}

func main() {
	openEnv()

	mysqlOption := getMysqlOption()
	// redisOption := getRedisOption()
	userService := service.NewUserService(mysqlOption)

	router := httprouter.New()
	router.GET("/healthz", Healthz)
	router.GET("/metrics", Metric)

	// router.GET("/signup", userService.Signup)
	router.POST("/signup", userService.Signup)

	fmt.Println("Starting HTTP Receiver")
	http.ListenAndServe(":1234", router)
}

func openEnv() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func getMysqlOption() database.MysqlOption {
	return database.MysqlOption{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
		Charset:  os.Getenv("MYSQL_CHARSET"),
	}
}

func getRedisOption() database.RedisOption {
	redisUrl, err := url.Parse(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	password, _ := redisUrl.User.Password()
	redisOpt := database.RedisOption{
		Host:     redisUrl.Hostname(),
		Port:     redisUrl.Port(),
		Password: password,
		Database: 0,
	}
	log.Printf("redis opt: %v", redisOpt)
	return redisOpt
}
