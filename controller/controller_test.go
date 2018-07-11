package controller_test

import (
	"net/http"
	"net/http/httptest"

	"bitbucket.org/armakuni/raindrops-mb2/controller"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Router() *mux.Router {
	r := controller.Start()
	return r
}

func init() {
	http.Handle("/", Router())
}

var _ = Describe("The API", func() {
	var (
		req          *http.Request
		mockRecorder *httptest.ResponseRecorder
		err          error
	)

	BeforeEach(func() {
		mockRecorder = httptest.NewRecorder()
	})

	Context("when no number is passed in", func() {
		JustBeforeEach(func() {
			req, err = http.NewRequest("GET", "http://example.com/", nil)
			Expect(err).ToNot(HaveOccurred())
			Router().ServeHTTP(mockRecorder, req)
		})

		It("responds with an appropriate error", func() {
			Expect(mockRecorder.Code).To(Equal(400))
			Expect(mockRecorder.Body.String()).To(Equal("To use this API please provide a 'number' param"))
		})
	})

	Context("when number is provided as not a number", func() {
		JustBeforeEach(func() {
			req, err = http.NewRequest("GET", "http://example.com/?number=thing", nil)
			Expect(err).ToNot(HaveOccurred())
			Router().ServeHTTP(mockRecorder, req)
		})

		It("responds with an appropriate error", func() {
			Expect(mockRecorder.Code).To(Equal(400))
			Expect(mockRecorder.Body.String()).To(Equal("The 'number' param provided was 'thing', it must be a valid integer"))
		})
	})

	Context("when a number divisible by 3 is sent", func() {
		JustBeforeEach(func() {
			req, err = http.NewRequest("GET", "http://example.com/?number=3", nil)
			Expect(err).ToNot(HaveOccurred())
			Router().ServeHTTP(mockRecorder, req)
		})

		It("responds with Pling", func() {
			Expect(mockRecorder.Code).To(Equal(200))
			Expect(mockRecorder.Body.String()).To(Equal("Pling"))
		})
	})

	Context("when the input is divisible by only 5", func() {
		JustBeforeEach(func() {
			req, err = http.NewRequest("GET", "http://example.com/?number=5", nil)
			Expect(err).ToNot(HaveOccurred())
			Router().ServeHTTP(mockRecorder, req)
		})

		It("responds with Plang", func() {
			Expect(mockRecorder.Code).To(Equal(200))
			Expect(mockRecorder.Body.String()).To(Equal("Plang"))
		})
	})

	Context("when the input is divisible by only 7", func() {
		JustBeforeEach(func() {
			req, err = http.NewRequest("GET", "http://example.com/?number=7", nil)
			Expect(err).ToNot(HaveOccurred())
			Router().ServeHTTP(mockRecorder, req)
		})

		It("responds with Plong", func() {
			Expect(mockRecorder.Code).To(Equal(200))
			Expect(mockRecorder.Body.String()).To(Equal("Plong"))
		})
	})

	Context("when the input is divisible by only 3 and 5", func() {
		JustBeforeEach(func() {
			req, err = http.NewRequest("GET", "http://example.com/?number=15", nil)
			Expect(err).ToNot(HaveOccurred())
			Router().ServeHTTP(mockRecorder, req)
		})

		It("responds with PlingPlang", func() {
			Expect(mockRecorder.Code).To(Equal(200))
			Expect(mockRecorder.Body.String()).To(Equal("PlingPlang"))
		})
	})

	Context("when the input is divisible by 3 and 7", func() {
		JustBeforeEach(func() {
			req, err = http.NewRequest("GET", "http://example.com/?number=21", nil)
			Expect(err).ToNot(HaveOccurred())
			Router().ServeHTTP(mockRecorder, req)
		})

		It("responds with PlingPlong", func() {
			Expect(mockRecorder.Code).To(Equal(200))
			Expect(mockRecorder.Body.String()).To(Equal("PlingPlong"))
		})
	})

	Context("when the input is divisible by 5 and 7", func() {
		JustBeforeEach(func() {
			req, err = http.NewRequest("GET", "http://example.com/?number=35", nil)
			Expect(err).ToNot(HaveOccurred())
			Router().ServeHTTP(mockRecorder, req)
		})

		It("responds with PlangPlong", func() {
			Expect(mockRecorder.Code).To(Equal(200))
			Expect(mockRecorder.Body.String()).To(Equal("PlangPlong"))
		})
	})

	Context("when the input is divisible by 3, 5 and 7", func() {
		JustBeforeEach(func() {
			req, err = http.NewRequest("GET", "http://example.com/?number=105", nil)
			Expect(err).ToNot(HaveOccurred())
			Router().ServeHTTP(mockRecorder, req)
		})

		It("responds with PlingPlangPlong", func() {
			Expect(mockRecorder.Code).To(Equal(200))
			Expect(mockRecorder.Body.String()).To(Equal("PlingPlangPlong"))
		})
	})

	Context("when the input is not divisible by 3, 5 or 7", func() {
		JustBeforeEach(func() {
			req, err = http.NewRequest("GET", "http://example.com/?number=2", nil)
			Expect(err).ToNot(HaveOccurred())
			Router().ServeHTTP(mockRecorder, req)
		})

		It("responds with the input number as a string", func() {
			Expect(mockRecorder.Code).To(Equal(200))
			Expect(mockRecorder.Body.String()).To(Equal("2"))
		})
	})
})
