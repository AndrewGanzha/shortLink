package auth

import (
	"encoding/json"
	"fmt"
	"learnProject/configs"
	"learnProject/pkg/response"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}
type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	authHandler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("/auth/login", authHandler.Login())
	router.HandleFunc("/auth/register", authHandler.Register())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			response.Json(err.Error(), w, http.StatusBadRequest)
		}
		res := LoginResponse{
			Token: "123",
		}
		response.Json(res, w, http.StatusOK)
	}
}
