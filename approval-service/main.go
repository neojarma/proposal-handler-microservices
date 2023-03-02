package main

import (
	"approval-service/connection"
	"log"

	"approval-service/router"
)

func main() {
	mysqlConn, err := connection.GetConnectionMySQL()
	if err != nil {
		log.Println("error connection to mysql, err:", err)
	}

	redisConn, err := connection.GetConnectionRedis()
	if err != nil {
		log.Println("cannot make a connection to redis, err:", err)
	}

	router.Router(mysqlConn, redisConn)

}
