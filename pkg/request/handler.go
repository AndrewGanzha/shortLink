package request

import (
	"learnProject/pkg/response"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		response.Json(err.Error(), *w, http.StatusBadRequest)
		return nil, err
	}
	err = IsValid(body)

	if err != nil {
		response.Json(err.Error(), *w, http.StatusBadRequest)
		return nil, err
	}

	return &body, nil
}
