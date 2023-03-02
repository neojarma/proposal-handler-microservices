package main

import (
	"log"
	"os"
	connection "user-service/connection"
	handler "user-service/handler"
	broker "user-service/redis"
	"user-service/repository"
	usecase "user-service/use-case"
)

func main() {

	mysqlConn, err := connection.GetConnectionMySQL()
	if err != nil {
		log.Println("error connection to mysql, err:", err)
	}

	repo := repository.NewRepository(mysqlConn)
	useCase := usecase.NewUserUseCase(repo)
	handler := handler.NewUserHandler(useCase)

	redisConn, err := connection.GetConnectionRedis()
	if err != nil {
		log.Println("cannot make a connection to redis, err:", err)
	}

	broker.ListenRedisEvent(redisConn, os.Getenv("USER_CHANNEL"), handler.Regist)
}
