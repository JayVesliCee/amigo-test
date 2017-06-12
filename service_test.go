package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {
	Describe("Testing Service", func() {

		Context("Starting a new service", func() {
			It("Should not return an error", func() {
				s, err := newService("cftest.json")
				Expect(err).To(BeNil())
				Expect(s).NotTo(BeNil())

				Expect(s.Config).NotTo(BeNil())
				Expect(s.DB).NotTo(BeNil())
			})
		})

		Context("Starting a new service with wrong file", func() {
			It("Should return an error", func() {
				s, err := newService("awf.json")
				Expect(err).NotTo(BeNil())
				Expect(s).To(BeNil())
			})
		})

		Context("Starting a new service with wrong DB", func() {
			It("Should return an error", func() {
				s, err := newService("config-test.json")
				Expect(err).NotTo(BeNil())
				Expect(s).To(BeNil())
			})
		})
	})
})
