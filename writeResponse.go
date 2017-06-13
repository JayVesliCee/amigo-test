/*
	Those functions could be extracted in an other package in case of reusability.
	It's a helper in order to send response through bytes or JSON
*/

package main

import (
	"encoding/json"
	"net/http"
)

func writeReponse(content string, headerStatus int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(headerStatus)
	w.Write([]byte(content))
}

func writeJSONResponse(v interface{}, headerStatus int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(headerStatus)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + " error on encoding data\n"))
	}
}
