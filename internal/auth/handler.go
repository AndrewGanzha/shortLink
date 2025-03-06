package auth

import (
	"learnProject/configs"
	"learnProject/pkg/request"
	"learnProject/pkg/response"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
}
type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	authHandler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", authHandler.Login())
	router.HandleFunc("POST /auth/register", authHandler.Register())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, r)

		if err != nil {
			return
		}

		handler.AuthService.Register(body.Email, body.Password, body.Name)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[LoginRequest](&w, r)

		if err != nil {
			return
		}
		res := LoginResponse{
			Token: "123",
		}
		response.Json(res, w, http.StatusOK)
	}
}
