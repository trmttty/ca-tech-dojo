package main

import (
	"github.com/trmttty/ca-tech-dojo/handler"
	"github.com/trmttty/ca-tech-dojo/infrastructure"
	"github.com/trmttty/ca-tech-dojo/usecase"
)

func main() {
	sqlHandler := infrastructure.NewSqlHandler()
	userRepository := infrastructure.NewUserRepository(sqlHandler)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	handler.Run(userHandler)
}
