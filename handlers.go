package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/pressly/chi"
)

type message struct {
	ID      uint
	Content string
}

func extractContext(r *http.Request) *gorm.DB {
	DB := r.Context().Value(psqlDB).(*gorm.DB)

	return DB
}

func retrieveMessageHandler(w http.ResponseWriter, r *http.Request) {
	messageID := strings.Title(chi.URLParam(r, "id"))
	db := extractContext(r)
	m := &message{}

	err := db.First(&m, messageID).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write([]byte("error on retrieving message from ID"))
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
	w.Write([]byte(m.Content + "\n"))
}

func receiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write([]byte("error on encoding data"))
	}
	m := &message{
		Content: string(bodyBytes),
	}

	db := extractContext(r)
	err = db.Create(&m).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write([]byte("error on creating message in DB"))
	}

	idToReturn := struct {
		ID uint
	}{
		m.ID,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if err = json.NewEncoder(w).Encode(idToReturn); err != nil {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write([]byte("error on encoding data"))
	}
}
