package auth

import (
	"learnProject/configs"
	"learnProject/pkg/jwt"
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

		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := jwt.NewJWT(handler.Auth.Secret).Create(email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := RegisterResponse{
			Token: token,
		}
		response.Json(res, w, http.StatusOK)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, r)

		if err != nil {
			return
		}
		email, err := handler.AuthService.Login(body.Email, body.Password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(handler.Auth.Secret).Create(email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := LoginResponse{
			Token: token,
		}
		response.Json(res, w, http.StatusOK)
	}
}
