package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Context("Test with correct file", func() {
		It("Should not return an error", func() {
			conf, err := loadConfig("config.json")
			Expect(err).To(BeNil())
			Expect(conf).NotTo(BeNil())
		})
	})

	Context("Test with incorrect file", func() {
		It("Should return an error", func() {
			conf, err := loadConfig("asdlkfh.json")
			Expect(err).NotTo(BeNil())
			Expect(conf).To(BeNil())
		})
	})

	Context("Test with wrong JSON", func() {
		It("Should return an error", func() {
			conf, err := loadConfig("conf-test.json")
			Expect(err).NotTo(BeNil())
			Expect(conf).To(BeNil())
		})
	})
})
