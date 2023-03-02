package main

import (
	"auth-service/connection"
	"auth-service/router"
	"log"
)

func main() {
	mysqlConn, err := connection.GetConnectionMySQL()
	if err != nil {
		log.Println("error while connecting to mysql, err:", err)
	}

	router.Router(mysqlConn)
}
