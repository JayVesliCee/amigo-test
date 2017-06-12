package main

import (
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers", func() {

	Describe("Receive Message Handler", func() {
		Context("Valide request", func() {
			It("Should return an ID", func() {
				reader := strings.NewReader("Hello ca va")
				req, err := http.NewRequest("POST", "/messages/", reader)
				Expect(err).To(BeNil())

				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(serv.receiveMessageHandler)
				handler.ServeHTTP(rr, req)
				Expect(rr.Code).To(BeIdenticalTo(http.StatusOK))

				expected := "{\"ID\":1}\n"
				Expect(rr.Body.String()).To(BeIdenticalTo(expected))
			})
		})
	})
})
