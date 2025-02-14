package response

import (
	"encoding/json"
	"net/http"
)

func Json(res any, writer http.ResponseWriter, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(res)
}
