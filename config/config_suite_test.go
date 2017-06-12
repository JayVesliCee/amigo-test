package config_test

import (
	"github.com/amigo-test/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config", func() {
	Context("Test with correct file", func() {
		It("Should not return an error", func() {
			conf, err := config.LoadConfig("../config.json")
			Expect(err).To(BeNil())
			Expect(conf).NotTo(BeNil())
		})
	})

	Context("Test with incorrect file", func() {
		It("Should return an error", func() {
			conf, err := config.LoadConfig("asdlkfh.json")
			Expect(err).NotTo(BeNil())
			Expect(conf).To(BeNil())
		})
	})

	Context("Test with wrong JSON", func() {
		It("Should return an error", func() {
			conf, err := config.LoadConfig("conf-test.json")
			Expect(err).NotTo(BeNil())
			Expect(conf).To(BeNil())
		})
	})
})
