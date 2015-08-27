package webster_test

import (
	. "github.com/ScarletTanager/webster"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*
 * You need to set key to your own API key from webster
 */

var _ = Describe("Webster", func() {
	var (
		err error
		key string
	)

	BeforeEach(func() {
		key = "a04e46f7-0cb8-41a6-a164-bd09f0046930"
		InitClient(key, false)
		err := Fetch("good")
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Initializing the Webster API client", func() {
		Context("Using a blank API key", func() {
			It("should return an error", func() {
				err = InitClient("", false)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("Using the default configuration", func() {
			It("should not return an error", func() {
				err = InitClient(key, true)
				Expect(err).NotTo(HaveOccurred())
			})

			Context("without testing the connection during init", func() {
				It("should never return an error", func() {
					err = InitClient("someboguskey", false)
					Expect(err).NotTo(HaveOccurred())
				})
			})
		})
	})

	Describe("Looking up to see whether a word exists", func() {
		Context("Word is known to exist", func() {
			It("should return true", func() {
				Expect(WordExists("good")).To(Equal(true))
			})
		})
	})

	Describe("Retrieving an iterating through a set of entries under a single word", func() {
		Context("Word exists and has multiple entries", func() {
			It("Should retrieve a set of entries with cardinality > 1", func() {
				Expect(EntryCount()).To(BeNumerically(">", 1))
			})

/*			It("Should store and return the same word", func() {
				Expect()
				}) */

			It("Should return the correct first entry", func() {
				Expect(FirstEntry().Id).To(Equal("good[1]"))
				Expect(CurrentEntry().Id).To(Equal("good[1]"))
			})

			It("Should iterate through the entries", func() {
				lastId := CurrentEntry().Id
				for NextEntry() != nil {
					Expect(CurrentEntry().Id).NotTo(Equal(lastId))
				}
			})
		})
	})

})
