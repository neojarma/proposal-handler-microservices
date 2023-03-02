package main

import (
	"building-service/connection"
	"building-service/router"
	"log"
)

func main() {
	mysqlConn, err := connection.GetConnectionMySQL()
	if err != nil {
		log.Println("error while connecting to mysql, err:", err)
	}

	redisConn, err := connection.GetConnectionRedis()
	if err != nil {
		log.Println("cannot make a connection to redis, err:", err)
	}

	router.Router(mysqlConn, redisConn)
}
