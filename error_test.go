package response_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v "github.com/pkulinovich/vresponse-go"
	"net/http"
)

var _ = Describe("Error", func() {
	var (
		err400 *v.Error
	)

	BeforeEach(func() {
		err400 = &v.Error{
			Title:   http.StatusText(http.StatusBadRequest),
			Details: "HTTP 400 Bad Request",
			Status:  http.StatusBadRequest,
			Code:    "INVALID_QUERY_PARAMS",
		}
	})

	Describe("check 400 bad request error", func() {
		Context("when parses successfully", func() {
			It("should populate the arguments correctly", func() {
				Expect(err400.Title).To(Equal(http.StatusText(http.StatusBadRequest)))
				Expect(err400.Details).To(Equal("HTTP 400 Bad Request"))
				Expect(err400.Status).To(Equal(http.StatusBadRequest))
				Expect(err400.Code).To(Equal("INVALID_QUERY_PARAMS"))
			})
		})
	})
})
