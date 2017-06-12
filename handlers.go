package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/amigo-test/helpers"
	"github.com/pressly/chi"
)

type message struct {
	ID      uint
	Content string
}

func (s *service) retrieveMessageHandler(w http.ResponseWriter, r *http.Request) {
	m := &message{}
	messageID := strings.Title(chi.URLParam(r, "id"))

	err := s.DB.First(&m, messageID).Error
	if err != nil {
		helpers.WriteReponse(err.Error()+" On retrieving message from ID\n", http.StatusBadRequest, w)
	}

	helpers.WriteReponse(m.Content+"\n", http.StatusOK, w)
}

func (s *service) receiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		helpers.WriteReponse(err.Error()+" On reading body\n", http.StatusBadRequest, w)
	}

	m := &message{
		Content: string(bodyBytes),
	}
	err = s.DB.Create(&m).Error
	if err != nil {
		helpers.WriteReponse(err.Error()+" On creating message in DB\n", http.StatusInternalServerError, w)
	}

	IDToReturn := struct {
		ID uint
	}{
		m.ID,
	}
	helpers.WriteJSONResponse(IDToReturn, http.StatusOK, w)
}
