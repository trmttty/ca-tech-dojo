package infrastructure

import (
	"github.com/trmttty/ca-tech-dojo/domain/model"
	"github.com/trmttty/ca-tech-dojo/domain/repository"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	userRepository := UserRepository{sqlHandler}
	return &userRepository
}

func (userRepo *UserRepository) Create(user *model.User) (*model.User, error) {
	res, err := userRepo.SqlHandler.Conn.Exec("INSERT INTO users (name) VALUES (?)", user.UserName)
	if err != nil {
		return user, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return user, err
	}
	user.ID = int(id)
	return user, err
}

func (userRepo *UserRepository) Update(user *model.User) (*model.User, error) {
	_, err := userRepo.SqlHandler.Conn.Exec("UPDATE users SET name=? WHERE id=?", user.UserName, user.ID)
	return user, err
}

func (userRepo *UserRepository) FindByID(id int) (*model.User, error) {
	user := model.User{ID: id}
	err := userRepo.SqlHandler.Conn.QueryRow("SELECT name, created_at, updated_at from users WHERE id = ?", id).Scan(&user.UserName, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}
