package repository

import (
	"github.com/trmttty/ca-tech-dojo/pkg/domain/model"
)

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	FindByID(id int) (*model.User, error)
}
