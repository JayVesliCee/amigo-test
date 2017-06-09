package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/amigo-test/helpers"
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
		helpers.WriteReponse(err.Error()+" On retrieving message from ID\n", http.StatusBadRequest, w)
	}

	helpers.WriteReponse(m.Content+"\n", http.StatusOK, w)
}

func receiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.WriteReponse(err.Error()+" On reading body\n", http.StatusBadRequest, w)
	}

	m := &message{
		Content: string(bodyBytes),
	}
	db := extractContext(r)
	err = db.Create(&m).Error
	if err != nil {
		helpers.WriteReponse(err.Error()+" On creating message in DB\n", http.StatusInternalServerError, w)
	}

	idToReturn := struct {
		ID uint
	}{
		m.ID,
	}
	helpers.WriteJSONResponse(idToReturn, http.StatusOK, w)
}
