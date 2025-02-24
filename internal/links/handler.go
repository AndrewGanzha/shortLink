package links

import (
	"learnProject/configs"
	"net/http"
)

type LinkHandler struct {
	*configs.Config
}
type LinkHandlerDeps struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	linkHandler := &LinkHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /link", linkHandler.Create())
	router.HandleFunc("PATCH /link/{id}", linkHandler.Update())
	router.HandleFunc("DELETE /link/{id}", linkHandler.Delete())
	router.HandleFunc("GET /{alias}", linkHandler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
