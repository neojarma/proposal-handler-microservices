package main

import (
	connection "api-gateway/connection"
	"api-gateway/router"
	"log"
)

func main() {
	redisConn, err := connection.GetConnectionRedis()
	if err != nil {
		log.Println("cannot make a connection to redis, err:", err)
	}

	router.Router(redisConn)
}
