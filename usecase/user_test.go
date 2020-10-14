package usecase

import (
	"testing"
	"time"

	"github.com/trmttty/ca-tech-dojo/domain/model"
	mock "github.com/trmttty/ca-tech-dojo/mock/mock_user_repository"

	"github.com/golang/mock/gomock"
)

func TestFindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testMock := mock.NewMockUserRepository(ctrl)
	expected := &model.User{
		ID:        1,
		UserName:  "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	testMock.EXPECT().FindByID(1).Return(expected, nil)

	testUsecase := NewUserUsecase(testMock)
	user, err := testUsecase.FindByID(1)
	if err != nil {
		t.Error("failed Test")
	}
	if user != expected {
		t.Error("failed Test")
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testMock := mock.NewMockUserRepository(ctrl)
	expected := &model.User{
		UserName: "test",
	}
	testMock.EXPECT().Create(expected).Return(expected, nil)

	testUsecase := NewUserUsecase(testMock)
	user, err := testUsecase.Create("test")
	if err != nil {
		t.Error("failed Test")
	}
	if user != expected {
		t.Error("failed Test")
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testMock := mock.NewMockUserRepository(ctrl)
	expected := &model.User{
		ID:       1,
		UserName: "test",
	}
	testMock.EXPECT().FindByID(1).Return(expected, nil)
	testMock.EXPECT().Update(expected).Return(expected, nil)

	testUsecase := NewUserUsecase(testMock)
	user, err := testUsecase.Update(1, "test")
	if err != nil {
		t.Error("failed Test")
	}
	if user != expected {
		t.Error("failed Test")
	}
}
