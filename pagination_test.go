package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v "github.com/pkulinovich/vresponse-go"
)

var _ = Describe("Pagination", func() {
	var (
		pagination         *v.Pagination
		page, limit, total uint64
	)

	JustBeforeEach(func() {
		pagination = v.NewPagination(page, limit, total)
	})

	Describe("", func() {
		Context("when total 100 items, limit 20 and page 2", func() {
			BeforeEach(func() {
				page = 2
				limit = 20
				total = 100
			})
			It("should populate the arguments correctly", func() {
				Expect(pagination.TotalRecord).To(Equal(total))
				Expect(pagination.Limit).To(Equal(limit))
				Expect(pagination.PageNumber).To(Equal(page))
				Expect(pagination.CurrentPage).To(Equal(page))
			})
			It("should calculate offset correctly", func() {
				offset := limit * (page - 1)
				Expect(pagination.Offset).To(Equal(offset))
			})
			It("should calculate total pages correctly", func() {
				Expect(pagination.TotalPage).To(Equal(uint64(5)))
			})
		})
	})

})
