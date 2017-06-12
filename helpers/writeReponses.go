package helpers

import (
	"encoding/json"
	"net/http"
)

// WriteReponse as bytes for standard output responses
func WriteReponse(content string, headerStatus int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(headerStatus)
	w.Write([]byte(content))
}

// WriteJSONResponse as Json when requiering to deal with specific clients
func WriteJSONResponse(v interface{}, headerStatus int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(headerStatus)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + " error on encoding data\n"))
	}
}
