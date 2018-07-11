package raindrops_test

import (
	"math/rand"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bitbucket.org/armakuni/raindrops-mb2/raindrops"
)

var _ = Describe("#Raindrops", func() {
	rand.Seed(time.Now().UTC().UnixNano())

	Context("when the input is divisible by only 3", func() {
		It("returns Pling", func() {
			for i := 1; i <= 50; i++ {
				number := rand.Intn(1000) * 3
				if number%5 == 0 || number%7 == 0 {
					continue
				}
				Expect(raindrops.Process(number)).To(Equal("Pling"))
			}
			Expect(raindrops.Process(3)).To(Equal("Pling"))
		})
	})

	Context("when the input is divisible by only 5", func() {
		It("returns Plang", func() {
			for i := 1; i <= 50; i++ {
				number := rand.Intn(1000) * 5
				if number%3 == 0 || number%7 == 0 {
					continue
				}
				Expect(raindrops.Process(number)).To(Equal("Plang"))
			}
			Expect(raindrops.Process(5)).To(Equal("Plang"))
		})
	})

	Context("when the input is divisible by only 7", func() {
		It("returns Plong", func() {
			for i := 1; i <= 50; i++ {
				number := rand.Intn(1000) * 7
				if number%3 == 0 || number%5 == 0 {
					continue
				}
				Expect(raindrops.Process(number)).To(Equal("Plong"))
			}
			Expect(raindrops.Process(7)).To(Equal("Plong"))
		})
	})

	Context("when the input is divisible by only 3 and 5", func() {
		It("returns PlingPlang", func() {
			for i := 1; i <= 50; i++ {
				number := rand.Intn(1000) * 3 * 5
				if number%7 == 0 {
					continue
				}
				Expect(raindrops.Process(number)).To(Equal("PlingPlang"))
			}
			Expect(raindrops.Process(15)).To(Equal("PlingPlang"))
		})
	})

	Context("when the input is divisible by 3 and 7", func() {
		It("returns PlingPlong", func() {
			for i := 1; i <= 50; i++ {
				number := rand.Intn(1000) * 3 * 7
				if number%5 == 0 {
					continue
				}
				Expect(raindrops.Process(number)).To(Equal("PlingPlong"))
			}
			Expect(raindrops.Process(21)).To(Equal("PlingPlong"))
		})
	})

	Context("when the input is divisible by 5 and 7", func() {
		It("returns PlangPlong", func() {
			for i := 1; i <= 50; i++ {
				number := rand.Intn(1000) * 5 * 7
				if number%3 == 0 {
					continue
				}
				Expect(raindrops.Process(number)).To(Equal("PlangPlong"))
			}
			Expect(raindrops.Process(35)).To(Equal("PlangPlong"))
		})
	})

	Context("when the input is divisible by 3, 5 and 7", func() {
		It("returns PlingPlangPlong", func() {
			for i := 1; i <= 50; i++ {
				number := rand.Intn(1000) * 3 * 5 * 7
				Expect(raindrops.Process(number)).To(Equal("PlingPlangPlong"))
			}
			Expect(raindrops.Process(105)).To(Equal("PlingPlangPlong"))
		})
	})

	Context("when the input is not divisible by 3, 5 or 7", func() {
		It("returns the input number as a string", func() {
			for i := 1; i <= 50; i++ {
				number := rand.Intn(1000)
				if number%3 == 0 || number%5 == 0 || number%7 == 0 {
					continue
				}
				Expect(raindrops.Process(number)).To(Equal(strconv.Itoa(number)))
			}
			Expect(raindrops.Process(2)).To(Equal("2"))
		})
	})
})
