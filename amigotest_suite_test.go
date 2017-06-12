package main

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var serv *service

func TestAmigoTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AmigoTest Suite")
}

var _ = BeforeSuite(func() {
	var err error

	serv, err = newService("cftest.json")
	Expect(err).NotTo(HaveOccurred())

	Expect(serv.Config).NotTo(BeNil())
	Expect(serv.DB).NotTo(BeNil())

})

var _ = AfterSuite(func() {
	serv.DB.Exec("DROP TABLE messages;")
	fmt.Println("Closing!")
	serv.Close()
})
