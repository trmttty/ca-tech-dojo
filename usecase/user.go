package usecase

import (
	"github.com/trmttty/ca-tech-dojo/domain/model"
	"github.com/trmttty/ca-tech-dojo/domain/repository"
)

type UserUsecase interface {
	FindByID(id int) (*model.User, error)
	Create(name string) (*model.User, error)
	Update(id int, name string) (*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	userUsecase := userUsecase{userRepo: userRepo}
	return &userUsecase
}

func (usecase *userUsecase) FindByID(id int) (*model.User, error) {
	user, err := usecase.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (usecase *userUsecase) Create(name string) (*model.User, error) {
	user, err := model.NewUser(name)
	if err != nil {
		return nil, err
	}

	createdUser, err := usecase.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (usecase *userUsecase) Update(id int, name string) (*model.User, error) {
	targetUser, err := usecase.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetUser.Set(name)
	if err != nil {
		return nil, err
	}

	updateUser, err := usecase.userRepo.Update(targetUser)
	if err != nil {
		return nil, err
	}

	return updateUser, err
}
