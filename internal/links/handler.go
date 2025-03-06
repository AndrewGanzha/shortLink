package links

import (
	"gorm.io/gorm"
	"learnProject/pkg/middleware"
	"learnProject/pkg/request"
	"learnProject/pkg/response"
	"net/http"
	"strconv"
)

type LinkHandler struct {
	LinkRepository *LinkRepository
}
type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	linkHandler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", linkHandler.Create())
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(linkHandler.Update()))
	router.HandleFunc("DELETE /link/{id}", linkHandler.Delete())
	router.HandleFunc("GET /{hash}", linkHandler.GoTo())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		for {
			existedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		response.Json(createdLink, w, http.StatusOK)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		response.Json(link, w, http.StatusOK)
	}
}

func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		_, err = handler.LinkRepository.GetById(uint(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = handler.LinkRepository.Delete(uint(id))

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(nil, w, http.StatusOK)
	}
}

func (handler *LinkHandler) FindById(id uint64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
