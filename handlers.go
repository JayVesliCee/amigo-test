/*
	Handler function will manage the request / response through the routing.
	We use service method in order to avoid the DB transition through a context.
*/

package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pressly/chi"
)

type message struct {
	ID      uint
	Content string
}

// retrieveMessageHandler simply parse the URL param and return the associated message as content.
func (s *service) retrieveMessageHandler(w http.ResponseWriter, r *http.Request) {
	m := &message{}
	messageID := strings.Title(chi.URLParam(r, "id"))

	err := s.DB.First(&m, messageID).Error
	if err != nil {
		writeReponse(err.Error()+" On retrieving message from ID\n", http.StatusBadRequest, w)
	}

	writeReponse(m.Content+"\n", http.StatusOK, w)
}

// receiveMessageHandler is a basic read and store handlerFunc.
// The particularity is to use an anonyme structure in order to return the wanted label
func (s *service) receiveMessageHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		writeReponse(err.Error()+" On reading body\n", http.StatusBadRequest, w)
	}

	m := &message{
		Content: string(bodyBytes),
	}
	err = s.DB.Create(&m).Error
	if err != nil {
		writeReponse(err.Error()+" On creating message in DB\n", http.StatusInternalServerError, w)
	}

	IDToReturn := struct {
		ID uint
	}{
		m.ID,
	}
	writeJSONResponse(IDToReturn, http.StatusOK, w)
}
