package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/trmttty/ca-tech-dojo/interface/auth"
	"github.com/trmttty/ca-tech-dojo/usecase"
)

type UserHandler interface {
	Post() http.HandlerFunc
	Get() http.HandlerFunc
	Put() http.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	userHandler := userHandler{userUsecase: userUsecase}
	return &userHandler
}

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

type UserGetResponse struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}

func (handler *userHandler) Post() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var userName UserCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&userName); err != nil {
			log.Printf("Json decode error, %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		createdUser, err := handler.userUsecase.Create(userName.Name)
		if err != nil {
			log.Printf("Create user DB error, %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var token UserCreateResponse
		token.Token, err = auth.CreateToken(createdUser.ID)
		if err != nil {
			log.Printf("Create token error, %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response, err := json.MarshalIndent(&token, "", " ")
		if err != nil {
			log.Printf("Encode token error, %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(response)
	}
}

func (handler *userHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		userID := r.Context().Value("user-id").(int)

		foundUser, err := handler.userUsecase.FindByID(userID)
		if err != nil {
			log.Printf("Get user name DB error, %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userName := UserGetResponse{Name: foundUser.UserName}
		response, err := json.MarshalIndent(&userName, "", " ")
		if err != nil {
			log.Printf("Encode token error, %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(response)
	}
}

func (handler *userHandler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var userName UserUpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&userName); err != nil {
			log.Printf("Json decode error, %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userID := r.Context().Value("user-id").(int)

		_, err := handler.userUsecase.Update(userID, userName.Name)
		if err != nil {
			log.Printf("Update user name DB error, %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
