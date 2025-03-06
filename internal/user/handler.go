package user

import (
	"learnProject/pkg/request"
	"learnProject/pkg/response"
	"net/http"
)

type UserHandler struct {
	UserRepository *UserRepository
}

type UserHandlerDeps struct {
	UserRepository *UserRepository
}

func NewUserHandler(router *http.ServeMux, deps UserHandlerDeps) {
	userHandler := &UserHandler{
		UserRepository: deps.UserRepository,
	}
	router.HandleFunc("POST /user/create", userHandler.Create())
}

func (hander *UserHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[UserCreateRequest](&w, r)
		if err != nil {
			return
		}
		user := NewUser(body.Email, body.Username, body.Password)
		createdUser, err := hander.UserRepository.CreateUser(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		response.Json(createdUser, w, http.StatusOK)
	}
}
