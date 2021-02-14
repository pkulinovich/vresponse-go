package response_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v "github.com/pkulinovich/vresponse-go"
)

var _ = Describe("Links", func() {
	var (
		links *v.Links

		path               string
		page, limit, total uint64
	)

	JustBeforeEach(func() {
		links = v.NewLinks(path, page, limit, total)
	})

	Describe("", func() {
		Context("when total 100 items, limit 20 and page 2", func() {
			BeforeEach(func() {
				path = "http://github.com/pkulinovich/vresponse-go"
				page = 2
				limit = 20
				total = 100
			})

			It("should build first path correctly", func() {
				first := fmt.Sprintf("%s?page[number]=1", path)
				Expect(links.First).To(Equal(&first))
			})
			It("should build last path correctly", func() {
				last := fmt.Sprintf("%s?page[number]=5", path)
				Expect(links.Last).To(Equal(&last))
			})
			It("should build next path correctly", func() {
				next := fmt.Sprintf("%s?page[number]=3", path)
				Expect(links.Next).To(Equal(&next))
			})
			It("should build prev path correctly", func() {
				prev := fmt.Sprintf("%s?page[number]=1", path)
				Expect(links.Prev).To(Equal(&prev))
			})
			It("should build self path correctly", func() {
				self := fmt.Sprintf("%s?page[number]=2", path)
				Expect(links.Self).To(Equal(&self))
			})
		})
		Context("when total less limit", func() {
			BeforeEach(func() {
				path = "http://github.com/pkulinovich/vresponse-go"
				page = 1
				limit = 20
				total = 10
			})

			It("should build first path correctly", func() {
				first := fmt.Sprintf("%s?page[number]=1", path)
				Expect(links.First).To(Equal(&first))
			})
			It("should build last path correctly", func() {
				last := fmt.Sprintf("%s?page[number]=1", path)
				Expect(links.Last).To(Equal(&last))
			})
			It("next link should be nil", func() {
				Expect(links.Next).To(BeNil())
			})
			It("prev link should be nil", func() {
				Expect(links.Prev).To(BeNil())
			})
			It("should build self path correctly", func() {
				self := fmt.Sprintf("%s?page[number]=1", path)
				Expect(links.Self).To(Equal(&self))
			})
		})
	})
})
