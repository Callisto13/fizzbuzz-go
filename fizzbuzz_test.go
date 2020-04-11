package fizzbuzz_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/callisto13/fizzbuzz"
)

var _ = Describe("Fizzbuzz", func() {
	Context("knows when a number is", func() {
		It("divisible by 3", func() {
			Expect(fizzbuzz.IsDivisibleBy(3, 3)).To(BeTrue())
		})

		It("knows when a number is NOT divisible by 3", func() {
			Expect(fizzbuzz.IsDivisibleBy(1, 3)).To(BeFalse())
		})

		It("knows when a number is divisible by 5", func() {
			Expect(fizzbuzz.IsDivisibleBy(5, 5)).To(BeTrue())
		})

		It("knows when a number is NOT divisible by 5", func() {
			Expect(fizzbuzz.IsDivisibleBy(1, 5)).To(BeFalse())
		})

		It("knows when a number is divisible by 3 and  5", func() {
			Expect(fizzbuzz.IsDivisibleBy(15, 15)).To(BeTrue())
		})

		It("knows when a number is NOT divisible by 3 or 5", func() {
			Expect(fizzbuzz.IsDivisibleBy(1, 15)).To(BeFalse())
		})
	})
})
