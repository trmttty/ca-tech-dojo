package injector

import (
	"github.com/trmttty/ca-tech-dojo/pkg/domain/repository"
	"github.com/trmttty/ca-tech-dojo/pkg/infrastructure"
	"github.com/trmttty/ca-tech-dojo/pkg/interface/handler"
	"github.com/trmttty/ca-tech-dojo/pkg/usecase"
)

func InjectDB() infrastructure.SqlHandler {
	sqlHandler := infrastructure.NewSqlHandler()
	return sqlHandler
}

func InjectUserRepository() repository.UserRepository {
	sqlHandler := InjectDB()
	return infrastructure.NewUserRepository(sqlHandler)
}

func InjectUserUsecase() usecase.UserUsecase {
	userRepo := InjectUserRepository()
	return usecase.NewUserUsecase(userRepo)
}

func InjectUserHandler() handler.UserHandler {
	userUsecase := InjectUserUsecase()
	return handler.NewUserHandler(userUsecase)
}
